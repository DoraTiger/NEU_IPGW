package commands

import (
	"fmt"

	"github.com/DoraTiger/NEU_IPGW/internal/handler"
	"github.com/DoraTiger/NEU_IPGW/internal/utils"
	"github.com/spf13/cobra"
)

var (
	forgetCredential bool
	forgetAccount    string
)

func init() {
	registerFlagsLogoutCmd(LogoutCmd)
}

func registerFlagsLogoutCmd(cmd *cobra.Command) {
	cmd.Flags().BoolVar(&forgetCredential, "forget", false, "delete saved local credential")
	cmd.Flags().StringVar(&forgetAccount, "account", "", "saved account username for --forget")
}

var LogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "command for logout",
	Run: func(cmd *cobra.Command, args []string) {
		// init ipgwHandler
		ipgwHandler := handler.NewIPGWHandler()
		ipgwHandler.SetLogger(logger)

		//logout from campus nerwork
		err := ipgwHandler.Logout()
		if err != nil {
			file, line := utils.GetErrorLocation()
			logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
			fmt.Println(err)
			return
		}

		if forgetCredential {
			if err := forgetSavedCredential(forgetAccount); err != nil {
				file, line := utils.GetErrorLocation()
				logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
				fmt.Println(err)
				return
			}
		}

		fmt.Println("logout success")
	},
}

func forgetSavedCredential(selectedAccount string) error {
	store, err := newCredentialStore()
	if err != nil {
		return err
	}

	if selectedAccount != "" {
		deleted, err := store.DeleteByUsername(selectedAccount)
		if err != nil {
			return err
		}
		if !deleted {
			return fmt.Errorf("saved account %q not found", selectedAccount)
		}
		fmt.Printf("saved account %q deleted\n", selectedAccount)
		return nil
	}

	username, deleted, err := store.DeleteLast()
	if err != nil {
		return err
	}
	if !deleted {
		return fmt.Errorf("no saved account to delete")
	}
	fmt.Printf("saved account %q deleted\n", username)
	return nil
}
