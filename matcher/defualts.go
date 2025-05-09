package matcher

import (
	"github.com/VincentBrodin/suddig/distance"
	"github.com/VincentBrodin/suddig/score"
)

// Points to the distance/LevenshteinDistance function
func DefualtDistance(query, target string) int32 {
	return distance.LevenshteinDistance(query, target)
}

// Does nothing, just returns the given string
func DefualtString(input string) string {
	return input
}

// Linear function for the score
func DefualtScore(query, target string, dist int32) float64 {
	return score.Linear(query, target, dist)
}
