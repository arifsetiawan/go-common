package tracing

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	//log.SetFormatter(&log.JSONFormatter{})
}

// SimpleTracing will calculate difference between current and previous log
type SimpleTracing struct {
	StartTime *time.Time
	Operation string
}

// NewSimpleTracing is
func NewSimpleTracing(appName string) *SimpleTracing {
	return &SimpleTracing{
		Operation: appName,
	}
}

// Start is
func (s *SimpleTracing) Start(tenant string, message string) {
	now := time.Now()
	s.StartTime = &now
	log.WithFields(log.Fields{
		"operation": s.Operation,
		"type":      "tracing",
		"tenant":    tenant,
	}).Info(message)
}

// End is
func (s *SimpleTracing) End(tenant string, message string) {
	now := time.Now()
	log.WithFields(log.Fields{
		"operation":  s.Operation,
		"type":       "tracing",
		"tenant":     tenant,
		"latency_ms": now.Sub(*s.StartTime).Nanoseconds() / 1e6,
	}).Info(message)
}
