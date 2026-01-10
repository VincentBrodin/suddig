package distance_test

import (
	"testing"

	"github.com/vincbro/suddig/distance"
)

func TestLevenshteinSameLength(t *testing.T) {
	dist := distance.LevenshteinDistance("book", "back")
	t.Logf("Expected 2 got %f\n", dist)
	if dist != 2 {
		t.FailNow()
	}
}

func TestLevenshteinTwoWords(t *testing.T) {
	dist := distance.LevenshteinDistance("hello", "hello world")
	t.Logf("Expected 6 got %f\n", dist)
	if dist != 6 {
		t.FailNow()
	}
}

func TestLevenshteinEmptyStrings(t *testing.T) {
	dist := distance.LevenshteinDistance("", "")
	t.Logf("Expected 0 got %f\n", dist)
	if dist != 0 {
		t.FailNow()
	}
}

func TestLevenshteinBigStrings(t *testing.T) {
	a := string(make([]byte, 1000))
	b := string(make([]byte, 1000))
	dist := distance.LevenshteinDistance(a, b)
	t.Logf("Expected 0 got %f\n", dist)
	if dist != 0 {
		t.FailNow()
	}
}

func TestLevenshteinRepeating(t *testing.T) {
	dist := distance.LevenshteinDistance("aabaa", "aaaa")
	t.Logf("Expected 1 got %f\n", dist)
	if dist != 1 {
		t.FailNow()
	}
}

func TestLevenshteinNonAscii(t *testing.T) {
	dist := distance.LevenshteinDistance("båt", "bat")
	t.Logf("Expected 1 got %f\n", dist)
	if dist != 1 {
		t.FailNow()
	}
}
