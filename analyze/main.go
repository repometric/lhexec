package analyze

import (
	"encoding/json"

	"github.com/repometric/lhexec/catalog"
	"github.com/repometric/lhexec/execute"
)

// Run function runs analyze for current context
func Run(context Context) []Result { // TODO: now its works only with local execution
	var result []Result
	for _, engine := range context.Engine {
		resultEngine := Result{
			Engine:  engine.Name,
			Results: []FileResult{},
			Errors:  []PipeError{},
		}
		var engineSpec = *catalog.Get(engine.Name)
		var buffer string
		for id, pipe := range engineSpec.Extr.Pipeline {
			executeContext := execute.Context{
				Binary:      pipe.Cmd,
				Stdin:       buffer,
				SuccessCode: pipe.Success,
			}
			if pipe.Engine {
				executeContext.WorkingDirectory = context.Project
				executeContext.Args = ArgBuilder(context, engine, engineSpec)
			}
			stdout, stderr := execute.Run(executeContext)
			if len(stderr) > 0 {
				resultEngine.Errors = append(resultEngine.Errors, PipeError{
					Message:    stderr,
					PipelineID: id,
				})
			} else {
				buffer = stdout
			}
		}
		json.Unmarshal([]byte(buffer), &(resultEngine.Results))
		result = append(result, resultEngine)
	}

	return result
}
