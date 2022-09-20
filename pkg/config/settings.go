package config

import (
	"fmt"
)

// RankingRules encapsulates the rules for ranking teams based on match results
type RankingRules struct {
	WinPoints  int `yaml:"win_points" env:"SPAN_WIN_POINTS" env-default:"3"`
	DrawPoints int `yaml:"draw_points" env:"SPAN_DRAW_POINTS" env-default:"1"`
	LossPoints int `yaml:"loss_points" env:"SPAN_LOSS_POINTS" env-default:"0"`
}

// Validate validates the rules configuration.
func (r RankingRules) Validate() error {
	if r.WinPoints < 0 {
		return fmt.Errorf("win points must be greater than or equal to 0")
	}
	if r.DrawPoints < 0 {
		return fmt.Errorf("draw points must be greater than or equal to 0")
	}
	if r.LossPoints < 0 {
		return fmt.Errorf("loss points must be greater than or equal to 0")
	}
	return nil
}

// Validatable is an interface that can be implemented by any struct that needs to be validated
type Validatable interface {
	Validate() error
}
