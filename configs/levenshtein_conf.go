package configs

import (
	"strings"

	"github.com/vincbro/suddig/distance"
	"github.com/vincbro/suddig/score"
)

func Levenshtein() Config {
	return Config{
		MinScore:     0.8,
		StringFunc:   strings.ToLower,
		DistanceFunc: distance.LevenshteinDistance,
		ScoreFunc:    score.Linear,
		TokenFunc:    DefualtToken,
	}
}
