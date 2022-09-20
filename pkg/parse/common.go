package parse

import (
	"github.com/pkg/errors"
	"spanchallenge/pkg/util"
	"spanchallenge/pkg/validate"
	"strconv"
	"strings"
)

// TeamScore takes a string in the input format TeamName score and returns the team name, score and any error
func TeamScore(str string) (string, int, error) {
	str = strings.TrimSpace(str)
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
