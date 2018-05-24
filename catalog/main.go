package catalog

import (
	"os"
	"path"
	"strings"

	"github.com/repometric/lhexec/extensions"
	"github.com/repometric/lhexec/models"
)

var (
	dir, _        = os.Getwd()
	hubFolderName = "hub"
	argsFileName  = "args.json"
	pipeFileName  = "pipe.json"
)

// Get function creates instance of engine
func Get(engine string) *Context {
	var (
		context        = Context{}
		engineArgsFile = models.Args{}
		engineFolder   = path.Join(dir, hubFolderName, engine)
	)

	extensions.GetObjectInFile(path.Join(engineFolder, argsFileName), &engineArgsFile)

	arguments := make([]Argument, 0, len(engineArgsFile.Definitions.Arguments.Properties))

	for key, value := range engineArgsFile.Definitions.Arguments.Properties {
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
		arguments = append(arguments, argument)
	}

	context.Args = Args{
		ID:         engineArgsFile.ID,
		Name:       engineArgsFile.Name,
		Arguments:  arguments,
		Delimeters: engineArgsFile.Delimeters,
	}

	extensions.GetObjectInFile(path.Join(engineFolder, pipeFileName), &(context.Extr))
	return &context
}
