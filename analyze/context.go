package analyze

import "catalog"

type Context struct {
	Engine      []catalog.Engine
	Environment string
	Project		string
	File		string
	Folder      string
}