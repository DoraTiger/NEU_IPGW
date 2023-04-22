package main

import (
	"fmt"
	"os"

	cmd "github.com/DoraTiger/NEU_IPGW/cmd/app/commands"
)

func main() {
	rootCmd := cmd.RootCmd
	rootCmd.AddCommand(
		cmd.VersionCmd,
		cmd.LoginCmd,
		cmd.LogoutCmd,
	)
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v", err)
	}
}
