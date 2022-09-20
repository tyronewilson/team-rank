package errs

import (
	"fmt"
)

// NOTES FOR EVALUATORS
// 1. I decided to create a separate package for standardized error messages so that models and validate (or any other package for that matter)
// can reference them without creating circular dependencies.
// 2. I this package errs because there are useful golang packages which are called errors and I don't want to keep tip toeing around
// import aliases etc.

func ErrMissingTeams(teamA, teamB string) error {
	return fmt.Errorf("match must have two teams but provided teams were '%s' vs '%s'", teamA, teamB)
}

func ErrStringDoesNotEndWithSpaceAndScore(str string) error {
	return fmt.Errorf("provided string must have a space and integer at the end to be considered a score got: '%s'", str)
}

func ErrTeamNameTooShort(teamName string) error {
	return fmt.Errorf("team name '%s' is too short, must be at least two characters", teamName)
}
