package audit

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type EventType string

const (
	EventPromptChange  EventType = "prompt_change"
	EventVersionCreate EventType = "version_create"
	EventRelease       EventType = "release"
	EventModelCall     EventType = "model_call"
	EventToolExecute   EventType = "tool_execute"
	EventAuthChange    EventType = "auth_change"
)

type AuditEntry struct {
	ID        string    `json:"id"`
	EventType EventType `json:"event_type"`
	Resource  string    `json:"resource"`
	Action    string    `json:"action"`
	Actor     string    `json:"actor"`
	Details   string    `json:"details"`
	PrevHash  string    `json:"prev_hash"`
	Hash      string    `json:"hash"`
	Timestamp time.Time `json:"timestamp"`
}

type Chain struct {
	db       interface{}
	logger   interface{}
	lastHash string
}

func NewChain(db, logger interface{}) *Chain {
	return &Chain{
		db:       db,
		logger:   logger,
		lastHash: "genesis",
	}
}

func (ch *Chain) Append(eventType EventType, resource, action, actor, details string) *AuditEntry {
	entry := &AuditEntry{
		ID:        "audit-" + time.Now().Format("20060102150405.000000"),
		EventType: eventType,
		Resource:  resource,
		Action:    action,
		Actor:     actor,
		Details:   details,
		PrevHash:  ch.lastHash,
		Timestamp: time.Now(),
	}
	entry.Hash = ch.computeHash(entry)
	ch.lastHash = entry.Hash
	return entry
}

func (ch *Chain) Verify(entries []*AuditEntry) bool {
	for i := 1; i < len(entries); i++ {
		if entries[i].PrevHash != entries[i-1].Hash {
			return false
		}
		if ch.computeHash(entries[i]) != entries[i].Hash {
			return false
		}
	}
	return true
}

func (ch *Chain) computeHash(entry *AuditEntry) string {
	data := entry.ID + string(entry.EventType) + entry.Resource + entry.Action + entry.Actor + entry.PrevHash
	h := sha256.Sum256([]byte(data))
	return hex.EncodeToString(h[:])
}