package main

import (
	"path"
	"testing"

	"github.com/repometric/lhexec/analyze"
	"github.com/repometric/lhexec/api"
	"github.com/repometric/lhexec/extensions"
)

var (
	API         = api.API{}
	ProjectPath = ""
)

func TestAnalyzeCmd(t *testing.T) {
	analyzeResult := analyze.Run(analyze.CLIContext{
		Engine:      "csslint",
		Config:      path.Join(ProjectPath, ".linterhub.json"),
		Environment: "local",
		Project:     ProjectPath,
		File:        "style.css",
	})
	if len(analyzeResult.Results) == 0 {
		t.Error("Result is nil")
	}
	t.Log(extensions.ConverObjectToJSON(analyzeResult, false))
}
