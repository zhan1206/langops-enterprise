package model

// Provider defines supported LLM providers
type Provider string

const (
	ProviderOpenAI    Provider = "openai"
	ProviderAnthropic Provider = "anthropic"
	ProviderGemini    Provider = "gemini"
	ProviderQwen      Provider = "qwen"
	ProviderDeepSeek  Provider = "deepseek"
	ProviderLocal     Provider = "local"
	ProviderCustom    Provider = "custom"
)

// ModelConfig defines a model's configuration
type ModelConfig struct {
	ID         string   json:"id"
	Provider   Provider json:"provider"
	ModelName  string   json:"model_name"
	Endpoint   string   json:"endpoint"
	APIKey     string   json:"api_key,omitempty"
	MaxTokens  int      json:"max_tokens"
	CostInput  float64  json:"cost_per_1k_input"
	CostOutput float64  json:"cost_per_1k_output"
}

// Router routes requests to the appropriate model
type Router struct {
	config interface{}
	logger interface{}
	models map[string]*ModelConfig
}

// NewRouter creates a new model router
func NewRouter(config, logger interface{}) *Router {
	return &Router{
		config: config,
		logger: logger,
		models: make(map[string]*ModelConfig),
	}
}

// Register registers a model configuration
func (r *Router) Register(cfg *ModelConfig) {
	r.models[cfg.ID] = cfg
}

// Route selects the best model for a request
func (r *Router) Route(req map[string]interface{}) (*ModelConfig, error) {
	// Default: route to first available model
	for _, m := range r.models {
		return m, nil
	}
	return nil, nil
}