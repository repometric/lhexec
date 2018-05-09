package models

//Config - stucture of config file
type Config struct {
	Engines []Engine     `json:"engines"`
	Ingores []IgnoreRule `json:"ignores"`
}

// Engine struct describes information for engine execution
type Engine struct {
	Name        string            `json:"name"`
	Environment string            `json:"install"`
	Active      bool              `json:"active"`
	Args        map[string]string `json:"args"`
}

// IgnoreRule struct describes ignore rules
type IgnoreRule struct {
	Mask     string `json:"mask"`
	RuleID   string `json:"ruleId"`
	Line     int    `json:"line"`
	IgnoreID string `json:"ignoreId"`
}

// InitEngine creates new instance of Engine struct with default params
func InitEngine() Engine {
	return Engine{
		Args: map[string]string{},
	}
}
