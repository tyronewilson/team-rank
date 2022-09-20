package models

import (
	"github.com/pkg/errors"
	"spanchallenge/pkg/errs"
)

// MatchResult embodies a single match result TeamA vs TeamB and their respective scores
type MatchResult struct {
	TeamA      string
	TeamB      string
	TeamAScore int
	TeamBScore int
}

// IsDraw returns true if the match is a draw
func (m MatchResult) IsDraw() bool {
	return m.TeamAScore == m.TeamBScore
}

// GetWinner returns the name of the winning team
// returns "", nil if it was in fact a draw
func (m MatchResult) GetWinner() (string, error) {
	if !m.HasTwoTeams() {
		return "", errs.ErrMissingTeams(m.TeamA, m.TeamB)
	}
	if m.IsDraw() {
		return "", nil
	}
	if m.TeamAScore > m.TeamBScore {
		return m.TeamA, nil
	}
	return m.TeamB, nil
}

// HasTwoTeams returns true if the match has two teams
// ideally validations would be a separate concern but this method is necessary to provide a meaningful result for GetWinner
// Given that we will do validations on models, it is highly likely that there would be a circular dependency if we were
// to put this function inside some kind of validation package
// The method is exported because we want to test it in the models_test package. It could be tested in the same package but given
// that the method is a simple statement of truth and innocuous, it is not worth the pollution to test it in the same package
func (m MatchResult) HasTwoTeams() bool {
	return m.TeamA != "" && m.TeamB != ""
}

// OpponentOf returns the name of the opponent of the input team
func (m MatchResult) OpponentOf(teamName string) (string, error) {
	if teamName == "" {
		return "", errors.New("provided name cannot be empty")
	}
	if teamName == m.TeamA {
		return m.TeamB, nil
	}
	if teamName == m.TeamB {
		return m.TeamA, nil
	}
	return "", errors.Errorf("team '%s' is not part of the match", teamName)
}

// ResultSet is a collection of MatchResults
type ResultSet []*MatchResult

type ResultStream chan *MatchResult

type TeamRank struct {
	TeamName string
	Points   int
	Rank     int
}

type TeamRankList []*TeamRank
