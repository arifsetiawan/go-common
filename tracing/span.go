package tracing

import (
	"time"

	"github.com/arifsetiawan/go-common/random"
	log "github.com/sirupsen/logrus"
)

// Span is
type Span struct {
	StartTime    time.Time
	Operation    string
	Tenant       string
	ID           string
	ParentID     string
	FollowFromID string
	Tags         map[string]interface{}
}

// StartSpan is
func StartSpan(tenant string, operation string, parentID string, followFromID string, tags map[string]interface{}) *Span {
	return &Span{
		Tenant:       tenant,
		Operation:    operation,
		StartTime:    time.Now(),
		ID:           random.GenerateAlphaNumeric(16),
		ParentID:     parentID,
		FollowFromID: followFromID,
		Tags:         tags,
	}
}

// Finish is
func (s *Span) Finish() {
	log.WithFields(log.Fields{
		"type":           "tracing",
		"operation":      s.Operation,
		"tenant":         s.Tenant,
		"span_id":        s.ID,
		"parent_id":      s.ParentID,
		"follow_from_id": s.FollowFromID,
		"tags":           s.Tags,
		"latency_ms":     time.Now().Sub(s.StartTime).Nanoseconds() / 1e6,
	}).Info("")	
}
