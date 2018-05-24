package analyze

import (
	"encoding/json"

	"github.com/repometric/lhexec/catalog"
	"github.com/repometric/lhexec/execute"
	"github.com/repometric/lhexec/extensions"
	"github.com/repometric/lhexec/models"
)

// Run function runs analyze for current context
func Run(data CLIContext) Result {

	var (
		context = Context{
			Project: Project{
				Path:   data.Project,
				Folder: data.Folder,
				File:   data.File,
			},
			Engine: models.Engine{
				Name:        data.Engine,
				Environment: data.Environment,
				Args:        map[string]string{},
			},
		}
		engineResult = Result{
			Engine:  context.Engine.Name,
			Results: []FileResult{},
			Errors:  []PipeError{},
		}
		engineSpec = *catalog.Get(context.Engine.Name)
		buffer     string
	)

	setConfig(
		getConfig(data.Config, data.Stdin),
		&context)

	if !context.Engine.Active {
		extensions.ShowError("Engine `" + context.Engine.Name + "` is disable. Please, activate it.")
	}

	for id, pipe := range engineSpec.Extr.Pipeline {
		executeContext := execute.Context{
			Binary:           pipe.Cmd,
			Stdin:            buffer,
			SuccessCode:      pipe.Success,
			WorkingDirectory: context.Project.Path,
		}

		if pipe.Engine {
			executeContext.Args = ArgBuilder(context, context.Engine, engineSpec)
			executeContext.Delimeters = engineSpec.Args.Delimeters

			if pipe.Args != nil {
				for key, value := range pipe.Args.(map[string]interface{}) {

					find := false
					for _, arg := range executeContext.Args {
						if arg.Key == key {
							find = true
							continue
						}
					}

					if !find {
						executeContext.Args = append(executeContext.Args, execute.Argument{
							Key:   key,
							Value: replaceResrvedName(value, context.Engine.Name),
						})
					}

				}
			}

		} else {

			executeContext.Args = []execute.Argument{
				execute.Argument{
					Key:   "",
					Value: replaceResrvedName(pipe.Args, context.Engine.Name),
				},
			}
		}

		//fmt.Println("command: "+strconv.Itoa(id), executeContext.Binary, executeContext.Args)
		stdout, stderr := execute.Run(executeContext)

		if len(stderr) > 0 {
			engineResult.Errors = append(engineResult.Errors, PipeError{
				Message:    stderr,
				PipelineID: id,
			})
		} else {
			buffer = stdout
		}
	}
	json.Unmarshal([]byte(buffer), &(engineResult.Results))
	return engineResult
}
