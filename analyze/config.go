package analyze

import (
	"bufio"
	"encoding/json"
	"io"
	"os"

	"github.com/repometric/lhexec/extensions"
	"github.com/repometric/lhexec/models"
)

func getConfig(path string, stdin bool) models.Config {
	config := models.Config{}
	if extensions.PathValidate(path) {
		extensions.GetObjectInFile(path, &config)
	}

	if stdin {
		reader := bufio.NewReader(os.Stdin)
		var output []rune

		for {
			input, _, err := reader.ReadRune()
			if err != nil && err == io.EOF {
				break
			}
			output = append(output, input)
		}

		json.Unmarshal([]byte(string(output)), &config)
	}
	return config
}

func setConfig(config models.Config, context *Context) error {
	if config.Engines != nil {
		for _, item := range config.Engines {
			if item.Name == context.Engine.Name {
				context.Engine.Args = item.Args
				context.Engine.Active = item.Active
			}
		}
	}

	if config.Ingores != nil {
		context.Ignores = config.Ingores
	}

	return nil
}
