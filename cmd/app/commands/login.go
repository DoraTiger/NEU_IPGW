package commands

import (
	"fmt"

	"errors"

	"github.com/DoraTiger/NEU_IPGW/pkg/handler"
	"github.com/DoraTiger/NEU_IPGW/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	username string
	password string
)

func init() {
	registerFlagsLoginCmd(LoginCmd)
}

func registerFlagsLoginCmd(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&username, "username", "u", "", "the username for authentication")
	cmd.MarkPersistentFlagRequired("username")
	cmd.PersistentFlags().StringVarP(&password, "password", "p", "", "the password for authentication")
	cmd.MarkPersistentFlagRequired("password")
}

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "commond for login",
	Run: func(cmd *cobra.Command, args []string) {
		// init err
		err := errors.New("")
		// init gwHandler
		gwHandler := handler.NewGWHandler()
		gwHandler.SetLogger(logger)
		// login to eone gw get cookie
		err = gwHandler.Login(username, password)
		if err != nil {
			file, line := utils.GetErrorLocation()
			logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
			fmt.Println(err)
			return
		}
		// init ipgwHandler
		ipgwHandler := handler.NewIPGWHandler()
		ipgwHandler.SetClient(gwHandler.GetClient())
		ipgwHandler.SetLogger(logger)
		//login to network gw
		msg, err := ipgwHandler.Login()
		if err != nil {
			file, line := utils.GetErrorLocation()
			logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
			fmt.Println(err)
			return
		}
		fmt.Println(msg)
	},
}
