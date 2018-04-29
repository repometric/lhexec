package api

import (
	"github.com/repometric/lhexec/analyze"

	"github.com/jinzhu/copier"
	"github.com/repometric/lhexec/catalog"
)

// API struct describes jsonrpc api of lhexec
type API struct{}

// GetInstance method returns new instance of api
func GetInstance() *API {
	return new(API)
}

// GetEngine api method returns representation of single engione by name (for test only)
func (*API) GetEngine(name *string, result *catalog.Engine) error {
	copier.Copy(&result, catalog.Get(*name))
	return nil
}

// Analyze api method executes analyze on given Context
func (*API) Analyze(context *analyze.Context, result *[]analyze.Result) error {
	copier.Copy(&result, analyze.Run(*context))
	return nil
}
