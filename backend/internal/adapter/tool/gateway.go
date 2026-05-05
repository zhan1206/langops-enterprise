package tool

// ToolType defines the type of external tool
type ToolType string

const (
	ToolTypeMCP      ToolType = "mcp"
	ToolTypeOpenAPI  ToolType = "openapi"
	ToolTypeDatabase ToolType = "database"
	ToolTypeInternal ToolType = "internal"
)

// ToolConfig defines a tool's configuration
type ToolConfig struct {
	ID          string   json:"id"
	Name        string   json:"name"
	Type        ToolType json:"type"
	Endpoint    string   json:"endpoint"
	Description string   json:"description"
	Enabled     bool     json:"enabled"
}

// Gateway provides unified tool access
type Gateway struct {
	config interface{}
	logger interface{}
	tools  map[string]*ToolConfig
}

// NewGateway creates a new tool gateway
func NewGateway(config, logger interface{}) *Gateway {
	return &Gateway{
		config: config,
		logger: logger,
		tools:  make(map[string]*ToolConfig),
	}
}

// Register registers a tool
func (g *Gateway) Register(cfg *ToolConfig) {
	g.tools[cfg.ID] = cfg
}

// Execute executes a tool call
func (g *Gateway) Execute(toolID string, params map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	result["tool_id"] = toolID
	result["status"] = "executed"
	return result, nil
}