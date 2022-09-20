package serialize

import (
	"fmt"
	"span-challenge/pkg/models"
)

func TeamRanksAsCSV(l models.TeamRankList) []byte {
	var output []byte
	for i, rank := range l {
		row := fmt.Sprintf("%d. %s, %d pts", i+1, rank.TeamName, rank.Points)
		row += "\n"
		output = append(output, []byte(row)...)
	}
	return output
}
