package main

import (
	"loglinter/analyzers"

	"golang.org/x/tools/go/analysis"
)

type AnalyzerPlugin struct{}

func (*AnalyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{analyzers.Analyzer}
}

func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzers.Analyzer}, nil
}
