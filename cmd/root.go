package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var logger ilogger

var rootCmd = &cobra.Command{
	Short: "Authentication microservice that will store the session in local storage.",
	Run:   root,
}

func Execute(l ilogger) error {
	logger = l
	return rootCmd.Execute()
}

func root(cmd *cobra.Command, args []string) {
	cmd.Usage()
}

type ilogger interface {
	Debug(string, ...zap.Field)
	Panic(string, ...zap.Field)
	Warn(string, ...zap.Field)
	Info(string, ...zap.Field)
}
