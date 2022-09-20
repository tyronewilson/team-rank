package usecase_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"span-challenge/internal/usecase"
	"span-challenge/pkg/fixtures"
	"span-challenge/pkg/models"
	"span-challenge/pkg/util"
	"strings"
	"testing"
)

func TestStreamResults(t *testing.T) {
	input := fixtures.ProvidedExampleInput
	expectedLines := strings.Split(string(input), "\n")
	reader := bytes.NewReader(input)
	errs := util.NewErrorCollector()
	stream := usecase.StreamResults(errs, reader)
	results := make(models.ResultSet, 0)
	for result := range stream {
		results = append(results, result)
	}
	assert.Len(t, results, len(expectedLines))
}
