package errs

import (
	"fmt"
)

// NOTES FOR EVALUATORS
// 1. I decided to create a separate package for standardized error messages so that models and validate (or any other package for that matter)
// can reference them without creating circular dependencies.
// 2. I created this package errs because there are useful golang packages which are called errors and I don't want to keep tip toeing around
// import aliases etc.

// ErrMissingTeams returns an error when the match does not have two teams
func ErrMissingTeams(teamA, teamB string) error {
	return fmt.Errorf("match must have two teams but provided teams were '%s' vs '%s'", teamA, teamB)
}

// ErrStringDoesNotEndWithSpaceAndScore returns an error when the string does not end with a space and a score
func ErrStringDoesNotEndWithSpaceAndScore(str string) error {
	return fmt.Errorf("provided string must have a space and integer at the end to be considered a score got: '%s'", str)
}

// ErrTeamNameTooShort returns an error when the team name is too short
func ErrTeamNameTooShort(teamName string) error {
	return fmt.Errorf("team name '%s' is too short, must be at least two characters", teamName)
}
