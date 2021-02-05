package aws

import "encoding/json"

type Message struct {
	Message     json.RawMessage `json:"Message"`
	MessageId   string          `json:"MessageId"`
	TopicArn    string          `json:"TopicArn"`
	MessageType string          `json:"Type"`
}
