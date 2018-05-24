package execute

// Context describes params necessary for command execution
type Context struct {
	Binary           string
	WorkingDirectory string
	Args             []Argument
	Stdin            string
	SuccessCode      int
	Delimeters       string
}

// Argument struct describes single argument
type Argument struct {
	Key   string
	Value string
}
