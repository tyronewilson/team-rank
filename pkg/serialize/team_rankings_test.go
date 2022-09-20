package serialize_test

import (
	"reflect"
	"span-challenge/pkg/models"
	"span-challenge/pkg/serialize"
	"testing"
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
				},
			},
			want: []byte("1. Team A, 2 pts\n"),
		}, {
			name: "with multiple teams",
			l: models.TeamRankList{
				&models.TeamRank{
					TeamName: "Team A",
					Points:   2,
				}, &models.TeamRank{
					TeamName: "Team B",
					Points:   2,
				}, &models.TeamRank{
					TeamName: "Team C",
					Points:   0,
				},
			},
			want: []byte("1. Team A, 2 pts\n2. Team B, 2 pts\n3. Team C, 0 pts\n"),
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
