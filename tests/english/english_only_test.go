package english

import (
	"loglinter/analyzers"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestEnglishOnly(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzers.Analyzer, "englishonly")
}
