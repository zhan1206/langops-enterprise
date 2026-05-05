package main

import (
	"log"
	"net/http"
	"os"

	"github.com/zhan1206/langops-enterprise/backend/internal/adapter/autogen"
	"github.com/zhan1206/langops-enterprise/backend/internal/adapter/crewai"
	"github.com/zhan1206/langops-enterprise/backend/internal/adapter/langchain"
	"github.com/zhan1206/langops-enterprise/backend/internal/adapter/llamaindex"
	"github.com/zhan1206/langops-enterprise/backend/internal/adapter/model"
	"github.com/zhan1206/langops-enterprise/backend/internal/adapter/openhands"
	"github.com/zhan1206/langops-enterprise/backend/internal/adapter/tool"
	"github.com/zhan1206/langops-enterprise/backend/internal/core/abtest"
	"github.com/zhan1206/langops-enterprise/backend/internal/core/collab"
	"github.com/zhan1206/langops-enterprise/backend/internal/core/degradation"
	"github.com/zhan1206/langops-enterprise/backend/internal/core/eval"
	"github.com/zhan1206/langops-enterprise/backend/internal/core/regression"
	"github.com/zhan1206/langops-enterprise/backend/internal/core/release"
	"github.com/zhan1206/langops-enterprise/backend/internal/core/rootcause"
	"github.com/zhan1206/langops-enterprise/backend/internal/core/traffic"
	"github.com/zhan1206/langops-enterprise/backend/internal/core/version"
	"github.com/zhan1206/langops-enterprise/backend/internal/observability/alert"
	"github.com/zhan1206/langops-enterprise/backend/internal/observability/cost"
	"github.com/zhan1206/langops-enterprise/backend/internal/observability/metrics"
	"github.com/zhan1206/langops-enterprise/backend/internal/observability/trace"
	"github.com/zhan1206/langops-enterprise/backend/internal/security/audit"
	"github.com/zhan1206/langops-enterprise/backend/internal/security/auth"
	"github.com/zhan1206/langops-enterprise/backend/internal/security/compliance"
	"github.com/zhan1206/langops-enterprise/backend/internal/security/datamask"
	"github.com/zhan1206/langops-enterprise/backend/internal/security/guard"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Server   ServerConfig   yaml:"server"
	Database DatabaseConfig yaml:"database"
	Redis    RedisConfig    yaml:"redis"
	Security SecurityConfig yaml:"security"
}

type ServerConfig struct {
	Port string yaml:"port"
	Mode string yaml:"mode"
}

type DatabaseConfig struct {
	DSN string yaml:"dsn"
}

type RedisConfig struct {
	Addr string yaml:"addr"
}

type SecurityConfig struct {
	JWTSecret string yaml:"jwt_secret"
}

