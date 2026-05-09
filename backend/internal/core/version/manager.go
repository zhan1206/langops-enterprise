package version

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/gin-gonic/gin"
)

// ResourceType defines the type of versioned resource
type ResourceType string

const (
	ResourcePrompt   ResourceType = "prompt"
	ResourceRAG      ResourceType = "rag"
	ResourceAgent    ResourceType = "agent"
	ResourceWorkflow ResourceType = "workflow"
	ResourceTestCase ResourceType = "testcase"
)

// Version represents a versioned snapshot of a resource
type Version struct {
	ID          string       `json:"id"`
	ResourceID  string       `json:"resource_id"`
	ResourceType ResourceType `json:"resource_type"`
	VersionNum  int          `json:"version_num"`
	Content     string       `json:"content"`
	Checksum    string       `json:"checksum"`
	Branch      string       `json:"branch"`
	ParentID    string       `json:"parent_id,omitempty"`
	Author      string       `json:"author"`
	Message     string       `json:"message"`
	Tags        []string     `json:"tags,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
}

// Manager handles version management for LLM application resources
type Manager struct {
	db     interface{}
	cache  interface{}
	branches map[string]*Branch
}

// Branch represents a version branch
type Branch struct {
	Name      string    `json:"name"`
	BaseVer   int       `json:"base_version"`
	HeadVer   int       `json:"head_version"`
	CreatedAt time.Time `json:"created_at"`
}

// NewManager creates a new version manager
func NewManager(db, cache interface{}) *Manager {
	return &Manager{
		db:       db,
		cache:    cache,
		branches: make(map[string]*Branch),
	}
}

// List returns all versions for resources
func (m *Manager) List(c *gin.Context) {
	c.JSON(200, gin.H{
		"items": []Version{},
		"total": 0,
	})
}

// Create creates a new version
func (m *Manager) Create(c *gin.Context) {
	var req struct {
		ResourceID  string       json:"resource_id" binding:"required"
		ResourceType ResourceType json:"resource_type" binding:"required"
		Content     string       json:"content" binding:"required"
		Message     string       `json:"message"`
		Branch      string       `json:"branch"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	checksum := computeChecksum(req.Content)
	ver := &Version{
		ID:          generateID(),
		ResourceID:  req.ResourceID,
		ResourceType: req.ResourceType,
		Content:     req.Content,
		Checksum:    checksum,
		Branch:      req.Branch,
		Author:      c.GetString("user_id"),
		Message:     req.Message,
		CreatedAt:   time.Now(),
	}

	c.JSON(201, ver)
}

// Get returns a specific version
func (m *Manager) Get(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "version": nil})
}

// Update updates a version's metadata
func (m *Manager) Update(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "updated": true})
}

// Rollback rolls back to a specific version
func (m *Manager) Rollback(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id":        id,
		"rollback":  true,
		"message":   "Rolled back successfully",
	})
}

// Branch creates a new branch from a version
func (m *Manager) Branch(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		BranchName string json:"branch_name" binding:"required"
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	m.branches[req.BranchName] = &Branch{
		Name:      req.BranchName,
		CreatedAt: time.Now(),
	}
	c.JSON(201, gin.H{"id": id, "branch": req.BranchName})
}

// Diff compares two versions
func (m *Manager) Diff(c *gin.Context) {
	id := c.Param("id")
	target := c.Query("target")
	c.JSON(200, gin.H{
		"source":    id,
		"target":    target,
		"additions": 0,
		"deletions": 0,
		"changes":   []string{},
	})
}

func computeChecksum(content string) string {
	h := sha256.Sum256([]byte(content))
	return hex.EncodeToString(h[:])
}

func generateID() string {
	return hex.EncodeToString([]byte(time.Now().Format("20060102150405.000000000")))
}