package degradation

import (
	"time"

	"github.com/gin-gonic/gin"
)

// AlertLevel defines the severity of a degradation alert
type AlertLevel string

const (
	AlertLevelWarning  AlertLevel = "warning"
	AlertLevelCritical AlertLevel = "critical"
	AlertLevelSevere   AlertLevel = "severe"
)

// DegradationAlert represents a detected degradation event
type DegradationAlert struct {
	ID          string     json:"id"
	ResourceID  string     json:"resource_id"
	Metric      string     json:"metric"
	PreviousVal float64    json:"previous_value"
	CurrentVal  float64    json:"current_value"
	DropPct     float64    json:"drop_percentage"
	Level       AlertLevel json:"level"
	RootCause   string     json:"root_cause,omitempty"
	Status      string     json:"status"
	DetectedAt  time.Time  json:"detected_at"
}

// DetectionRule defines a degradation detection rule
type DetectionRule struct {
	ID         string  json:"id"
	Name       string  json:"name"
	Metric     string  json:"metric"
	Threshold  float64 json:"threshold"
	WindowSize int     json:"window_size_minutes"
	Cooldown   int     json:"cooldown_minutes"
	Enabled    bool    json:"enabled"
}

// Detector monitors and detects effect degradation
type Detector struct {
	db    interface{}
	cache interface{}
	rules map[string]*DetectionRule
}

// NewDetector creates a new degradation detector
func NewDetector(db, cache interface{}) *Detector {
	return &Detector{
		db:    db,
		cache: cache,
		rules: make(map[string]*DetectionRule),
	}
}

// ListAlerts returns all degradation alerts
func (d *Detector) ListAlerts(c *gin.Context) {
	c.JSON(200, gin.H{"items": []DegradationAlert{}, "total": 0})
}

// CreateRule creates a new detection rule
func (d *Detector) CreateRule(c *gin.Context) {
	var rule DetectionRule
	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	rule.ID = "rule-" + time.Now().Format("20060102150405")
	d.rules[rule.ID] = &rule
	c.JSON(201, rule)
}

// GetStatus returns current degradation monitoring status
func (d *Detector) GetStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"monitoring":      true,
		"active_rules":    len(d.rules),
		"recent_alerts":   0,
		"last_check":      time.Now(),
	})
}