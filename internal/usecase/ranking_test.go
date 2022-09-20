package usecase_test

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"span-challenge/internal/usecase"
	"span-challenge/pkg/config"
	"span-challenge/pkg/models"
	"testing"
)

func TestSortByPointsDesc(t *testing.T) {
	rankA := &models.TeamRank{
		TeamName: "Team A",
		Points:   42,
	}
	rankB := &models.TeamRank{
		TeamName: "Team B",
		Points:   21,
	}
	inOrder := models.TeamRankList{rankA, rankB}
	outOfOrder := models.TeamRankList{rankB, rankA}
	usecase.SortByPointsDesc(outOfOrder)
	assert.Equal(t, inOrder, outOfOrder)
	rankC := &models.TeamRank{
		TeamName: "Team C",
		Points:   42,
	}
	inOrder = models.TeamRankList{rankA, rankC, rankB}
	outOfOrder = models.TeamRankList{rankB, rankA, rankC}
	usecase.SortByPointsDesc(outOfOrder)
	assert.Equal(t, inOrder, outOfOrder)
	// Ensure that we don't have items swapping if they have the same points
	inOrder = models.TeamRankList{rankC, rankA, rankB}
	outOfOrder = models.TeamRankList{rankB, rankC, rankA}
	usecase.SortByPointsDesc(outOfOrder)
	assert.Equal(t, inOrder, outOfOrder)
}

func TestSortByTeamNameAsc(t *testing.T) {
	// basic example of sorting by team name
	teamA := &models.TeamRank{TeamName: "Team A", Points: 1}
	teamB := &models.TeamRank{TeamName: "Team B", Points: 3}
	teamC := &models.TeamRank{TeamName: "Team C", Points: 42}
	inOrder := models.TeamRankList{teamA, teamB, teamC}
	outOfOrder := models.TeamRankList{teamB, teamC, teamA}
	usecase.SortByTeamNameAsc(outOfOrder)
	assert.Equal(t, inOrder, outOfOrder)

	// check ranking with subtle differences in team names including number of letters
	randomA := &models.TeamRank{TeamName: "AAAAA", Points: 1}
	randomAA := &models.TeamRank{TeamName: "AAAAAA", Points: 1}
	randomAB := &models.TeamRank{TeamName: "AAAAAB", Points: 1}
	randomABB := &models.TeamRank{TeamName: "AAAAABB", Points: 1}
	inOrder = models.TeamRankList{randomA, randomAA, randomAB, randomABB}
	outOfOrder = models.TeamRankList{randomAB, randomAA, randomABB, randomA}
	usecase.SortByTeamNameAsc(outOfOrder)
	assert.Equal(t, inOrder, outOfOrder)
}

func TestTallyResultSet(t *testing.T) {
	noPoints := func() int { return 0 }
	onePoint := func() int { return 1 }
	threePoints := func() int { return 3 }
	results := models.ResultSet{
		{
			TeamA:      "Team 1",
			TeamB:      "Team 2",
			TeamAScore: 4,
			TeamBScore: 2,
		}, {
			TeamA:      "Team 2",
			TeamB:      "Team 3",
			TeamAScore: 1,
			TeamBScore: 1,
		}, {
			TeamA:      "Team 3",
			TeamB:      "Team 1",
			TeamAScore: 0,
			TeamBScore: 3,
		},
	}
	streamResults := func() models.ResultStream {
		out := make(models.ResultStream)
		go func() {
			for _, result := range results {
				out <- result
			}
			close(out)
		}()
		return out
	}
	emptyStream := func() models.ResultStream {
		out := make(models.ResultStream)
		go func() {
			close(out)
		}()
		return out
	}

	type args struct {
		results        models.ResultStream
		winPointsFunc  func() int
		drawPointsFunc func() int
		lossPointsFunc func() int
	}
	tests := []struct {
		name    string
		args    args
		want    usecase.TeamTally
		wantErr bool
	}{
		{
			name: "with empty input",
			args: args{
				results:        emptyStream(),
				winPointsFunc:  threePoints,
				drawPointsFunc: onePoint,
				lossPointsFunc: noPoints,
			},
			want:    usecase.TeamTally{},
			wantErr: false,
		}, {
			name: "with valid input",
			args: args{
				results:        streamResults(),
				winPointsFunc:  threePoints,
				drawPointsFunc: onePoint,
				lossPointsFunc: noPoints,
			},
			want: usecase.TeamTally{
				"Team 1": 6,
				"Team 2": 1,
				"Team 3": 1,
			},
			wantErr: false,
		}, {
			name: "with no points assigned for anything",
			args: args{
				results:        streamResults(),
				winPointsFunc:  noPoints,
				drawPointsFunc: noPoints,
				lossPointsFunc: noPoints,
			},
			want: usecase.TeamTally{
				"Team 1": 0,
				"Team 2": 0,
				"Team 3": 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := usecase.TallyResultSet(tt.args.results, tt.args.winPointsFunc, tt.args.drawPointsFunc, tt.args.lossPointsFunc)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equalf(t, tt.want, got, "TallyResultSet(%v, winPoinsFunc, drawPointsFunc, lossPointsFunc)", tt.args.results)
		})
	}
}

