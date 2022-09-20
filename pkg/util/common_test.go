package util_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"span-challenge/pkg/util"
	"testing"
)

func TestMaybePanic(t *testing.T) {
	assert.NotPanics(t, func() {
		util.MaybePanic(nil)
	})
	assert.Panics(t, func() {
		util.MaybePanic(fmt.Errorf("test"))
	})
}

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
			name:  "with empty string",
			args:  args{str: ""},
			want:  "",
			want1: "",
		}, {
			name:  "with single word",
			args:  args{str: "test"},
			want:  "",
			want1: "",
		}, {
			name:  "with two words",
			args:  args{str: "one two"},
			want:  "one",
			want1: "two",
		}, {
			name:  "with three words",
			args:  args{str: "one two three"},
			want:  "one two",
			want1: "three",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := util.SplitOnLastSpace(tt.args.str)
			assert.Equalf(t, tt.want, got, "SplitOnLastSpace(%v)", tt.args.str)
			assert.Equalf(t, tt.want1, got1, "SplitOnLastSpace(%v)", tt.args.str)
		})
	}
}
