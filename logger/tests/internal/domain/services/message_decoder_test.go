package services_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Lucashcr/openservicebroker/logger/internal/domain/services"
)

func TestShouldDecodeMessageWithMetadata(t *testing.T) {
	encodedMessage := []byte("{\"timestamp\": \"2025-07-27T00:00:00Z\",\"serviceName\": \"worker02\",\"job\": \"Provision-Postgresql\",\"level\": \"DEBUG\",\"line\": \"Provisioned successfuly\",\"metadata\": {\"operationID\": \"teste123\"}}")

	message, err := services.DecodeMessage(encodedMessage)

	if err != nil {
		t.Fatalf("Error decoding message: %v", err)
	}

	expectedTimestamp := time.Date(2025, 7, 27, 0, 0, 0, 0, time.UTC)
	if !message.Timestamp.Equal(expectedTimestamp) {
		t.Fatalf("Expected timestamp to be '2025-07-27T00:00:00Z', got '%s'", message.Timestamp)
	}

	if message.ServiceName != "worker02" {
		t.Fatalf("Expected service name to be 'worker02', got '%s'", message.ServiceName)
	}

	if message.Job != "Provision-Postgresql" {
		t.Fatalf("Expected job to be 'Provision-Postgresql', got '%s'", message.Job)
	}

	if message.Level != "DEBUG" {
		t.Fatalf("Expected level to be 'DEBUG', got '%s'", message.Level)
	}

	if message.Line != "Provisioned successfuly" {
		t.Fatalf("Expected line to be 'Provisioned successfuly', got '%s'", message.Line)
	}

	if message.Metadata["operationID"] != "teste123" {
		t.Fatalf("Expected operationID to be 'teste123', got '%s'", message.Line)
	}
}

func TestShouldDecodeMessageWithoutMetadata(t *testing.T) {
	encodedMessage := []byte("{\"timestamp\": \"2025-07-27T00:00:00Z\",\"serviceName\": \"worker02\",\"job\": \"Provision-Postgresql\",\"level\": \"DEBUG\",\"line\": \"Provisioned successfuly\"}")

	message, err := services.DecodeMessage(encodedMessage)

	if err != nil {
		t.Fatalf("Error decoding message: %v", err)
	}

	expectedTimestamp := time.Date(2025, 7, 27, 0, 0, 0, 0, time.UTC)
	if !message.Timestamp.Equal(expectedTimestamp) {
		t.Fatalf("Expected timestamp to be '2025-07-27T00:00:00Z', got '%s'", message.Timestamp)
	}

	if message.ServiceName != "worker02" {
		t.Fatalf("Expected service name to be 'worker02', got '%s'", message.ServiceName)
	}

	if message.Job != "Provision-Postgresql" {
		t.Fatalf("Expected job to be 'Provision-Postgresql', got '%s'", message.Job)
	}

	if message.Level != "DEBUG" {
		t.Fatalf("Expected level to be 'DEBUG', got '%s'", message.Level)
	}

	if message.Line != "Provisioned successfuly" {
		t.Fatalf("Expected line to be 'Provisioned successfuly', got '%s'", message.Line)
	}
}

func TestShouldNotDecodeMessageWithoutTimestamp(t *testing.T) {
	encodedMessage := []byte("{\"serviceName\": \"worker02\",\"job\": \"Provision-Postgresql\",\"level\": \"DEBUG\",\"line\": \"Provisioned successfuly\",\"metadata\": {\"operationID\": \"teste123\"}}")

	_, err := services.DecodeMessage(encodedMessage)
	if err == nil {
		t.Fatalf("Error expected when decoding message without timestamp, got nil")
	}
}

func TestShouldNotDecodeMessageWithoutJob(t *testing.T) {
	encodedMessage := []byte("{\"timestamp\": \"2025-07-27T00:00:00Z\",\"serviceName\": \"worker02\",\"level\": \"DEBUG\",\"line\": \"Provisioned successfuly\",\"metadata\": {\"operationID\": \"teste123\"}}")

	_, err := services.DecodeMessage(encodedMessage)
	if err == nil {
		t.Fatalf("Error expected when decoding message without job, got nil")
	}
}

func TestShouldNotDecodeMessageWithoutLevel(t *testing.T) {
	encodedMessage := []byte("{\"timestamp\": \"2025-07-27T00:00:00Z\",\"serviceName\": \"worker02\",\"job\": \"Provision-Postgresql\",\"line\": \"Provisioned successfuly\",\"metadata\": {\"operationID\": \"teste123\"}}")

	_, err := services.DecodeMessage(encodedMessage)
	if err == nil {
		t.Fatalf("Error expected when decoding message without level, got nil")
	}
}

func TestShouldNotDecodeMessageWithoutLine(t *testing.T) {
	encodedMessage := []byte("{\"timestamp\": \"2025-07-27T00:00:00Z\",\"serviceName\": \"worker02\",\"job\": \"Provision-Postgresql\",\"level\": \"DEBUG\",\"metadata\": {\"operationID\": \"teste123\"}}")

	_, err := services.DecodeMessage(encodedMessage)
	if err == nil {
		t.Fatalf("Error expected when decoding message without line, got nil")
	}
}

func TestShouldNotDecodeMessageWithoutServiceName(t *testing.T) {
	encodedMessage := []byte("{\"timestamp\": \"2025-07-27T00:00:00Z\",\"job\": \"Provision-Postgresql\",\"level\": \"DEBUG\",\"line\": \"Provisioned successfuly\",\"metadata\": {\"operationID\": \"teste123\"}}")

	_, err := services.DecodeMessage(encodedMessage)
	if err == nil {
		t.Fatalf("Error expected when decoding message without service name, got nil")
	}
}

func TestShouldNotDecodeMessageDueToInvalidTimestamp(t *testing.T) {
	jsonInput := []byte("{\"timestamp\": \"invalid\",\"serviceName\": \"worker02\",\"job\": \"Provision-Postgresql\",\"level\": \"DEBUG\",\"line\": \"Provisioned successfuly\"}")
	_, err := services.DecodeMessage(jsonInput)
	if err == nil {
		t.Fatal("expected error for invalid timestamp, got nil")
	}
}

func TestShouldNotDecodeMessageDueToUnknownField(t *testing.T) {
	jsonInput := []byte(`{"unknown":"???""}`)
	_, err := services.DecodeMessage(jsonInput)
	if err == nil {
		t.Fatal("expected error for unknown field, got nil")
	}
}

func TestShouldNotDecodeMessageDueToInvalidJSON(t *testing.T) {
	jsonInput := []byte(`{"field1":"value",`)
	_, err := services.DecodeMessage(jsonInput)
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
}

func TestShouldNotDecodeMessageDueToNullJSON(t *testing.T) {
	jsonInput := []byte(`{}`)
	msg, err := services.DecodeMessage(jsonInput)
	fmt.Println(msg)
	if err == nil {
		t.Fatal("expected error for null JSON, got nil")
	}
}
