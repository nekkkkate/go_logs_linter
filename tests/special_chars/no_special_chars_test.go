package special_chars

import (
	"loglinter/analyzers"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestNoSpecialChars(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzers.Analyzer, "nospecialchars")
}
