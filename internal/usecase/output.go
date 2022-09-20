package usecase

import (
	"github.com/rs/zerolog/log"
	"io"
	"spanchallenge/pkg/models"
	"spanchallenge/pkg/serialize"
)

// WriteRankingsCSV writes the rankings to the writer in CSV format
func WriteRankingsCSV(list models.TeamRankList, w io.Writer) error {
	log.Debug().Caller().Msgf("writing rankings to CSV")
	if list == nil {
		log.Warn().Caller().Msgf("no rankings to write")
		return nil
	}
	bc, err := w.Write(serialize.TeamRanksAsCSV(list))
	if err != nil {
		return err
	}
	log.Trace().Caller().Int("bytes", bc).Msg("wrote rankings")
	return nil
}
