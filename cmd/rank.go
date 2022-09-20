package cmd

// Copyright © 2022 Tyrone Wilson <tdubs241083@gmail.com>

import (
	"github.com/spf13/cobra"
)

var rankCmd = &cobra.Command{
	Use:   "rank",
	Short: "Takes an input file of team match results and outputs a ranked list of teams",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(rankCmd)
}
