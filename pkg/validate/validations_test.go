package validate_test

import (
	"spanchallenge/pkg/models"
	"spanchallenge/pkg/validate"
	"testing"
)

func TestIsValidMatchResult(t *testing.T) {
	type args struct {
		result models.MatchResult
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "when there are two blank teams",
			args:    args{result: models.MatchResult{TeamA: "", TeamB: ""}},
			want:    false,
			wantErr: true,
		}, {
			name:    "when Team A is blank",
			args:    args{result: models.MatchResult{TeamA: "", TeamB: "Team B"}},
			want:    false,
			wantErr: true,
		}, {
			name:    "when Team B is blank",
			args:    args{result: models.MatchResult{TeamA: "Team A", TeamB: ""}},
			want:    false,
			wantErr: true,
		}, {
			name:    "when there are two non-blank teams",
			args:    args{result: models.MatchResult{TeamA: "Team A", TeamB: "Team B"}},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validate.IsValidMatchResult(tt.args.result)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsValidMatchResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsValidMatchResult() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidScoreString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "when the input is empty",
			args:    args{str: ""},
			want:    false,
			wantErr: true,
		}, {
			name:    "when the input is a single digit",
			args:    args{str: "1"},
			want:    false,
			wantErr: true,
		}, {
			name:    "when the input is a single digit with a space",
			args:    args{str: "1 "},
			want:    false,
			wantErr: true,
		}, {
			name:    "when the input is a single digit with a space and a digit but less than two chars in the team name",
			args:    args{str: "A 1"},
			want:    false,
			wantErr: true,
		}, {
			name:    "when the input is a single digit with a space and a digit and at least two chars in the team name",
			args:    args{str: "AB 1"},
			want:    true,
			wantErr: false,
		}, {
			name:    "when the team name is complex with multiple spaces and there is no integer at the end",
			args:    args{str: "My Awesome Team 123 B"},
			want:    false,
			wantErr: true,
		}, {
			name:    "when the team name is complex with multiple spaces and there's a valid score at the end",
			args:    args{str: "My Awesome Team 123 B 42"},
			want:    true,
			wantErr: false,
		}, {
			name:    "when the team name is valid but there's a non integer score at the end",
			args:    args{str: "My Awesome Team 123 B 42.5"},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validate.IsValidScoreString(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsValidScoreString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsValidScoreString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
