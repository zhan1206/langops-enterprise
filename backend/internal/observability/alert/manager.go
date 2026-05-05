package alert

import "time"

type AlertLevel string

const (
	AlertInfo    AlertLevel = "info"
	AlertWarning AlertLevel = "warning"
	AlertError   AlertLevel = "error"
	AlertCritical AlertLevel = "critical"
)

type AlertChannel string

const (
	ChannelEmail    AlertChannel = "email"
	ChannelWeChat   AlertChannel = "wechat"
	ChannelDingTalk AlertChannel = "dingtalk"
	ChannelSlack    AlertChannel = "slack"
	ChannelWebhook  AlertChannel = "webhook"
)

type AlertRule struct {
	ID        string       json:"id"
	Name      string       json:"name"
	Metric    string       json:"metric"
	Condition string       json:"condition"
	Threshold float64      json:"threshold"
	Level     AlertLevel   json:"level"
	Channels  []AlertChannel json:"channels"
	Enabled   bool         json:"enabled"
}

type Alert struct {
	ID        string     json:"id"
	RuleID    string     json:"rule_id"
	Level     AlertLevel json:"level"
	Message   string     json:"message"
	Resolved  bool       json:"resolved"
	CreatedAt time.Time  json:"created_at"
}

type Manager struct {
	config interface{}
	logger interface{}
	rules  []AlertRule
}

func NewManager(config, logger interface{}) *Manager {
	return &Manager{config: config, logger: logger}
}

func (m *Manager) SendAlert(rule *AlertRule, message string) error {
	// Send alert through configured channels
	return nil
}

func (m *Manager) ListAlerts(resolved bool) []Alert {
	return []Alert{}
}