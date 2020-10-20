package contexts

// Context : context parameter that gets passed arround by creator functions
type Context struct {
	ContextType string
	GetValue    map[string]string
}

// NewContext : abstraction generator function for Context struct
func NewContext(contextType string) *Context {
	contextMap := make(map[string]string)
	return &Context{contextType, contextMap}
}
