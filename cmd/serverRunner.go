package cmd

import (
	"github.com/dmatusiewicz/auth/pkg/server"
	"github.com/spf13/cobra"
)

var signingKey = []byte("default_signing_key")

func init() {
	serverRunnerCmd.PersistentFlags().BytesHexVarP(&signingKey, "signingKey", "s", signingKey, "Signing Key for generating JWT tokens")
	rootCmd.AddCommand(serverRunnerCmd)
}

var serverRunnerCmd = &cobra.Command{
	Run:   serverRunner,
	Short: "Start the authentication server",
	Use:   "server",
}

func serverRunner(cmd *cobra.Command, args []string) {
	logger.Debug("Starting auth server.")
	server.Run(signingKey, logger)
}
