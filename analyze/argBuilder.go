package analyze

import (
	"os"
	"path"

	"github.com/repometric/lhexec/catalog"
	"github.com/repometric/lhexec/execute"
)

// ArgBuilder methods create splice of arguments for engine execution
func ArgBuilder(context Context, engine Engine, engineSpec catalog.Engine) []execute.Argument {
	var result []execute.Argument
	for _, argument := range engineSpec.Args.Arguments {
		if argument.Prefix == "linterhub" {
			switch argument.ID {
			case "path", "filename":
				analyzePath := context.Folder
				if len(context.File) > 0 {
					analyzePath = path.Join(analyzePath, context.File)
				}
				result = append(result, execute.Argument{
					Key:   argument.Name,
					Value: analyzePath,
				})
				break
			case "reporter":
				dir, _ := os.Getwd()
				reporterPath := path.Join(dir, "hub", engine.Name, argument.Value)
				result = append(result, execute.Argument{
					Key:   argument.Name,
					Value: reporterPath,
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
