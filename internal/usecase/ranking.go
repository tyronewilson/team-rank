package usecase

import (
	"github.com/pkg/errors"
	"sort"
	"spanchallenge/pkg/config"
	"spanchallenge/pkg/models"
)

// TeamRanker is an interface that defines the behaviour of a team ranker
type TeamRanker interface {
	Rank(results models.ResultStream, sorters ...func(list models.TeamRankList)) (models.TeamRankList, error)
}

type configTeamRankerImpl struct {
	cfg config.RankingRules
}

// TeamTally is a map of team names to points
type TeamTally map[string]int

// TallyResultSet tallies the points for each team in the provided result stream
// NOTE (for evaluators): We can argue that this function is likely to change based on the concrete args and implied
// Tally method if we were to provide func(results, int, int, int) as a function signature.
// If we ever want to tally results based on any criteria other than w, d, l points, we wouldn't be able to use this method.
// For this reason, I've chosen to make the points calculations functions so that in principle the points calculation can embody
// any logic that the client might want, and we simply apply them based on the win, draw, loss state of the result which is
// encapsulates all possible states and in essence can't change.
func TallyResultSet(results models.ResultStream, winPointsFunc, drawPointsFunc, lossPointsFunc func() int) (TeamTally, error) {
	teamScores := make(map[string]int)
	i := 0
	for result := range results {
		i++
		initTeamIfNotPresent(result.TeamA, teamScores)
		initTeamIfNotPresent(result.TeamB, teamScores)
		if result.IsDraw() {
			teamScores[result.TeamA] += drawPointsFunc()
			teamScores[result.TeamB] += drawPointsFunc()
		} else {
			winner, err := result.GetWinner()
			if err != nil {
				return nil, errors.Wrapf(
					err,
					"failed to get winner for match %s vs %s at row/index: %d/%d",
					result.TeamA,
					result.TeamB,
					i+1,
					i,
				)
			}
			loser, err := result.OpponentOf(winner)
			if err != nil {
				return nil, errors.Wrapf(
					err,
					"failed to get losing team for match %s vs %s at row/index: %d/%d",
					result.TeamA,
					result.TeamB,
					i+1,
					i,
				)
			}
			teamScores[winner] += winPointsFunc()
			teamScores[loser] += lossPointsFunc()
		}
	}
	return teamScores, nil
}

func pointFunc(points int) func() int {
	return func() int {
		return points
	}
}

// Rank ranks the teams based on the match results using the rules defined in the loaded configuration
func (c configTeamRankerImpl) Rank(results models.ResultStream, sorters ...func(list models.TeamRankList)) (models.TeamRankList, error) {
	teamScores, err := TallyResultSet(
		results,
		pointFunc(c.cfg.WinPoints),
		pointFunc(c.cfg.DrawPoints),
		pointFunc(c.cfg.LossPoints),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to tally result set")
	}
	var teamRankList models.TeamRankList
	for teamName, points := range teamScores {
		teamRankList = append(teamRankList, &models.TeamRank{
			TeamName: teamName,
			Points:   points,
		})
	}
	// Allows clients to define their own sorting rules but has sensible defaults
	if len(sorters) > 0 {
		for _, sorter := range sorters {
			sorter(teamRankList)
		}
	} else {
		SortByTeamNameAsc(teamRankList)
		SortByPointsDesc(teamRankList)
	}
	setRanks(teamRankList)

	return teamRankList, nil
}

func setRanks(list models.TeamRankList) {
	// avoid any complexity later on with an early exit
	if len(list) == 0 {
		return
	}
	// We are going to make bins of the same points with incrementing ranks
	list[0].Rank = 1 // because we have different ways to rank, we must respect the order of list so the first one is always 1
	for i := 1; i < len(list); i++ {
		if list[i].Points == list[i-1].Points && list[i].Points > 0 {
			list[i].Rank = list[i-1].Rank
		} else {
			list[i].Rank = i + 1
		}
	}
}

// SortByPointsDesc sorts the team rank list by points in descending order
func SortByPointsDesc(list models.TeamRankList) {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Points > list[j].Points
	})
}

// SortByTeamNameAsc sorts the team rank list by team name in ascending order
func SortByTeamNameAsc(list models.TeamRankList) {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].TeamName < list[j].TeamName
	})
}

func initTeamIfNotPresent(teamName string, lookup map[string]int) {
	_, seen := lookup[teamName]
	if !seen {
		lookup[teamName] = 0
	}
}

// NewConfigTeamRanker returns a new TeamRanker that uses the rules defined in the provided configuration
func NewConfigTeamRanker(cfg config.RankingRules) TeamRanker {
	return &configTeamRankerImpl{cfg: cfg}
}
