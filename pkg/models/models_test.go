package models_test

import (
	"span-challenge/pkg/models"
	"testing"
)

func TestMatchResult_IsDraw(t *testing.T) {
	type fields struct {
		TeamA      string
		TeamB      string
		TeamAScore int
		TeamBScore int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			// We don't care about the team names, only the outcome, so having blank names shouldn't matter
			name: "with empty fields and equal scores",
			fields: fields{
				TeamA:      "",
				TeamB:      "",
				TeamAScore: 0,
				TeamBScore: 0,
			},
			want: true,
		}, {
			name: "with non-empty fields and equal scores",
			fields: fields{
				TeamA:      "Team A",
				TeamB:      "Team B",
				TeamAScore: 1,
				TeamBScore: 1,
			},
			want: true,
		}, {
			name: "with non-empty fields and unequal scores where Team A has the higher score",
			fields: fields{
				TeamA:      "Team A",
				TeamB:      "Team B",
				TeamAScore: 1,
				TeamBScore: 0,
			},
			want: false,
		}, {
			name: "with non-empty fields and unequal scores where Team B has the higher score",
			fields: fields{
				TeamA:      "Team A",
				TeamB:      "Team B",
				TeamAScore: 0,
				TeamBScore: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := models.MatchResult{
				TeamA:      tt.fields.TeamA,
				TeamB:      tt.fields.TeamB,
				TeamAScore: tt.fields.TeamAScore,
				TeamBScore: tt.fields.TeamBScore,
			}
			if got := m.IsDraw(); got != tt.want {
				t.Errorf("IsDraw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchResult_HasTwoTeams(t *testing.T) {
	type fields struct {
		TeamA      string
		TeamB      string
		TeamAScore int
		TeamBScore int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "when there are two blank teams",
			fields: fields{
				TeamA: "",
				TeamB: "",
			},
			want: false,
		}, {
			name: "when Team A is blank",
			fields: fields{
				TeamA: "",
				TeamB: "Team B",
			},
			want: false,
		}, {
			name: "when Team B is blank",
			fields: fields{
				TeamA: "Team A",
				TeamB: "",
			},
			want: false,
		}, {
			name: "when there are two non-blank teams",
			fields: fields{
				TeamA: "Team A",
				TeamB: "Team B",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := models.MatchResult{
				TeamA:      tt.fields.TeamA,
				TeamB:      tt.fields.TeamB,
				TeamAScore: tt.fields.TeamAScore,
				TeamBScore: tt.fields.TeamBScore,
			}
			if got := m.HasTwoTeams(); got != tt.want {
				t.Errorf("HasTwoTeams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchResult_GetWinner(t *testing.T) {
	type fields struct {
		TeamA      string
		TeamB      string
		TeamAScore int
		TeamBScore int
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "when there are two blank teams",
			fields: fields{
				TeamA: "",
				TeamB: "",
			},
			want:    "",
			wantErr: true,
		}, {
			name: "when Team A is blank",
			fields: fields{
				TeamA: "",
				TeamB: "Team B",
			},
			want:    "",
			wantErr: true,
		}, {
			name: "when Team B is blank",
			fields: fields{
				TeamA: "Team A",
				TeamB: "",
			},
			want:    "",
			wantErr: true,
		}, {
			name: "when there are two non-blank teams and Team A has the higher score",
			fields: fields{
				TeamA:      "Team A",
				TeamB:      "Team B",
				TeamAScore: 1,
				TeamBScore: 0,
			},
			want:    "Team A",
			wantErr: false,
		}, {
			name: "when there are two non-blank teams and Team B has the higher score",
			fields: fields{
				TeamA:      "Team A",
				TeamB:      "Team B",
				TeamAScore: 0,
				TeamBScore: 1,
			},
			want:    "Team B",
			wantErr: false,
		}, {
			name: "when there are two non-blank teams and the scores are equal",
			fields: fields{
				TeamA:      "Team A",
				TeamB:      "Team B",
				TeamAScore: 1,
				TeamBScore: 1,
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := models.MatchResult{
				TeamA:      tt.fields.TeamA,
				TeamB:      tt.fields.TeamB,
				TeamAScore: tt.fields.TeamAScore,
				TeamBScore: tt.fields.TeamBScore,
			}
			got, err := m.GetWinner()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWinner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetWinner() got = %v, want %v", got, tt.want)
			}
		})
	}
}
