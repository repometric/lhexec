package catalog

// Extr structure describes extra engine data
type Extr struct {
	// Engine id, case insensitive unique identifier, required only if the engine name is not unique
	ID string `json:"id,omitempty"`

	// Case sensitive official engine name, may not be unique
	Name string `json:"name"`

	// Ordered chain of elements, containing command-line options, which allows to get a specific output format after executing
	Pipeline []PipelineChunk `json:"pipeline"`

	// Object, which stores execution information for engine
	Environment Environment `json:"environment"`
}

// PipelineChunk is a single element in pipeline chain
type PipelineChunk struct {
	// The executable string. Stdout of previous execution must be passed as stdin to next one
	Cmd string `json:"cmd,omitempty"`

	// Success exit code, if not set then it's assumed equal to 0
	Success int `json:"success,omitempty"`

	// Is the command-line option an engine executable name
	Engine bool `json:"engine,omitempty"`
}

// Environment struct stores execution information for engine
type Environment struct {
	// Whether filename masks are supported by engine, default is true
	Masks bool `json:"cmd,omitempty"`

	// The engine version, equal to the output of version command, it may be used for debugging purposes
	Version string `json:"version,omitempty"`
}

// InitEnvironment creates new instance of Environment struct with default params
func InitEnvironment() Environment {
	return Environment{
		Masks: true,
	}
}
