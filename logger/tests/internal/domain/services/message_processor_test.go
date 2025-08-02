package services_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Lucashcr/openservicebroker/logger/internal/domain/models"
	"github.com/Lucashcr/openservicebroker/logger/internal/domain/services"
)

func GenerateLokiMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, _ := io.ReadAll(r.Body)

		var payload models.Payload
		err := json.Unmarshal(body, &payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if len(payload.Streams) != 1 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}))
}

func TestShouldProcessMessage(t *testing.T) {
	mockServer := GenerateLokiMockServer()
	defer mockServer.Close()

	message := models.Message{
		Timestamp:   time.Date(2025, 7, 27, 0, 0, 0, 0, time.UTC),
		ServiceName: "worker02",
		Job:         "Provision-Postgresql",
		Level:       "DEBUG",
		Line:        "Provisioned successfully",
	}

	err := services.ProcessMessage(message, mockServer.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestShouldReturnErrorOnInvalidLokiResponse(t *testing.T) {
	mockServer := GenerateLokiMockServer()
	defer mockServer.Close()

	message := models.Message{
		Timestamp:   time.Date(2025, 7, 27, 0, 0, 0, 0, time.UTC),
		ServiceName: "worker02",
		Job:         "Provision-Postgresql",
		Level:       "DEBUG",
		Line:        "Provisioned successfully",
	}

	invalidURL := mockServer.URL + "/invalid"

	err := services.ProcessMessage(message, invalidURL)
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}
