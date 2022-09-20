package parse_test

import (
	"spanchallenge/pkg/parse"
	"spanchallenge/pkg/util"
	"testing"
)

func TestSplitOnLastSpace(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name:  "when the string is empty",
			args:  args{str: ""},
			want:  "",
			want1: "",
		}, {
			name:  "when the string has no spaces",
			args:  args{str: "abc"},
			want:  "",
			want1: "",
		}, {
			name:  "when the string has one space",
			args:  args{str: "abc def"},
			want:  "abc",
			want1: "def",
		}, {
			name:  "when the string has multiple spaces",
			args:  args{str: "abc def ghi"},
			want:  "abc def",
			want1: "ghi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := util.SplitOnLastSpace(tt.args.str)
			if got != tt.want {
				t.Errorf("SplitOnLastSpace() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SplitOnLastSpace() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParseScore(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   int
		wantErr bool
	}{
		{
			name:    "when the input is empty",
			args:    args{str: ""},
			want:    "",
			want1:   0,
			wantErr: true,
		}, {
			name:    "when the input has no spaces",
			args:    args{str: "abc"},
			want:    "",
			want1:   0,
			wantErr: true,
		}, {
			name:    "when the input has at least one space but no valid score at the end",
			args:    args{str: "abc def"},
			want:    "",
			want1:   0,
			wantErr: true,
		}, {
			name:    "when the input has at least one space and a valid score at the end",
			args:    args{str: "abc 1"},
			want:    "abc",
			want1:   1,
			wantErr: false,
		}, {
			name:    "when the input has more than one space and a valid score at the end",
			args:    args{str: "abc def 42"},
			want:    "abc def",
			want1:   42,
			wantErr: false,
		}, {
			name:    "should not care if there are non alpha characters in the team name",
			args:    args{str: "abc 54 def 42"},
			want:    "abc 54 def",
			want1:   42,
			wantErr: false,
		}, {
			name:    "should have an error if a non integer score is provided",
			args:    args{str: "abc def 42.1"},
			want:    "",
			want1:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parse.TeamScore(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("TeamScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TeamScore() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("TeamScore() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
