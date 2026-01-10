package score_test

import (
	"testing"

	"github.com/vincbro/suddig/score"
)

func TestLinearZero(t *testing.T) {
	score := score.Linear("", "aaaaaa", 6)
	t.Logf("Expected 0 got %f\n", score)
	if score != 0 {
		t.FailNow()
	}
}

func TestLinearHalf(t *testing.T) {
	score := score.Linear("", "aaaaaa", 3)
	t.Logf("Expected 0.5 got %f\n", score)
	if score != 0.5 {
		t.FailNow()
	}
}

func TestLinearFull(t *testing.T) {
	score := score.Linear("", "aaaaaa", 0)
	t.Logf("Expected 1 got %f\n", score)
	if score != 1 {
		t.FailNow()
	}
}
