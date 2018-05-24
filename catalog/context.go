package catalog

import "github.com/repometric/lhexec/models"

// Context structure
type Context struct {
	Extr models.Extr `json:"extr"`
	Args Args        `json:"args"`
}

// Args structure describes engine dependencies
type Args struct {
	ID         string     `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	Delimeters string     `json:"delimeters,omitempty"`
	Arguments  []Argument `json:"arguments,omitempty"`
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
