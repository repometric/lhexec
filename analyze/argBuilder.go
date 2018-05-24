package analyze

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/repometric/lhexec/catalog"
	"github.com/repometric/lhexec/execute"
	"github.com/repometric/lhexec/models"
)

var (
	dir, _        = os.Getwd()
	hubFolderPath = path.Join(dir, "hub")
)

// ArgBuilder methods create splice of arguments for engine execution
func ArgBuilder(context Context, engine models.Engine, engineSpec catalog.Context) []execute.Argument {
	var result []execute.Argument
	for _, argument := range engineSpec.Args.Arguments {
		if argument.Prefix == "linterhub" {
			switch argument.ID {
			case "path", "filename":
				analyzePath := context.Project.Folder
				if len(context.Project.File) > 0 {
					analyzePath = path.Join(analyzePath, context.Project.File)
				}
				result = append(result, execute.Argument{
					Key:   argument.Name,
					Value: analyzePath,
				})
				break
			}

		} else {
			if len(argument.Value) > 0 {
				result = append(result, execute.Argument{
					Key:   argument.ID,
					Value: argument.Value,
				})
			}
		}
	}

	for key, value := range engine.Args {
		for _, v := range engineSpec.Args.Arguments {
			if v.Prefix == "args" && v.ID == key {
				result = append(result, execute.Argument{
					Key:   key,
					Value: value,
				})
			}
		}
	}
	return result
}

func replaceResrvedName(data interface{}, engineName string) string {
	result := fmt.Sprintf("%v", data)
	result = strings.Replace(result, "{{hub}}", hubFolderPath, -1)
	result = strings.Replace(result, "{{engine}}", engineName, -1)

	return result
}
