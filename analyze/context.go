package analyze

import "github.com/repometric/lhexec/models"

//CLIContext - s
type CLIContext struct {
	Engine      string
	Environment string
	Project     string
	File        string
	Folder      string
	Stdin       bool
	Config      string
}

// Context struct describes all information for execution
type Context struct {
	Engine  models.Engine
	Project Project
	Ignores []models.IgnoreRule
}

type Project struct {
	Path   string
	Folder string
	File   string
}
