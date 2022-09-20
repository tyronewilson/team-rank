package serialize

import (
	"fmt"
	"span-challenge/pkg/models"
	"span-challenge/pkg/platform"
)

func TeamRanksAsCSV(l models.TeamRankList) []byte {
	var output []byte
	for _, team := range l {
		row := fmt.Sprintf("%d. %s, %d %s%s", team.Rank, team.TeamName, team.Points, pointsSuffix(team.Points), platform.LineSeparator)
		output = append(output, []byte(row)...)
	}
	return output
}

func pointsSuffix(points int) string {
	if points == 1 {
		return "pt"
	}
	return "pts"
}
