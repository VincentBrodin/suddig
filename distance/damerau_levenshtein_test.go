package distance_test

import (
	"testing"

	"github.com/vincbro/suddig/distance"
)

func TestDamerauLevenshteinTransposition(t *testing.T) {
	dist := distance.DamerauLevenshteinDistance("ac", "ca")
	t.Logf("Expected 1 got %f\n", dist)
	if dist != 1 {
		t.FailNow()
	}
}

func TestDamerauLevenshteinTransposition2(t *testing.T) {
	dist := distance.DamerauLevenshteinDistance("abcd", "acbd")
	t.Logf("Expected 1 got %f\n", dist)
	if dist != 1 {
		t.FailNow()
	}
}

func TestDamerauLevenshteinTransposition3(t *testing.T) {
	dist := distance.DamerauLevenshteinDistance("abcdef", "abdcef")
	t.Logf("Expected 1 got %f\n", dist)
	if dist != 1 {
		t.FailNow()
	}
}

func TestDamerauLevenshteinTransposition4(t *testing.T) {
	dist := distance.DamerauLevenshteinDistance("converse", "conevrse")
	t.Logf("Expected 1 got %f\n", dist)
	if dist != 1 {
		t.FailNow()
	}
}

func TestDamerauLevenshteinSameLength(t *testing.T) {
	dist := distance.DamerauLevenshteinDistance("book", "back")
	t.Logf("Expected 2 got %f\n", dist)
	if dist != 2 {
		t.FailNow()
	}
}

func TestDamerauLevenshteinTwoWords(t *testing.T) {
	dist := distance.DamerauLevenshteinDistance("hello", "hello world")
	t.Logf("Expected 6 got %f\n", dist)
	if dist != 6 {
		t.FailNow()
	}
}

func TestDamerauLevenshteinEmptyStrings(t *testing.T) {
	dist := distance.DamerauLevenshteinDistance("", "")
	t.Logf("Expected 0 got %f\n", dist)
	if dist != 0 {
		t.FailNow()
	}
}

func TestDamerauLevenshteinBigStrings(t *testing.T) {
	a := string(make([]byte, 1000))
	b := string(make([]byte, 1000))
	dist := distance.DamerauLevenshteinDistance(a, b)
	t.Logf("Expected 0 got %f\n", dist)
	if dist != 0 {
		t.FailNow()
	}
}

func TestDamerauLevenshteinRepeating(t *testing.T) {
	dist := distance.DamerauLevenshteinDistance("aabaa", "aaaa")
	t.Logf("Expected 1 got %f\n", dist)
	if dist != 1 {
		t.FailNow()
	}
}

func TestDamerauLevenshteinNonAscii(t *testing.T) {
	dist := distance.DamerauLevenshteinDistance("båt", "bat")
	t.Logf("Expected 1 got %f\n", dist)
	if dist != 1 {
		t.FailNow()
	}
}
