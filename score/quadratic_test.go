package score_test

import (
	"testing"

	"github.com/vincbro/suddig/score"
)

func TestQuadraticZero(t *testing.T) {
	score := score.Quadratic("", "aaaaaa", 6)
	t.Logf("Expected 0 got %f\n", score)
	if score != 0 {
		t.FailNow()
	}
}

func TestQuadraticHalf(t *testing.T) {
	// ratio = 1 - 3/6 = 0.5, squared => 0.25
	score := score.Quadratic("", "aaaaaa", 3)
	t.Logf("Expected 0.25 got %f\n", score)
	if score != 0.25 {
		t.FailNow()
	}
}

func TestQuadraticFull(t *testing.T) {
	score := score.Quadratic("", "aaaaaa", 0)
	t.Logf("Expected 1 got %f\n", score)
	if score != 1 {
		t.FailNow()
	}
}
