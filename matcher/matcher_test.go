package matcher_test

import (
	"sort"
	"testing"

	"github.com/vincbro/suddig/configs"
	"github.com/vincbro/suddig/matcher"
)

func TestParalellelMatcher(t *testing.T) {
	m := matcher.New(configs.DamerauLevenshtein())
	items := []string{"Test", "Hello", "Cool", "World", "Golang", "House", "Boat", "Other stuff"}
	results := m.ParallelMatch("test", items)
	t.Log(results)
	t.Log(items)
}

func TestParalellelRanker(t *testing.T) {
	m := matcher.New(configs.DamerauLevenshtein())
	items := []string{"Test", "Hello", "Cool", "World", "Golang", "House", "Boat", "Other stuff"}
	results := m.ParallelRank("test", items)
	t.Log(results)
	t.Log(items)
}

func TestTokenMatcher(t *testing.T) {
	m := matcher.New(configs.DamerauLevenshtein())
	items := []string{"Hello", "Hello World", "Hello House", "World", "Best World", "House Top", "Boat Low", "Other stuff"}
	results := m.RankMatchesTokenized("Hello World", m.TokenizeAll(items))
	sort.Slice(items, func(i, j int) bool {
		return results[i] > results[j]
	})
	t.Log(items)
}

func TestTokenMatcherParalellel(t *testing.T) {
	m := matcher.New(configs.DamerauLevenshtein())
	items := []string{"Hello", "Hello World", "Hello House", "World", "Best World", "House Top", "Boat Low", "Other stuff"}
	results := m.ParallelRankTokenized("Hello World", m.TokenizeAll(items))
	sort.Slice(items, func(i, j int) bool {
		return results[i] > results[j]
	})
	t.Log(items)
}

// func TestParalelleMatcher(t *testing.T) {
// 	m := matcher.New(configs.Levenshtein())
// 	url := "https://sample-files.com/downloads/documents/txt/long-doc.txt"
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		t.FailNow()
// 	}
// 	b, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		t.FailNow()
// 	}
// 	str := string(b)
// 	in := strings.Split(str, " ")
// 	for range 10 {
// 		in = append(in, in...)
// 	}
//
// 	t.Logf("Size: %d\n", len(in))
//
// 	start := time.Now()
// 	out := m.ParallelMatch("lorem", in)
// 	t.Logf("Found %d matches\n", len(out))
// 	elapsed := time.Since(start)
// 	t.Logf("Paralelle took %s\n", elapsed)
//
// 	start = time.Now()
// 	out = m.FindMatches("lorem", in)
// 	t.Logf("Found %d matches\n", len(out))
// 	elapsed = time.Since(start)
// 	t.Logf("Normal took %s\n", elapsed)
// }
//
//
// func TestParalelleRanker(t *testing.T) {
// 	m := matcher.New(configs.DamerauLevenshtein())
// 	url := "https://sample-files.com/downloads/documents/txt/long-doc.txt"
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		t.FailNow()
// 	}
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		t.FailNow()
// 	}
// 	str := string(body)
// 	in := strings.Split(str, " ")
// 	for range 10 {
// 		in = append(in, in...)
// 	}
//
// 	t.Logf("Size: %d\n", len(in))
//
// 	start := time.Now()
// 	a := m.ParallelRank("l", in)
// 	t.Log(len(a))
// 	elapsed := time.Since(start)
// 	t.Logf("Paralelle took %s\n", elapsed)
//
// 	start = time.Now()
// 	b := m.RankMatches("lorem", in)
// 	t.Log(len(b))
// 	elapsed = time.Since(start)
// 	t.Logf("Normal took %s\n", elapsed)
// }
