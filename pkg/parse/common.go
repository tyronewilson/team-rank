package parse

import (
	"github.com/pkg/errors"
	"span-challenge/pkg/util"
	"span-challenge/pkg/validate"
	"strconv"
)

// ParseScore takes a string in the input format TeamName score and returns the team name, score and any error
func ParseScore(str string) (string, int, error) {
	if valid, err := validate.IsValidScoreString(str); !valid {
		return "", 0, err
	}
	teamName, scoreStr := util.SplitOnLastSpace(str)
	score, err := strconv.Atoi(scoreStr)
	if err != nil {
		return "", 0, errors.Wrapf(err, "failed to parse score %s", scoreStr)
	}
	return teamName, score, nil
}
