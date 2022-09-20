package usecase_test

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"io"
	"span-challenge/internal/usecase"
	"span-challenge/pkg/models"
	"span-challenge/pkg/platform"
	"strings"
	"testing"
)

type badWriter struct {
}

func (b badWriter) String() string {
	return ""
}

func (b badWriter) Write(_ []byte) (n int, err error) {
	return 0, errors.New("no reason, me am just a bad writer")
}

func TestWriteRankingsCSV(t *testing.T) {
	type args struct {
		list models.TeamRankList
		w    io.Writer
	}
	tests := []struct {
		name    string
		args    args
		wantW   []byte
		wantErr bool
	}{
		{
			name:    "with nil list",
			args:    args{w: bytes.NewBuffer(nil)},
			wantW:   nil,
			wantErr: false,
		}, {
			name:    "with empty list",
			args:    args{list: models.TeamRankList{}, w: bytes.NewBuffer([]byte{})},
			wantW:   []byte{},
			wantErr: false,
		}, {
			name:    "with non-empty list",
			args:    args{list: models.TeamRankList{{TeamName: "Team A", Points: 1, Rank: 1}}, w: bytes.NewBuffer([]byte{})},
			wantW:   []byte("1. Team A, 1 pt" + platform.LineSeparator),
			wantErr: false,
		}, {
			name: "with multiple teams",
			args: args{
				list: models.TeamRankList{
					{TeamName: "Team A", Points: 2, Rank: 1},
					{TeamName: "Team B", Points: 2, Rank: 1},
					{TeamName: "Team C", Points: 0, Rank: 3},
				},
				w: bytes.NewBuffer([]byte{}),
			},
			wantW:   []byte(strings.ReplaceAll("1. Team A, 2 pts\n1. Team B, 2 pts\n3. Team C, 0 pts\n", `\n`, platform.LineSeparator)),
			wantErr: false,
		}, {
			name:    "with a bad writer",
			args:    args{list: models.TeamRankList{{TeamName: "Team A", Points: 1}}, w: &badWriter{}},
			wantW:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := tt.args.w.(fmt.Stringer)
			err := usecase.WriteRankingsCSV(tt.args.list, tt.args.w)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equalf(t, string(tt.wantW), w.String(), "WriteRankingsCSV(%v, %v)", tt.args.list, w)
		})
	}
}
