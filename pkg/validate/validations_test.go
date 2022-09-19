package validate_test

import (
	"span-challenge/pkg/models"
	"span-challenge/pkg/validate"
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
