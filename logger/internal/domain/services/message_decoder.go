package services

import (
	"bytes"
	"encoding/json"

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

	return messageData, nil
}
