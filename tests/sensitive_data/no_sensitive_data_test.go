package sensitive_data

import (
	"loglinter/analyzers"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestNoSensitiveData(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzers.Analyzer, "nosensitivedata")
}
