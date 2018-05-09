package models

//argsParsed asdasd
type Args struct {
	// Engine id, case insensitive unique identifier, required only if the engine name is not unique
	ID string `json:"id,omitempty"`

	// Case sensitive official engine name, may not be unique
	Name        string      `json:"name"`
	Delimeters  string      `json:"delimeters,omitempty"`
	Definitions Definitions `json:"definitions"`
}

type Definitions struct {
	Arguments Arguments `json:"arguments"`
}

type Arguments struct {
	Properties map[string]Properties `json:"properties"`
}

type Properties struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Default     string `json:"default,omitempty"`
}
