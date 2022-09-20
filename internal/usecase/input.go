package usecase

import (
	"bufio"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
	"span-challenge/pkg/models"
	"span-challenge/pkg/parse"
	"span-challenge/pkg/util"
	"strings"
)

// StreamResults takes a list of input readers and an error handler and streams marshalled results from each line for each reader
func StreamResults(errs util.ErrorHandler, inputs ...io.Reader) models.ResultStream {
	out := make(models.ResultStream)
	go func() {
		defer close(out)
		for i, input := range inputs {
			_scanner := bufio.NewScanner(input)
			count := 0
			// optionally, resize scanner's capacity for lines over 64K, see next example
			for _scanner.Scan() {
				count++
				line := _scanner.Text()
				items := strings.Split(line, ",")
				if len(items) != 2 {
					log.Error().Caller().Int("inputNo", i).Int("line", count).Msgf("invalid input: %s", line)
					errs.HandleError(errors.Errorf("invalid line %s at %d in input: %d", line, count, i))
					continue
				}
				teamA, scoreA, err := parse.ParseScore(items[0])
				if err != nil {
					log.Error().Err(err).Caller().Int("input", i).Int("line", count).Msgf("unable to parse score: %s", items[0])
					errs.HandleError(errors.Wrapf(err, "input: %d line: %d", i, count))
					continue
				}
				teamB, scoreB, err := parse.ParseScore(items[1])
				if err != nil {
					log.Error().Err(err).Caller().Int("input", i).Int("line", count).Msgf("unable to parse score: %s", items[1])
					errs.HandleError(errors.Wrapf(err, "input: %d line: %d", i, count))
					continue
				}
				out <- &models.MatchResult{
					TeamA:      teamA,
					TeamB:      teamB,
					TeamAScore: scoreA,
					TeamBScore: scoreB,
				}
			}
			if err := _scanner.Err(); err != nil {
				errs.HandleError(err)
			}
		}
	}()
	return out
}
