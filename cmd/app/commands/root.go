package commands

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	cfg "github.com/DoraTiger/NEU_IPGW/config"
)

var (
	logger     = logrus.New()
	log_level  string
	log_format string
)

func init() {
	registerFlagsRootCmd(RootCmd)
}

func registerFlagsRootCmd(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&log_level, "log_level", cfg.DefaultLogLevel, "log level")
	cmd.PersistentFlags().StringVar(&log_format, "log_format", cfg.DefaultLogFormat, "log format")
}

// RootCmd is the root command for NEU_IPGW.
// config log level and log format
var RootCmd = &cobra.Command{
	Use:   "NEU_IPGW",
	Short: "NEU_IPGW is a command line interface for NEU internet connect",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		if cmd.Name() == VersionCmd.Name() {
			return nil
		}

		// set log level (info by default)
		switch log_level {
		case "error":
			logger.SetLevel(logrus.ErrorLevel)
			break
		case "warn":
			logger.SetLevel(logrus.WarnLevel)
			break
		case "info":
			logger.SetLevel(logrus.InfoLevel)
			break
		case "debug":
			logger.SetLevel(logrus.DebugLevel)
			break
		case "trace":
			logger.SetLevel(logrus.TraceLevel)
			break
		default:
			logger.SetLevel(logrus.InfoLevel)
		}

		// If the log_format flag value is "json", the log output format is set to JSON
		if log_format == "json" {
			logger.SetFormatter(&logrus.JSONFormatter{})
		}

		// Set the log output path
		logger.SetOutput(os.Stdout)

		return nil
	},
}
