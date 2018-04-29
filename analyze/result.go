package analyze

type Result struct {
	// Engine name
	Engine  string       `json:"name"`
	Results []FileResult `json:"results"`
	Errors  []PipeError  `json:"errors"`
}

type PipeError struct {
	// The description of the error
	Message string `json:"message"`

	// The id or crashed pipeline
	PipelineID int `json:"pipelineId"`
}

// FileResult struct
type FileResult struct {
	// The path relative to the analysis root
	Path string `json:"path"`

	// List of messages in the path
	Messages []Message `json:"messages"`
}

// Message struct describes error and where is he located
type Message struct {
	// The short description of the message
	Message string `json:"message"`

	// The explanatory text of the message
	Description string `json:"description,omitempty"`

	// The severity of the message: “verbose”, “hint”, “information”, “warning” or “error”
	Severity string `json:"severity"`

	// The line where the message is located
	Line int `json:"line"`

	// The end line where the message is located (the same as line by default)
	LineEnd int `json:"lineEnd,omitempty"`

	// The column where the message is located
	Column int `json:"column,omitempty"`

	// The end column where the message is located
	ColumnEnd int `json:"columnEnd,omitempty"`

	// The id of the rule that produced the message
	RuleID string `json:"ruleId,omitempty"`

	// The name of the rule that produced the message
	RuleName string `json:"ruleName,omitempty"`

	// The namespace of the rule that produced the message
	RuleNs string `json:"ruleNs,omitempty"`
}
