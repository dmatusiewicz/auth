package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var commitId string
var buildDate string
var builder string

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Run:   version,
	Use:   "version",
	Short: "Prints the version of the application.",
}

func version(cmd *cobra.Command, args []string) {
	logger.Info("Printing version",
		zap.String("CommitId", commitId),
		zap.String("BuildDate", buildDate),
		zap.String("Builder", builder),
	)
}
