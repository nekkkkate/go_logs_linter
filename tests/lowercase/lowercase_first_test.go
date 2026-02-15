package lowercase

import (
	"loglinter/analyzers"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestLowerCaseFirst(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzers.Analyzer, "lowercasefirst")
}
