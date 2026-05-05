package regression

import (
	"time"

	"github.com/gin-gonic/gin"
)

// TestCase represents a test case for regression testing
type TestCase struct {
	ID          string            json:"id"
	Name        string            json:"name"
	Category    string            json:"category"
	Input       string            json:"input"
	Expected    string            json:"expected"
	Variables   map[string]string json:"variables,omitempty"
	Tags        []string          json:"tags,omitempty"
	Priority    int               json:"priority"
	Enabled     bool              json:"enabled"
}

// TestSuite represents a collection of test cases
type TestSuite struct {
	ID          string     json:"id"
	Name        string     json:"name"
	Description string     json:"description"
	Cases       []TestCase json:"cases"
	CreatedAt   time.Time  json:"created_at"
	UpdatedAt   time.Time  json:"updated_at"
}

// TestResult represents the result of a test case execution
type TestResult struct {
	CaseID     string    json:"case_id"
	Passed     bool      json:"passed"
	Actual     string    json:"actual"
	Score      float64   json:"score"
	Duration   int64     json:"duration_ms"
	Error      string    json:"error,omitempty"
	ExecutedAt time.Time json:"executed_at"
}

// SuiteRun represents a complete regression test run
type SuiteRun struct {
	ID          string       json:"id"
	SuiteID     string       json:"suite_id"
	VersionID   string       json:"version_id"
	Status      string       json:"status"
	Results     []TestResult json:"results"
	TotalCases  int          json:"total_cases"
	PassedCases int          json:"passed_cases"
	FailedCases int          json:"failed_cases"
	PassRate    float64      json:"pass_rate"
	StartedAt   time.Time    json:"started_at"
	CompletedAt time.Time    json:"completed_at,omitempty"
}

// Suite manages regression test suites
type Suite struct {
	db    interface{}
	cache interface{}
}

// NewSuite creates a new regression test suite manager
func NewSuite(db, cache interface{}) *Suite {
	return &Suite{db: db, cache: cache}
}

// List returns all test suites
func (s *Suite) List(c *gin.Context) {
	c.JSON(200, gin.H{"items": []TestSuite{}, "total": 0})
}

// Create creates a new test suite
func (s *Suite) Create(c *gin.Context) {
	var suite TestSuite
	if err := c.ShouldBindJSON(&suite); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	suite.ID = "suite-" + time.Now().Format("20060102150405")
	suite.CreatedAt = time.Now()
	c.JSON(201, suite)
}

// Run executes a regression test suite
func (s *Suite) Run(c *gin.Context) {
	id := c.Param("id")
	run := &SuiteRun{
		ID:         "run-" + time.Now().Format("20060102150405"),
		SuiteID:    id,
		Status:     "running",
		StartedAt:  time.Now(),
	}
	c.JSON(200, run)
}

// GetResult returns the result of a suite run
func (s *Suite) GetResult(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "status": "completed", "pass_rate": 0.92})
}