## suddig {adj.} /²s'ɵdːɪg/

*(Swedish for "fuzzy")*

A fast, flexible, and modular fuzzy finder library for Go. Built for both quick, one-line matches and deep customization, suddig provides simple functions (`Match`, `Distance`, `Score`, `FindMatches`, `RankMatches`) for everyday use, and a fully configurable `matcher` package for advanced scenarios.

---

## Features

* **Modular Architecture**: Swap or extend normalization, distance, and scoring components.
* **Simple API**: Call `suddig.Match`, `suddig.Distance`, `suddig.Score`, `suddig.FindMatches` & `suddig.RankMatches` out of the box.
* **Advanced Configuration**: Instantiate a `matcher.Matcher` with custom `configs.Config` to tweak behavior.
* **Parallel Processing**: Leverage all available CPU cores to perform matching and ranking in parallel—ideal for large datasets, with near-linear speedups and minimal overhead.
* **Mutliple Distance Algorithms**: Both `Levenshtein` & `Damerau Levenshtein` are supported out of the box, and you can easily add more.

---

## Installation

```bash
go get github.com/vincbro/suddig
```

---

## Usage

```go
package main

import (
	"fmt"
	"os"

	"github.com/vincbro/suddig"
)

func main() {
	args := os.Args[1:]
	argc := len(args)

	if argc < 2 {
		fmt.Println("Need 2 arguments")
		os.Exit(1)
	}
	s1, s2 := args[0], args[1]

	match := suddig.Match(s1, s2)

	if match {
		fmt.Printf("%s and %s match!\n", s1, s2)
	} else {
		fmt.Printf("%s and %s do not match\n", s1, s2)
	}
}
```

