package models

type Stream struct {
	Labels map[string]string `json:"stream"`
	Values [][]any           `json:"values"`
}

type Payload struct {
	Streams []Stream `json:"streams"`
}
