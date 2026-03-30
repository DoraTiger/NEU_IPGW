package commands

import (
	"fmt"
	"strings"

	"github.com/DoraTiger/NEU_IPGW/internal/handler"
	"github.com/DoraTiger/NEU_IPGW/internal/utils"
	"github.com/spf13/cobra"
)

var (
	username string
	password string
	saveCred bool
	account  string
)

func init() {
	registerFlagsLoginCmd(LoginCmd)
}

func registerFlagsLoginCmd(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&username, "username", "u", "", "the username for authentication")
	cmd.Flags().StringVarP(&password, "password", "p", "", "the password for authentication")
	cmd.Flags().BoolVar(&saveCred, "save", false, "save credential locally and overwrite same username")
	cmd.Flags().StringVar(&account, "account", "", "saved account username to load")
}

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "command for login",
	Run: func(cmd *cobra.Command, args []string) {
		credentialUsername := strings.TrimSpace(username)
		credentialPassword := password

		store, err := newCredentialStore()
		if err != nil {
			fmt.Println(err)
			return
		}

		if credentialUsername == "" && credentialPassword == "" {
			cred, err := loadCredentialForLogin(store, strings.TrimSpace(account))
			if err != nil {
				file, line := utils.GetErrorLocation()
				logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
				fmt.Println(err)
				return
			}
			credentialUsername = cred.Username
			credentialPassword = cred.Password
		} else if credentialUsername == "" || credentialPassword == "" {
			fmt.Println("both username and password are required when using inline credential")
			return
		}

		// init gwHandler
		gwHandler := handler.NewGWHandler()
		gwHandler.SetLogger(logger)
		// login to eone gw get cookie
		err = gwHandler.Login(credentialUsername, credentialPassword)
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

		if saveCred {
			if err := store.Save(credentialUsername, credentialPassword); err != nil {
				file, line := utils.GetErrorLocation()
				logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
				fmt.Println(err)
				return
			}
		} else {
			_ = store.MarkLast(credentialUsername)
		}

		fmt.Println(msg)
	},
}

func loadCredentialForLogin(store *utils.CredentialStore, selectedAccount string) (*utils.SavedCredential, error) {
	if selectedAccount != "" {
		cred, err := store.LoadByUsername(selectedAccount)
		if err != nil {
			if err == utils.ErrAccountNotFound {
				return nil, fmt.Errorf("saved account %q not found", selectedAccount)
			}
			return nil, err
		}
		return cred, nil
	}

	cred, err := store.LoadLast()
	if err != nil {
		if err == utils.ErrAccountNotFound {
			return nil, fmt.Errorf("no inline credential and no saved account available")
		}
		return nil, err
	}
	return cred, nil
}
