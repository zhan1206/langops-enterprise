package release

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Environment defines deployment environments
type Environment string

const (
	EnvDevelopment Environment = "development"
	EnvStaging     Environment = "staging"
	EnvPreprod     Environment = "preproduction"
	EnvProduction  Environment = "production"
)

// ReleaseStatus defines the status of a release
type ReleaseStatus string

const (
	StatusPending   ReleaseStatus = "pending"
	StatusApproved  ReleaseStatus = "approved"
	StatusCanary    ReleaseStatus = "canary"
	StatusProgress  ReleaseStatus = "in_progress"
	StatusComplete  ReleaseStatus = "completed"
	StatusRolledBack ReleaseStatus = "rolled_back"
	StatusFailed    ReleaseStatus = "failed"
)

// Release represents a version release
type Release struct {
	ID           string        json:"id"
	VersionID    string        json:"version_id"
	ResourceID   string        json:"resource_id"
	FromEnv      Environment   json:"from_env"
	ToEnv        Environment   json:"to_env"
	Status       ReleaseStatus json:"status"
	CanaryPct    int           json:"canary_percentage"
	ApprovedBy   string        json:"approved_by,omitempty"
	ApprovedAt   *time.Time    json:"approved_at,omitempty"
	RollbackOn   string        json:"rollback_conditions"
	CreatedAt    time.Time     json:"created_at"
	CompletedAt  *time.Time    json:"completed_at,omitempty"
}

// Manager handles release lifecycle
type Manager struct {
	db    interface{}
	cache interface{}
}

// NewManager creates a new release manager
func NewManager(db, cache interface{}) *Manager {
	return &Manager{db: db, cache: cache}
}

// List returns all releases
func (m *Manager) List(c *gin.Context) {
	c.JSON(200, gin.H{"items": []Release{}, "total": 0})
}

// Create creates a new release
func (m *Manager) Create(c *gin.Context) {
	var rel Release
	if err := c.ShouldBindJSON(&rel); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	rel.ID = "rel-" + time.Now().Format("20060102150405")
	rel.Status = StatusPending
	rel.CreatedAt = time.Now()
	c.JSON(201, rel)
}

// Approve approves a release for deployment
func (m *Manager) Approve(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "status": "approved"})
}

// StartCanary starts a canary deployment
func (m *Manager) StartCanary(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Percentage int json:"percentage" binding:"required"
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"id":       id,
		"status":   "canary",
		"canary_pct": req.Percentage,
	})
}

// Rollback rolls back a release
func (m *Manager) Rollback(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "status": "rolled_back"})
}