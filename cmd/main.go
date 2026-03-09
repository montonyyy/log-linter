package main

import (
	"log-linter/logcheck"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(logcheck.Analyzer)
}
