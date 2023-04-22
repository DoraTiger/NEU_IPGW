package commands

import (
	"fmt"

	"errors"

	"github.com/DoraTiger/NEU_IPGW/pkg/handler"
	"github.com/DoraTiger/NEU_IPGW/pkg/utils"
	"github.com/spf13/cobra"
)

var LogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "commond for logout",
	Run: func(cmd *cobra.Command, args []string) {
		// init err
		err := errors.New("")

		// init ipgwHandler
		ipgwHandler := handler.NewIPGWHandler()
		ipgwHandler.SetLogger(logger)

		//logout from campus nerwork
		err = ipgwHandler.Logout()
		if err != nil {
			file, line := utils.GetErrorLocation()
			logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
			fmt.Println(err)
			return
		}
		fmt.Println("logout success")
	},
}
