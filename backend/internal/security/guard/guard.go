package guard

type GuardLevel string

const (
	GuardAllow    GuardLevel = "allow"
	GuardWarn     GuardLevel = "warn"
	GuardBlock    GuardLevel = "block"
	GuardSanitize GuardLevel = "sanitize"
)

type GuardRule struct {
	ID          string     json:"id"
	Name        string     json:"name"
	Category    string     json:"category"
	Pattern     string     json:"pattern"
	Level       GuardLevel json:"level"
	Description string     json:"description"
	Enabled     bool       json:"enabled"
}

type GuardResult struct {
	RuleID    string     json:"rule_id"
	RuleName  string     json:"rule_name"
	Category  string     json:"category"
	Level     GuardLevel json:"level"
	Matched   bool       json:"matched"
	Original  string     json:"original,omitempty"
	Sanitized string     json:"sanitized,omitempty"
}

type Guard struct {
	config interface{}
	logger interface{}
	rules  []GuardRule
}

func NewGuard(config, logger interface{}) *Guard {
	return &Guard{
		config: config,
		logger: logger,
		rules:  defaultGuardRules(),
	}
}

func (g *Guard) CheckInput(content string) []GuardResult {
	return g.check(content)
}

func (g *Guard) CheckOutput(content string) []GuardResult {
	return g.check(content)
}

func (g *Guard) check(content string) []GuardResult {
	var results []GuardResult
	for _, rule := range g.rules {
		if !rule.Enabled {
			continue
		}
		results = append(results, GuardResult{
			RuleID:   rule.ID,
			RuleName: rule.Name,
			Category: rule.Category,
			Level:    rule.Level,
			Matched:  false,
		})
	}
	return results
}

func defaultGuardRules() []GuardRule {
	return []GuardRule{
		{ID: "gr-1", Name: "Malicious Instruction", Category: "injection", Level: GuardBlock, Enabled: true},
		{ID: "gr-2", Name: "PII Exposure", Category: "privacy", Level: GuardSanitize, Enabled: true},
		{ID: "gr-3", Name: "Harmful Content", Category: "safety", Level: GuardBlock, Enabled: true},
		{ID: "gr-4", Name: "Compliance Violation", Category: "compliance", Level: GuardWarn, Enabled: true},
	}
}