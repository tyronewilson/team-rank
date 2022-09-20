package usecase_test

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"spanchallenge/internal/usecase"
	"spanchallenge/pkg/fixtures"
	"spanchallenge/pkg/models"
	"spanchallenge/pkg/platform"
	"spanchallenge/pkg/util"
	"strings"
	"testing"
)

func TestStreamResults(t *testing.T) {
	input := fixtures.ProvidedExampleInput
	expectedLines := strings.Split(string(input), platform.LineSeparator)
	reader := bytes.NewReader(input)
	errs := util.NewErrorCollector()
	stream := usecase.StreamResults(errs, reader)
	results := make(models.ResultSet, 0)
	for result := range stream {
		results = append(results, result)
	}
	assert.Len(t, results, len(expectedLines))
}

func BenchmarkStreamResults10(b *testing.B) {
	benchStreamResults(b, 10)
}

func BenchmarkStreamResults10000(b *testing.B) {
	benchStreamResults(b, 10000)
}

func benchStreamResults(b *testing.B, inputLimit int) {
	input := getInputSet(inputLimit)
	for i := 0; i < b.N; i++ {
		reader := bytes.NewReader(input)
		errs := util.NewErrorCollector()
		stream := usecase.StreamResults(errs, reader)
		for range stream {
		}
	}
}

func getInputSet(limit int) []byte {
	result := make([]byte, 0)
	for i := 0; i < limit; i++ {
		one := fmt.Sprintf("%s%d %d", "Team A", i, rand.Intn(10))
		two := fmt.Sprintf("%s%d %d", "Team B", i, rand.Intn(10))
		row := strings.Join([]string{one, two}, ",") + platform.LineSeparator
		result = append(result, []byte(row)...)
	}
	return result
}
