package rootcause

import (
	"time"

	"github.com/gin-gonic/gin"
)

// RootCauseCategory defines the category of a root cause
type RootCauseCategory string

const (
	CategoryPrompt      RootCauseCategory = "prompt"
	CategoryRAG         RootCauseCategory = "rag"
	CategoryModel       RootCauseCategory = "model"
	CategoryTool        RootCauseCategory = "tool"
	CategoryData        RootCauseCategory = "data"
	CategoryConfig      RootCauseCategory = "config"
)

// Analysis represents a root cause analysis result
type Analysis struct {
	ID           string             `json:"id"`
	ResourceID   string             `json:"resource_id"`
	TriggerEvent string             `json:"trigger_event"`
	Category     RootCauseCategory  `json:"category"`
	Confidence   float64            `json:"confidence"`
	Description  string             `json:"description"`
	Evidence     []string           `json:"evidence"`
	Suggestions  []Suggestion       `json:"suggestions"`
	TraceIDs     []string           `json:"trace_ids"`
	AnalyzedAt   time.Time          `json:"analyzed_at"`
}

// Suggestion represents an optimization suggestion
type Suggestion struct {
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Priority    int               `json:"priority"`
	Description string            `json:"description"`
	Action      string            `json:"action"`
	Params      map[string]string `json:"params,omitempty"`
	AutoApply   bool              `json:"auto_apply"`
}

// Analyzer performs root cause analysis
type Analyzer struct {
	db    interface{}
	cache interface{}
}

// NewAnalyzer creates a new root cause analyzer
func NewAnalyzer(db, cache interface{}) *Analyzer {
	return &Analyzer{db: db, cache: cache}
}

// Analyze triggers a root cause analysis
func (a *Analyzer) Analyze(c *gin.Context) {
	var req struct {
		ResourceID  string json:"resource_id" binding:"required"
		TriggerEvent string json:"trigger_event" binding:"required"
		TraceIDs    []string `json:"trace_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	analysis := &Analysis{
		ID:           "rca-" + time.Now().Format("20060102150405"),
		ResourceID:   req.ResourceID,
		TriggerEvent: req.TriggerEvent,
		Category:     CategoryPrompt,
		Confidence:   0.85,
		Description:  "Analysis in progress",
		Evidence:     []string{},
		TraceIDs:     req.TraceIDs,
		AnalyzedAt:   time.Now(),
	}

	c.JSON(200, analysis)
}

// Get returns a specific analysis
func (a *Analyzer) Get(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "analysis": nil})
}

// GetSuggestions returns optimization suggestions for an analysis
func (a *Analyzer) GetSuggestions(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id": id,
		"suggestions": []Suggestion{
			{
				ID:          "sug-1",
				Type:        "prompt_optimization",
				Priority:    1,
				Description: "Refine prompt template for better accuracy",
				Action:      "update_prompt",
				AutoApply:   false,
			},
		},
	})
}