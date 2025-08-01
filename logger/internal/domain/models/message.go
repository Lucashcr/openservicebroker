package models

import (
	"strconv"
	"time"
)

type Message struct {
	Timestamp   time.Time      `json:"timestamp"`
	Level       string         `json:"level"`
	ServiceName string         `json:"serviceName"`
	Job         string         `json:"job"`
	Line        string         `json:"line"`
	Metadata    map[string]any `json:"metadata,omitempty"`
}

func (v Message) Format() []any {
	result := make([]any, 0, 3)
	result = append(result, strconv.FormatInt(v.Timestamp.UnixNano(), 10))
	result = append(result, v.Line)
	if v.Metadata != nil {
		result = append(result, v.Metadata)
	}
	return result
}
