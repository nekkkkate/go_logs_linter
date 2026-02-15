package analyzers

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "loglinter",
	Doc:  "lints log statements in Go code",
	Run:  run,
}
