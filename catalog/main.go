package catalog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const hubFolder = "hub"

// Get function creates instance of engine
func Get(engine string) *Engine {
	var result Engine

	argsFile, e := ioutil.ReadFile(path.Join(hubFolder, engine, "args.json"))
	if e != nil {
		fmt.Printf("Catch error while reading args file: %v\n", e)
		os.Exit(1)
	}

	extrFile, e := ioutil.ReadFile(path.Join(hubFolder, engine, "extr.json"))
	if e != nil {
		fmt.Printf("Catch error while reading extr file: %v\n", e)
		os.Exit(1)
	}

	var args argsParsed

	json.Unmarshal(argsFile, &(args))

	v := make([]Argument, 0, len(args.Definitions.Arguments.Properties))

	for key, value := range args.Definitions.Arguments.Properties {
		argument := Argument{
			ID:          value.ID,
			Name:        key,
			Type:        value.Type,
			Description: value.Description,
			Value:       value.Default,
		}
		splited := strings.Split(value.ID, ":")
		if len(splited) > 1 {
			argument.Prefix = splited[0]
			argument.ID = splited[1]
		}
		v = append(v, argument)
	}

	result.Args = Args{
		ID:        args.ID,
		Name:      args.Name,
		Arguments: v,
	}

	json.Unmarshal(extrFile, &(result.Extr))

	return &result
}
