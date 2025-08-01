package models_test

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/Lucashcr/openservicebroker/logger/internal/domain/models"
)

func TestMessageFormat(t *testing.T) {
	message := models.Message{
		Timestamp:   time.Now(),
		Level:       "DEBUG",
		ServiceName: "worker01",
		Job:         "provision-service-01",
		Line:        "Service 01 provisioned successfully",
		Metadata:    make(map[string]any),
	}

	values := message.Format()

	expectedValues := []any{strconv.FormatInt(message.Timestamp.UnixNano(), 10), message.Line}
	if message.Metadata != nil {
		expectedValues = append(expectedValues, message.Metadata)
	}

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected %v, got %v", expectedValues, values)
	}
}
