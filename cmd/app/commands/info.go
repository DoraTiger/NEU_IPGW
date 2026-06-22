package commands

import (
	"errors"
	"fmt"

	"github.com/DoraTiger/NEU_IPGW/internal/handler"
	"github.com/DoraTiger/NEU_IPGW/internal/i18n"
	"github.com/DoraTiger/NEU_IPGW/internal/utils"
	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "show current IPGW usage info",
	Run: func(cmd *cobra.Command, args []string) {
		lang := currentLanguage()
		ipgwHandler := handler.NewIPGWHandler()
		ipgwHandler.SetLogger(logger)

		info, err := ipgwHandler.FetchOnlineInfo()
		if err != nil {
			if errors.Is(err, handler.ErrIPGWNotOnline) {
				data := map[string]interface{}{
					"error": i18n.T(lang, i18n.MsgNotLoggedIn),
				}
				utils.Output(data, func() {
					fmt.Println(i18n.T(lang, i18n.MsgNotLoggedIn))
				})
				return
			}
			file, line := utils.GetErrorLocation()
			logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
			data := map[string]interface{}{
				"error": fmt.Sprintf("%s: %v", i18n.T(lang, i18n.MsgUsageFetchFail), err),
			}
			utils.Output(data, func() {
				fmt.Printf("%s: %v\n", i18n.T(lang, i18n.MsgUsageFetchFail), err)
			})
			return
		}

		printOnlineInfo(lang, info)
	},
}

func printOnlineInfo(lang string, info *handler.OnlineInfo) {
	data := map[string]interface{}{
		"username": info.Username,
		"traffic":  info.Traffic(),
		"duration": info.Duration(),
		"balance":  info.Balance,
		"ip":       info.IP,
	}

	utils.Output(data, func() {
		fmt.Printf("%s: %s\n", i18n.T(lang, i18n.LabelAccount), info.Username)
		fmt.Printf("%s: %s\n", i18n.T(lang, i18n.LabelTraffic), info.Traffic())
		fmt.Printf("%s: %s\n", i18n.T(lang, i18n.LabelDuration), info.Duration())
		fmt.Printf("%s: %.2f\n", i18n.T(lang, i18n.LabelBalance), info.Balance)
		fmt.Printf("%s: %s\n", i18n.T(lang, i18n.LabelOnlineIP), info.IP)
	})
}
