package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Lucashcr/openservicebroker/logger/internal/domain/models"
)

func ProcessMessage(message models.Message, lokiUrl string) error {
	labels := make(map[string]string)
	labels["service_name"] = message.ServiceName
	labels["job"] = message.Job
	labels["detected_level"] = message.Level

	values := make([][]any, 0, 1)
	values = append(values, message.Format())

	streams := make([]models.Stream, 0, 1)
	streams = append(streams, models.Stream{
		Labels: labels,
		Values: values,
	})

	payload := models.Payload{Streams: streams}
	payloadJson, err := json.Marshal(&payload)
	if err != nil {
		return err
	}
	fmt.Println(string(payloadJson))

	response, err := http.Post(lokiUrl, "application/json", bytes.NewReader(payloadJson))
	if err != nil {
		return err
	}

	if response.StatusCode != 204 {
		errorMessage := fmt.Sprint("Unexpected response status: ", response.Status)
		return errors.New(errorMessage)
	}

	log.Println("Logs sent successfully!")
	return nil
}
