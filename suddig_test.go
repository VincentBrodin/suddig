package suddig_test

import (
	"testing"

	"github.com/VincentBrodin/suddig"
)

func TestSameLength(t *testing.T) {
	dist := suddig.Distance("book", "back")
	t.Logf("Expected 2 got %d\n", dist)
	if dist != 2 {
		t.FailNow()
	}
}

func TestTwoWords(t *testing.T) {
	dist := suddig.Distance("hello", "hello world")
	t.Logf("Expected 6 got %d\n", dist)
	if dist != 6 {
		t.FailNow()
	}
}

func TestEmptyStrings(t *testing.T) {
	dist := suddig.Distance("", "")
	t.Logf("Expected 0 got %d\n", dist)
	if dist != 0 {
		t.FailNow()
	}
}

func TestBigStrings(t *testing.T) {
	a := string(make([]byte, 1000))
	b := string(make([]byte, 1000))
	dist := suddig.Distance(a, b)
	t.Logf("Expected 0 got %d\n", dist)
	if dist != 0 {
		t.FailNow()
	}
}

func TestRepeating(t *testing.T) {
	dist := suddig.Distance("aabaa", "aaaa")
	t.Logf("Expected 1 got %d\n", dist)
	if dist != 1 {
		t.FailNow()
	}
}

func TestMatch(t *testing.T) {
	if suddig.Match("book", "back") {
		t.FailNow()
	}

	if suddig.Match("book", "boat") {
		t.FailNow()
	}

	if !suddig.Match("house", "bouse") {
		t.FailNow()
	}

}
