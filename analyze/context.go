package analyze

// Context struct describes all information for execution
type Context struct {
	Engine       []Engine
	Project      string
	File         string
	Folder       string
	Stdin        bool
	StdinContent string
	Ignores      []IgnoreRule
}

// Engine struct describes information for engine execution
type Engine struct {
	Name        string
	Environment string
	Args        map[string]string
}

// IgnoreRule struct describes ignore rules
type IgnoreRule struct {
	Mask   string
	RuleID string
	Line   int
}
