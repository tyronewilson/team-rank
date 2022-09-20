package serialize

import (
	"fmt"
	"spanchallenge/pkg/models"
	"spanchallenge/pkg/platform"
)

// TeamRanksAsCSV takes a list of team ranks and returns a CSV string
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
