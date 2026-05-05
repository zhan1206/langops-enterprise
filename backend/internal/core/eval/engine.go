package eval

import (
	"math"
	"time"

	"github.com/gin-gonic/gin"
)

// Dimension represents an evaluation dimension
type Dimension string

const (
	DimensionAccuracy      Dimension = "accuracy"
	DimensionRelevance     Dimension = "relevance"
	DimensionFactuality    Dimension = "factuality"
	DimensionCompliance    Dimension = "compliance"
	DimensionFluency       Dimension = "fluency"
	DimensionCoherence     Dimension = "coherence"
	DimensionHarmfulness   Dimension = "harmfulness"
	DimensionEfficiency    Dimension = "efficiency"
	DimensionContextuality Dimension = "contextuality"
	DimensionCompleteness  Dimension = "completeness"
)

// EvalResult represents a single evaluation result
type EvalResult struct {
	Dimension Dimension  json:"dimension"
	Score     float64    json:"score"
	Weight    float64    json:"weight"
	Details   string     json:"details,omitempty"
	Passed    bool       json:"passed"
}

// EvalReport represents a complete evaluation report
type EvalReport struct {
	ID          string       json:"id"
	ResourceID  string       json:"resource_id"
	Version     int          json:"version"
	Results     []EvalResult json:"results"
	TotalScore  float64      json:"total_score"
	PassRate    float64      json:"pass_rate"
	EvalModel   string       json:"eval_model"
	CompletedAt time.Time    json:"completed_at"
}

// EvalConfig defines evaluation configuration
type EvalConfig struct {
	Dimensions []DimensionConfig json:"dimensions"
	EvalModel  string            json:"eval_model"
	Threshold  float64           json:"threshold"
}

// DimensionConfig defines a dimension's evaluation config
type DimensionConfig struct {
	Dimension Dimension json:"dimension"
	Weight    float64   json:"weight"
	Threshold float64   json:"threshold"
}

// Engine is the multi-dimensional evaluation engine
type Engine struct {
	db    interface{}
	cache interface{}
}

// NewEngine creates a new evaluation engine
func NewEngine(db, cache interface{}) *Engine {
	return &Engine{db: db, cache: cache}
}

// List returns all evaluations
func (e *Engine) List(c *gin.Context) {
	c.JSON(200, gin.H{
		"items": []EvalReport{},
		"total": 0,
	})
}

// Create creates a new evaluation task
func (e *Engine) Create(c *gin.Context) {
	var cfg EvalConfig
	if err := c.ShouldBindJSON(&cfg); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Apply default dimensions if none specified
	if len(cfg.Dimensions) == 0 {
		cfg.Dimensions = defaultDimensions()
	}

	c.JSON(201, gin.H{
		"id":     "eval-" + time.Now().Format("20060102150405"),
		"config": cfg,
		"status": "created",
	})
}

// Run executes an evaluation
func (e *Engine) Run(c *gin.Context) {
	id := c.Param("id")

	// Simulate evaluation execution
	results := simulateEvaluation()
	totalScore := computeTotalScore(results)
	passRate := computePassRate(results)

	report := &EvalReport{
		ID:          id,
		Results:     results,
		TotalScore:  totalScore,
		PassRate:    passRate,
		EvalModel:   "gpt-4o",
		CompletedAt: time.Now(),
	}

	c.JSON(200, report)
}

// GetReport returns the evaluation report
func (e *Engine) GetReport(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "report": nil})
}

func defaultDimensions() []DimensionConfig {
	return []DimensionConfig{
		{Dimension: DimensionAccuracy, Weight: 0.2, Threshold: 0.8},
		{Dimension: DimensionRelevance, Weight: 0.15, Threshold: 0.75},
		{Dimension: DimensionFactuality, Weight: 0.2, Threshold: 0.85},
		{Dimension: DimensionCompliance, Weight: 0.15, Threshold: 0.9},
		{Dimension: DimensionFluency, Weight: 0.1, Threshold: 0.7},
		{Dimension: DimensionHarmfulness, Weight: 0.2, Threshold: 0.95},
	}
}

func simulateEvaluation() []EvalResult {
	dims := []Dimension{
		DimensionAccuracy, DimensionRelevance, DimensionFactuality,
		DimensionCompliance, DimensionFluency, DimensionHarmfulness,
	}
	results := make([]EvalResult, len(dims))
	for i, d := range dims {
		score := 0.7 + math.Round(0.3*float64(i%5)/4*100) / 100
		results[i] = EvalResult{
			Dimension: d,
			Score:     score,
			Weight:    0.15,
			Passed:    score >= 0.8,
		}
	}
	return results
}

func computeTotalScore(results []EvalResult) float64 {
	var totalWeight, weightedSum float64
	for _, r := range results {
		weightedSum += r.Score * r.Weight
		totalWeight += r.Weight
	}
	if totalWeight == 0 {
		return 0
	}
	return math.Round(weightedSum/totalWeight*100) / 100
}

func computePassRate(results []EvalResult) float64 {
	if len(results) == 0 {
		return 0
	}
	passed := 0
	for _, r := range results {
		if r.Passed {
			passed++
		}
	}
	return float64(passed) / float64(len(results))
}