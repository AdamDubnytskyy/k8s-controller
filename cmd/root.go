/*
Copyright © 2025 Oleh Adam Dubnytskyy <adam.dubnytskyy@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/AdamDubnytskyy/k8s-controller/pkg/logger"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var logLevel string

var rootCmd = &cobra.Command{
	Use:   "k8s-controller",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Init(logLevel)
		log.Info().Str("version", "dev").Msg("Starting k8s-controller")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "Set log level: trace, debug, info, warn, error")
}
