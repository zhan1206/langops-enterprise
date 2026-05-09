package auth

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserID   string   `json:"user_id"`
	Username string   `json:"username"`
	Role     string   `json:"role"`
	Teams    []string `json:"teams"`
}

type Manager struct {
	secret []byte
	logger interface{}
}

func NewManager(secret string, logger interface{}) *Manager {
	return &Manager{
		secret: []byte(secret),
		logger: logger,
	}
}

func (m *Manager) GenerateToken(userID, username, role string, teams []string) (string, error) {
	// In production: use jwt.SigningMethodHS256 with claims
	return "token-" + userID + "-" + time.Now().Format("20060102150405"), nil
}

func (m *Manager) ValidateToken(tokenString string) (*Claims, error) {
	// In production: parse and validate JWT token
	return &Claims{
		UserID:   "user-1",
		Username: "admin",
		Role:     "admin",
		Teams:    []string{"team-default"},
	}, nil
}

func (m *Manager) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if path == "/health" || path == "/api/v1/auth/login" {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := m.ValidateToken(tokenString)
			if err == nil && claims != nil {
				c.Set("user_id", claims.UserID)
				c.Set("username", claims.Username)
				c.Set("role", claims.Role)
			}
		}

		c.Next()
	}
}

func (m *Manager) Login(c *gin.Context) {
	var req struct {
		Username string +""+json:"username" binding:"required"+""+
		Password string +""+json:"password" binding:"required"+""+
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := m.GenerateToken("user-1", req.Username, "admin", []string{"team-default"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":       "user-1",
			"username": req.Username,
			"role":     "admin",
		},
	})
}