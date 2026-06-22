package commands

import (
	"fmt"
	"time"

	"github.com/DoraTiger/NEU_IPGW/internal/utils"
	"github.com/spf13/cobra"
)

var AccountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "commands for saved accounts",
}

var AccountsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list saved accounts",
	Run: func(cmd *cobra.Command, args []string) {
		store, err := newCredentialStore()
		if err != nil {
			fmt.Println(err)
			return
		}

		accounts, err := store.ListAccounts()
		if err != nil {
			file, line := utils.GetErrorLocation()
			logger.Debug(fmt.Sprintf("Error in file %s, line %d: %v", file, line, err))
			fmt.Println(err)
			return
		}

		if len(accounts) == 0 {
			data := map[string]interface{}{
				"accounts": []interface{}{},
			}
			utils.Output(data, func() {
				fmt.Println("no saved accounts")
			})
			return
		}

		accountList := make([]map[string]interface{}, len(accounts))
		for i, item := range accounts {
			accountList[i] = map[string]interface{}{
				"username":   item.Username,
				"is_last":    item.IsLast,
				"updated_at": item.UpdatedAt.Local().Format(time.RFC3339),
			}
		}

		data := map[string]interface{}{
			"accounts": accountList,
		}
		utils.Output(data, func() {
			fmt.Printf("%-20s %-8s %-25s\n", "username", "last", "updated_at")
			for _, item := range accounts {
				last := ""
				if item.IsLast {
					last = "*"
				}
				timeText := item.UpdatedAt.Local().Format(time.RFC3339)
				fmt.Printf("%-20s %-8s %-25s\n", item.Username, last, timeText)
			}
		})
	},
}

func init() {
	AccountsCmd.AddCommand(AccountsListCmd)
}
