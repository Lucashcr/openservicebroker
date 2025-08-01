package services

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/Lucashcr/openservicebroker/logger/internal/domain/models"
)

func DecodeMessage(body []byte) (models.Message, error) {
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()

	var messageData models.Message
	err := decoder.Decode(&messageData)
	if err != nil {
		return models.Message{}, err
	}

	if messageData.Timestamp.IsZero() {
		return models.Message{}, errors.New("missing field: timestamp")
	}

	if messageData.Job == "" {
		return models.Message{}, errors.New("missing fields: job")
	}

	if messageData.Level == "" {
		return models.Message{}, errors.New("missing fields: job")
	}

	if messageData.Line == "" {
		return models.Message{}, errors.New("missing fields: job")
	}

	if messageData.ServiceName == "" {
		return models.Message{}, errors.New("missing fields: job")
	}

	return messageData, nil
}
