package catalog

// Args structure describes engine dependencies
type Args struct {
	ID        string     `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Arguments []Argument `json:"arguments,omitempty"`
}

// Argument structure describes single engine dependency
type Argument struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Prefix      string `json:"prefix,omitempty"`
	Description string `json:"description,omitempty"`
	Value       string `json:"value,omitempty"`
	Type        string `json:"type,omitempty"`
}

type argsParsed struct {
	// Engine id, case insensitive unique identifier, required only if the engine name is not unique
	ID string `json:"id,omitempty"`

	// Case sensitive official engine name, may not be unique
	Name string `json:"name"`

	Definitions definitionsParsed `json:"definitions"`
}

type definitionsParsed struct {
	Arguments argumentsParsed `json:"arguments"`
}

type argumentsParsed struct {
	Properties map[string]propertiesParsed `json:"properties"`
}

type propertiesParsed struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Default     string `json:"default,omitempty"`
}
