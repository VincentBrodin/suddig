package suddig

import (
	"github.com/vincbro/suddig/configs"
	"github.com/vincbro/suddig/matcher"
)

// Match computes a normalized similarity score between two strings.
// It returns true if the similarity score is 80% or above.
// Match uses the base Distance function.
func Match(query, target string) bool {
	maxDist := max(len(query), len(target))
	percent := 1.0 - float64(Distance(query, target))/float64(maxDist)
	return percent >= 0.8
}

// Returns a normalized linear score based on distance.
// See Distance for the raw difference between the strings.
func Score(query, target string) float64 {
	m := matcher.New(configs.Defualt())
	return m.Score(query, target)
}

// Returns the distance based on the levenshtein distance.
// The Levenshtein algorithm returns a number based on how many edits is needed for the query to match the target.
// O(m*n) where m & n are the length of the input strigns.
// See Score for the percental difference between the strings.
func Distance(query, target string) float64 {
	m := matcher.New(configs.Defualt())
	return m.Distance(query, target)
}

// FindMatches returns all strings in targets whose similarity
// to query is ≥ 80%.
func FindMatches(query string, targets []string) []string {
	m := matcher.New(configs.Defualt())
	return m.ParallelMatch(query, targets)
}

// RankMatches returns all the similarity score for all strings in the given array.
// The returned array matches the targets array.
func RankMatches(query string, targets []string) []float64 {
	m := matcher.New(configs.Defualt())
	return m.ParallelRank(query, targets)

}
