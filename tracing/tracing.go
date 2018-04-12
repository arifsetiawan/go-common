package tracing

import (
	"time"

	"github.com/arifsetiawan/go-common/env"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

// SimpleTracing will calculate difference between current and previous log
type SimpleTracing struct {
	StartTime *time.Time
	AppName string
}

// NewSimpleTracing is
func NewSimpleTracing(appName string) *SimpleTracing {
	return &SimpleTracing{
		AppName: appName,
	}
}

// Start is
func (s *SimpleTracing) Start(tenant string, message string) {

	now := time.Now()
	s.StartTime = &now
	log.WithFields(log.Fields{
		"app":    s.AppName,
		"type":   "tracing",
		"tenant": tenant,
	}).Info(message)

}

// End is
func (s *SimpleTracing) End(tenant string, message string) {

	now := time.Now()
	log.WithFields(log.Fields{
		"app":    s.AppName,
		"type":    "tracing",
		"tenant":  tenant,
		"latency": now.Sub(*s.StartTime).Nanoseconds(),
	}).Info(message)
}
