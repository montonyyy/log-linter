package main

import (
	"log-linter/logcheck"

	"golang.org/x/tools/go/analysis"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{logcheck.Analyzer}, nil
}
