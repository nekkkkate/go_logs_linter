package main

import (
	"loglinter/analyzers"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzers.Analyzer)
}
