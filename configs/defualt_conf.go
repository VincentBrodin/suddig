package configs

import (
	"strings"

	"github.com/vincbro/suddig/distance"
	"github.com/vincbro/suddig/score"
)

func Defualt() Config {
	return Config{
		MinScore:     0.8,
		StringFunc:   DefualtString,
		DistanceFunc: DefualtDistance,
		ScoreFunc:    DefualtScore,
		TokenFunc:    DefualtToken,
	}
}

// Points to the distance/LevenshteinDistance function
func DefualtDistance(query, target string) float64 {
	return distance.LevenshteinDistance(query, target)
}

// Does nothing, just returns the given string
func DefualtString(input string) string {
	return input
}

// Linear function for the score
func DefualtScore(query, target string, dist float64) float64 {
	return score.Linear(query, target, dist)
}

// Splits tokens based on spaces
func DefualtToken(str string) []string {
	return strings.Split(str, " ")
}
