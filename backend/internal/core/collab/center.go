package collab

import (
	"time"

	"github.com/gin-gonic/gin"
)

// ReviewStatus defines the status of a review
type ReviewStatus string

const (
	ReviewPending  ReviewStatus = "pending"
	ReviewApproved ReviewStatus = "approved"
	ReviewRejected ReviewStatus = "rejected"
	ReviewChanges  ReviewStatus = "changes_requested"
)

// Review represents a code/config review
type Review struct {
	ID          string       `json:"id"`
	ResourceID  string       `json:"resource_id"`
	VersionID   string       `json:"version_id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Author      string       `json:"author"`
	Reviewers   []string     `json:"reviewers"`
	Status      ReviewStatus `json:"status"`
	Comments    []Comment    `json:"comments"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

// Comment represents a review comment
type Comment struct {
	ID        string    `json:"id"`
	ReviewID  string    `json:"review_id"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	Line      int       `json:"line,omitempty"`
	Resolved  bool      `json:"resolved"`
	CreatedAt time.Time `json:"created_at"`
}

// Center manages collaboration and reviews
type Center struct {
	db      interface{}
	cache   interface{}
	reviews map[string]*Review
}

// NewCenter creates a new collaboration center
func NewCenter(db, cache interface{}) *Center {
	return &Center{
		db:      db,
		cache:   cache,
		reviews: make(map[string]*Review),
	}
}

// ListReviews returns all reviews
func (cc *Center) ListReviews(c *gin.Context) {
	c.JSON(200, gin.H{"items": []Review{}, "total": 0})
}

// CreateReview creates a new review request
func (cc *Center) CreateReview(c *gin.Context) {
	var review Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	review.ID = "rev-" + time.Now().Format("20060102150405")
	review.Status = ReviewPending
	review.CreatedAt = time.Now()
	cc.reviews[review.ID] = &review
	c.JSON(201, review)
}

// Approve approves a review
func (cc *Center) Approve(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "status": "approved"})
}

// Reject rejects a review
func (cc *Center) Reject(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "status": "rejected"})
}

// AddComment adds a comment to a review
func (cc *Center) AddComment(c *gin.Context) {
	var comment Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	comment.ID = "cmt-" + time.Now().Format("20060102150405")
	comment.CreatedAt = time.Now()
	c.JSON(201, comment)
}