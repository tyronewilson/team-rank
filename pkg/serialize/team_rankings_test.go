package serialize_test

import (
	"math/rand"
	"reflect"
	"span-challenge/pkg/models"
	"span-challenge/pkg/platform"
	"span-challenge/pkg/serialize"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestTeamRankList_SerializeCSV(t *testing.T) {
	tests := []struct {
		name string
		l    models.TeamRankList
		want []byte
	}{
		{
			name: "with nil list",
			l:    nil,
			want: nil,
		}, {
			name: "with empty list",
			l:    models.TeamRankList{},
			want: nil,
		}, {
			name: "with one team",
			l: models.TeamRankList{
				&models.TeamRank{
					TeamName: "Team A",
					Points:   2,
					Rank:     1,
				},
			},
			want: []byte("1. Team A, 2 pts" + platform.LineSeparator),
		}, {
			name: "with multiple teams",
			l: models.TeamRankList{
				&models.TeamRank{
					TeamName: "Team A",
					Points:   2,
					Rank:     1,
				}, &models.TeamRank{
					TeamName: "Team B",
					Points:   2,
					Rank:     1,
				}, &models.TeamRank{
					TeamName: "Team C",
					Points:   0,
					Rank:     3,
				},
			},
			want: []byte(strings.ReplaceAll("1. Team A, 2 pts\n1. Team B, 2 pts\n3. Team C, 0 pts\n", `\n`, platform.LineSeparator)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := serialize.TeamRanksAsCSV(tt.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TeamRanksAsCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTeamRanksAsCSV10(b *testing.B) {
	input := randomTeamRankList(10)
	for i := 0; i < b.N; i++ {
		serialize.TeamRanksAsCSV(input)
	}
}

func BenchmarkTeamRanksAsCSV1000(b *testing.B) {
	input := randomTeamRankList(1000)
	for i := 0; i < b.N; i++ {
		serialize.TeamRanksAsCSV(input)
	}
}

func BenchmarkTeamRanksAsCSV100000(b *testing.B) {
	input := randomTeamRankList(100000)
	for i := 0; i < b.N; i++ {
		serialize.TeamRanksAsCSV(input)
	}
}

func randomTeamRankList(limit int) models.TeamRankList {
	result := make(models.TeamRankList, limit)
	for i := 0; i < limit; i++ {
		r := &models.TeamRank{
			TeamName: "Team " + strconv.Itoa(i+1),
			Points:   rand.Intn(10),
			Rank:     rand.Intn(limit),
		}
		result[i] = r
	}
	return result
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
