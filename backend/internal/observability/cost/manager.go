package cost

import "time"

type CostRecord struct {
	ID         string  `json:"id"`
	ProjectID  string  `json:"project_id"`
	Team       string  `json:"team"`
	Scenario   string  `json:"scenario"`
	Model      string  `json:"model"`
	InputTokens  int64   `json:"input_tokens"`
	OutputTokens int64   `json:"output_tokens"`
	Cost       float64 `json:"cost"`
	Currency   string  `json:"currency"`
	Timestamp  time.Time `json:"timestamp"`
}

type Budget struct {
	ID        string  `json:"id"`
	ProjectID string  `json:"project_id"`
	Team      string  `json:"team"`
	Monthly   float64 `json:"monthly_budget"`
	Used      float64 `json:"used"`
	Remaining float64 `json:"remaining"`
	AlertPct  float64 `json:"alert_percentage"`
}

type CostSummary struct {
	ProjectID   string  `json:"project_id"`
	Period      string  `json:"period"`
	TotalCost   float64 `json:"total_cost"`
	InputTokens  int64   `json:"total_input_tokens"`
	OutputTokens int64   `json:"total_output_tokens"`
	ByModel     map[string]float64 `json:"by_model"`
	ByScenario  map[string]float64 `json:"by_scenario"`
}

type Manager struct {
	config interface{}
	logger interface{}
}

func NewManager(config, logger interface{}) *Manager {
	return &Manager{config: config, logger: logger}
}

func (m *Manager) Record(record *CostRecord) {
	// Record cost usage
}

func (m *Manager) GetBudget(projectID, team string) *Budget {
	return &Budget{
		ProjectID: projectID,
		Team:      team,
		Monthly:   10000.0,
		AlertPct:  80.0,
	}
}

func (m *Manager) GetSummary(projectID, period string) *CostSummary {
	return &CostSummary{
		ProjectID: projectID,
		Period:    period,
		ByModel:   make(map[string]float64),
		ByScenario: make(map[string]float64),
	}
}