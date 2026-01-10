package configs

import (
	"strings"

	"github.com/vincbro/suddig/distance"
	"github.com/vincbro/suddig/score"
)

func DamerauLevenshtein() Config {
	return Config{
		MinScore:     0.8,
		StringFunc:   strings.ToLower,
		DistanceFunc: distance.DamerauLevenshteinDistance,
		ScoreFunc:    score.Linear,
		TokenFunc:    DefualtToken,
	}
}
