package cmd

// Copyright © 2022 Tyrone Wilson <tdubs241083@gmail.com>

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"io"
	"os"
	"span-challenge/cmd/answer"
	"span-challenge/cmd/ask"
	"span-challenge/internal/usecase"
	"span-challenge/pkg/config"
	"span-challenge/pkg/util"
)

var rankCmd = &cobra.Command{
	Use:   "rank",
	Short: "Takes an input file of team match results and outputs a ranked list of teams",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal().Msg("missing input file(s) must have at least one file argument to rank")
		}
		type cmdConfig struct {
			Rules config.RankingRules
		}
		var cfg cmdConfig
		util.MaybePanic(loadConfig(cmd, &cfg))
		runner := usecase.NewConfigTeamRanker(cfg.Rules)
		filenames := args
		err := util.CheckAllFilesExist(filenames)
		util.MaybePanic(err)
		var inputs []io.Reader
		var toClose []io.Closer
		for _, filename := range filenames {
			filename := filename
			file, err := os.Open(filename)
			util.MaybePanic(err)
			inputs = append(inputs, file)
			toClose = append(toClose, file)
		}
		defer func() {
			for _, closer := range toClose {
				_ = closer.Close()
			}
		}()
		errs := util.NewErrorCollector()
		results := usecase.StreamResults(errs, inputs...)
		rankings, err := runner.Rank(results)
		util.MaybePanic(err)
		outputType, err := cmd.Flags().GetString("dest")
		util.MaybePanic(err)
		if outputType == "" {
			outputType, err = ask.OutputType()
			util.MaybePanic(err)
		}
		var output io.Writer
		switch outputType {
		case answer.FileOutput:
			filename, err := cmd.Flags().GetString("filename")
			util.MaybePanic(err)
			if filename == "" {
				filename, err = ask.OutputFilename()
				util.MaybePanic(err)
			}
			file, err := os.Create(filename)
			util.MaybePanic(err)
			defer func() {
				err := file.Close()
				util.MaybePanic(err)
			}()
			output = file
		case answer.StdOut:
			output = os.Stdout
		}
		err = usecase.WriteRankingsCSV(rankings, output)
		util.MaybePanic(err)
	},
}

func loadConfig(cmd *cobra.Command, dest interface{}) error {
	filename, err := cmd.Flags().GetString("config")
	if err != nil {
		return err
	}
	if filename != "" {
		log.Info().Msgf("loading env config from %s", filename)
		err = cleanenv.ReadConfig(filename, dest)
		if err != nil {
			return err
		}
		return nil
	}
	log.Info().Msg("load environment config")
	return cleanenv.ReadEnv(dest)
}

func init() {
	rootCmd.AddCommand(rankCmd)

	// TODO provide different format styles
	// rankCmd.Flags().StringP("format", "f", "csv", "output format")
	rankCmd.Flags().StringP("dest", "d", "", "--dest file or -d stdout")
	rankCmd.Flags().StringP("filename", "n", "", "--filename output.csv or -n output.csv")
}
