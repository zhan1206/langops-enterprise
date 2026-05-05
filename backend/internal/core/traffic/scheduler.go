package traffic

import (
	"time"

	"github.com/gin-gonic/gin"
)

// RuleType defines the type of traffic rule
type RuleType string

const (
	RuleTypePercentage RuleType = "percentage"
	RuleTypeUserLabel  RuleType = "user_label"
	RuleTypeScenario   RuleType = "scenario"
	RuleTypeRegion     RuleType = "region"
	RuleTypeComplexity RuleType = "complexity"
)

// TrafficRule defines how traffic is distributed
type TrafficRule struct {
	ID          string            json:"id"
	Name        string            json:"name"
	Type        RuleType          json:"type"
	ResourceID  string            json:"resource_id"
	Conditions  map[string]string json:"conditions"
	TargetID    string            json:"target_version_id"
	Weight      int               json:"weight"
	Priority    int               json:"priority"
	Enabled     bool              json:"enabled"
	CreatedAt   time.Time         json:"created_at"
}

// Scheduler manages traffic distribution
type Scheduler struct {
	db    interface{}
	cache interface{}
	rules map[string]*TrafficRule
}

// NewScheduler creates a new traffic scheduler
func NewScheduler(db, cache interface{}) *Scheduler {
	return &Scheduler{
		db:    db,
		cache: cache,
		rules: make(map[string]*TrafficRule),
	}
}

// ListRules returns all traffic rules
func (s *Scheduler) ListRules(c *gin.Context) {
	c.JSON(200, gin.H{"items": []TrafficRule{}, "total": 0})
}

// CreateRule creates a new traffic rule
func (s *Scheduler) CreateRule(c *gin.Context) {
	var rule TrafficRule
	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	rule.ID = "rule-" + time.Now().Format("20060102150405")
	rule.CreatedAt = time.Now()
	s.rules[rule.ID] = &rule
	c.JSON(201, rule)
}

// UpdateRule updates a traffic rule
func (s *Scheduler) UpdateRule(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "updated": true})
}

// DeleteRule deletes a traffic rule
func (s *Scheduler) DeleteRule(c *gin.Context) {
	id := c.Param("id")
	delete(s.rules, id)
	c.JSON(200, gin.H{"id": id, "deleted": true})
}