package validate

import (
	"span-challenge/pkg/errs"
	"span-challenge/pkg/models"
)

func IsValidMatchResult(result models.MatchResult) (bool, error) {
	if !result.HasTwoTeams() {
		return false, errs.ErrMissingTeams(result.TeamA, result.TeamB)
	}
	return true, nil
}
