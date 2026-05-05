package datamask

import (
	"regexp"
	"strings"
)

type MaskRule struct {
	ID       string json:"id"
	Name     string json:"name"
	Category string json:"category"
	Pattern  string json:"pattern"
	MaskChar string json:"mask_char"
	KeepLen  int    json:"keep_visible_length"
	Enabled  bool   json:"enabled"
}

type MaskResult struct {
	RuleID   string json:"rule_id"
	Original string json:"original"
	Masked   string json:"masked"
	Matched  bool   json:"matched"
}

type Engine struct {
	config interface{}
	logger interface{}
	rules  []MaskRule
}

func NewEngine(config, logger interface{}) *Engine {
	return &Engine{
		config: config,
		logger: logger,
		rules:  defaultMaskRules(),
	}
}

func (e *Engine) Mask(content string) (string, []MaskResult) {
	var results []MaskResult
	masked := content
	for _, rule := range e.rules {
		if !rule.Enabled {
			continue
		}
		re, err := regexp.Compile(rule.Pattern)
		if err != nil {
			continue
		}
		matches := re.FindAllString(masked, -1)
		for _, match := range matches {
			replacement := maskValue(match, rule.MaskChar, rule.KeepLen)
			masked = strings.Replace(masked, match, replacement, 1)
			results = append(results, MaskResult{
				RuleID:   rule.ID,
				Original: match,
				Masked:   replacement,
				Matched:  true,
			})
		}
	}
	return masked, results
}

func maskValue(value, maskChar string, keepLen int) string {
	if keepLen >= len(value) {
		return value
	}
	if keepLen > 0 {
		return value[:keepLen] + strings.Repeat(maskChar, len(value)-keepLen)
	}
	return strings.Repeat(maskChar, len(value))
}

func defaultMaskRules() []MaskRule {
	return []MaskRule{
		{ID: "mr-1", Name: "ID Card", Category: "pii", Pattern: \d{17}[\dXx], MaskChar: "*", KeepLen: 4, Enabled: true},
		{ID: "mr-2", Name: "Phone", Category: "pii", Pattern: 1[3-9]\d{9}, MaskChar: "*", KeepLen: 3, Enabled: true},
		{ID: "mr-3", Name: "Bank Card", Category: "financial", Pattern: \d{16,19}, MaskChar: "*", KeepLen: 4, Enabled: true},
		{ID: "mr-4", Name: "Email", Category: "pii", Pattern: [a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}, MaskChar: "*", KeepLen: 3, Enabled: true},
		{ID: "mr-5", Name: "API Key", Category: "secret", Pattern: (?:sk|pk)-[a-zA-Z0-9]{20,}, MaskChar: "*", KeepLen: 3, Enabled: true},
	}
}