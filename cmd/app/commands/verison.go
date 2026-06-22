package commands

import (
	"fmt"

	"github.com/DoraTiger/NEU_IPGW/internal/utils"
	"github.com/DoraTiger/NEU_IPGW/version"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "show the version info",
	Run: func(cmd *cobra.Command, args []string) {
		data := map[string]interface{}{
			"version": version.BuildVersion,
			"build":   version.BuildTime,
			"repo":    version.BuildRepo,
		}

		utils.Output(data, func() {
			fmt.Println(version.BuildVersion)
		})
	},
}
