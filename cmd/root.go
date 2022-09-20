package cmd

// Copyright © 2022 Tyrone Wilson <tdubs241083@gmail.com>

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "CLI tool to read, parse and rank team scores",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		logLevel, err := cmd.Flags().GetString("log-level")
		if err != nil {
			return err
		}
		level, err := zerolog.ParseLevel(logLevel)
		if err != nil {
			return err
		}
		zerolog.SetGlobalLevel(level)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	rootCmd.PersistentFlags().StringP("config", "c", "", "--config settings1.env or -c settings2.env")
	rootCmd.PersistentFlags().StringP("log-level", "l", "info", "--log-level debug or -l debug")
}