func main() {
	cfg := Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Mode: getEnv("GIN_MODE", "release"),
		},
		Database: DatabaseConfig{
			DSN: getEnv("DATABASE_DSN", "root:password@tcp(localhost:3306)/langops?parseTime=true"),
		},
		Redis: RedisConfig{
			Addr: getEnv("REDIS_ADDR", "localhost:6379"),
		},
		Security: SecurityConfig{
			JWTSecret: getEnv("JWT_SECRET", "langops-default-secret-change-me"),
		},
	}

	gin.SetMode(cfg.Server.Mode)
	r := gin.Default()

	// Middleware
	authManager := auth.NewManager(cfg.Security.JWTSecret, nil)
	r.Use(authManager.Middleware())

	// Initialize core modules
	versionMgr := version.NewManager(nil, nil)
	evalEngine := eval.NewEngine(nil, nil)
	regressionSuite := regression.NewSuite(nil, nil)
	degradationDetector := degradation.NewDetector(nil, nil)
	rootcauseAnalyzer := rootcause.NewAnalyzer(nil, nil)
	releaseManager := release.NewManager(nil, nil)
	abtestEngine := abtest.NewEngine(nil, nil)
	trafficScheduler := traffic.NewScheduler(nil, nil)
	collabCenter := collab.NewCenter(nil, nil)

	// Initialize adapters
	_ = langchain.NewAdapter(nil, nil)
	_ = llamaindex.NewAdapter(nil, nil)
	_ = autogen.NewAdapter(nil, nil)
	_ = crewai.NewAdapter(nil, nil)
	_ = openhands.NewAdapter(nil, nil)
	_ = model.NewRouter(nil, nil)
	_ = tool.NewGateway(nil, nil)

	// Initialize security modules
	_ = audit.NewChain(nil, nil)
	_ = guard.NewGuard(nil, nil)
	_ = datamask.NewEngine(nil, nil)
	_ = compliance.NewEngine(nil, nil)

	// Initialize observability modules
	_ = trace.NewTracer(nil, nil)
	_ = metrics.NewCollector(nil, nil)
	_ = cost.NewManager(nil, nil)
	_ = alert.NewManager(nil, nil)

	// API routes
	api := r.Group("/api/v1")
	{
		// Version management
		versions := api.Group("/versions")
		{
			versions.GET("", versionMgr.List)
			versions.POST("", versionMgr.Create)
			versions.GET("/:id", versionMgr.Get)
			versions.PUT("/:id", versionMgr.Update)
			versions.POST("/:id/rollback", versionMgr.Rollback)
			versions.POST("/:id/branch", versionMgr.Branch)
			versions.GET("/:id/diff", versionMgr.Diff)
		}

		// Evaluation
		evals := api.Group("/evaluations")
		{
			evals.GET("", evalEngine.List)
			evals.POST("", evalEngine.Create)
			evals.POST("/:id/run", evalEngine.Run)
			evals.GET("/:id/report", evalEngine.GetReport)
		}

		// Regression testing
		regressions := api.Group("/regressions")
		{
			regressions.GET("", regressionSuite.List)
			regressions.POST("", regressionSuite.Create)
			regressions.POST("/:id/run", regressionSuite.Run)
			regressions.GET("/:id/result", regressionSuite.GetResult)
		}

		// Degradation detection
		degradations := api.Group("/degradations")
		{
			degradations.GET("/alerts", degradationDetector.ListAlerts)
			degradations.POST("/rules", degradationDetector.CreateRule)
			degradations.GET("/status", degradationDetector.GetStatus)
		}

		// Root cause analysis
		analyses := api.Group("/analyses")
		{
			analyses.POST("", rootcauseAnalyzer.Analyze)
			analyses.GET("/:id", rootcauseAnalyzer.Get)
			analyses.GET("/:id/suggestions", rootcauseAnalyzer.GetSuggestions)
		}

		// Release management
		releases := api.Group("/releases")
		{
			releases.GET("", releaseManager.List)
			releases.POST("", releaseManager.Create)
			releases.POST("/:id/approve", releaseManager.Approve)
			releases.POST("/:id/canary", releaseManager.StartCanary)
			releases.POST("/:id/rollback", releaseManager.Rollback)
		}

		// A/B testing
		abtests := api.Group("/abtests")
		{
			abtests.GET("", abtestEngine.List)
			abtests.POST("", abtestEngine.Create)
			abtests.POST("/:id/start", abtestEngine.Start)
			abtests.POST("/:id/stop", abtestEngine.Stop)
			abtests.GET("/:id/result", abtestEngine.GetResult)
		}

		// Traffic scheduling
		traffics := api.Group("/traffic")
		{
			traffics.GET("/rules", trafficScheduler.ListRules)
			traffics.POST("/rules", trafficScheduler.CreateRule)
			traffics.PUT("/rules/:id", trafficScheduler.UpdateRule)
			traffics.DELETE("/rules/:id", trafficScheduler.DeleteRule)
		}

		// Collaboration
		collabs := api.Group("/collab")
		{
			collabs.GET("/reviews", collabCenter.ListReviews)
			collabs.POST("/reviews", collabCenter.CreateReview)
			collabs.POST("/reviews/:id/approve", collabCenter.Approve)
			collabs.POST("/reviews/:id/reject", collabCenter.Reject)
			collabs.POST("/comments", collabCenter.AddComment)
		}
	}

	log.Printf("LangOps Enterprise server starting on :%s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}