func Test_configTeamRankerImpl_Rank(t *testing.T) {
	defaultRules := config.RankingRules{
		WinPoints:  3,
		DrawPoints: 1,
		LossPoints: 0,
	}
	results := models.ResultSet{
		{
			TeamA:      "Team 1",
			TeamB:      "Team 2",
			TeamAScore: 4,
			TeamBScore: 2,
		}, {
			TeamA:      "Team 2",
			TeamB:      "Team 3",
			TeamAScore: 2,
			TeamBScore: 1,
		}, {
			TeamA:      "Team 3",
			TeamB:      "Team 1",
			TeamAScore: 0,
			TeamBScore: 3,
		},
	}
	streamResults := func() models.ResultStream {
		out := make(models.ResultStream)
		go func() {
			for _, result := range results {
				out <- result
			}
			close(out)
		}()
		return out
	}
	emptyStream := func() models.ResultStream {
		out := make(models.ResultStream)
		go func() {
			close(out)
		}()
		return out
	}

	type fields struct {
		cfg config.RankingRules
	}
	type args struct {
		results models.ResultStream
		sorters []func(list models.TeamRankList)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.TeamRankList
		wantErr bool
	}{
		{
			name:   "with empty input and no sorters",
			fields: fields{cfg: defaultRules},
			args: args{
				results: emptyStream(),
				sorters: nil,
			},
			want:    nil,
			wantErr: false,
		}, {
			name:   "with non-empty input and empty sorters",
			fields: fields{cfg: defaultRules},
			args: args{
				results: streamResults(),
				sorters: nil,
			},
			want: models.TeamRankList{
				{TeamName: "Team 1", Points: 6},
				{TeamName: "Team 2", Points: 3},
				{TeamName: "Team 3", Points: 0},
			},
			wantErr: false,
		}, {
			name:   "with non-empty input and non-standard sorters",
			fields: fields{cfg: defaultRules},
			args: args{
				results: streamResults(),
				sorters: []func(list models.TeamRankList){
					// inline function to reverse sort by points
					func(list models.TeamRankList) {
						sort.Slice(list, func(i, j int) bool {
							return list[i].Points < list[j].Points
						})
					},
				},
			},
			want: models.TeamRankList{
				{TeamName: "Team 3", Points: 0},
				{TeamName: "Team 2", Points: 3},
				{TeamName: "Team 1", Points: 6},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := usecase.NewConfigTeamRanker(tt.fields.cfg)
			got, err := c.Rank(tt.args.results, tt.args.sorters...)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			for i, rank := range tt.want {
				assert.Equal(t, rank.TeamName, got[i].TeamName)
				assert.Equal(t, rank.Points, got[i].Points)
			}
		})
	}
}
