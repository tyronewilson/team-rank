package validate

import (
	"github.com/pkg/errors"
	"spanchallenge/pkg/errs"
	"spanchallenge/pkg/models"
	"spanchallenge/pkg/util"
	"strconv"
)

func IsValidMatchResult(result models.MatchResult) (bool, error) {
	if !result.HasTwoTeams() {
		return false, errs.ErrMissingTeams(result.TeamA, result.TeamB)
	}
	return true, nil
}

func IsValidScoreString(str string) (bool, error) {
	teamName, scoreStr := util.SplitOnLastSpace(str)
	if teamName == "" || scoreStr == "" {
		return false, errs.ErrStringDoesNotEndWithSpaceAndScore(str)
	}
	if _, err := strconv.Atoi(scoreStr); err != nil {
		return false, errors.Wrapf(err, "score '%s' is not a valid score, expected an integer", scoreStr)
	}
	if len(teamName) < 2 {
		return false, errs.ErrTeamNameTooShort(teamName)
	}
	return true, nil
}
