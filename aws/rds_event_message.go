package aws

// RdsEventMessage is a AWS RDS Events Message template which needs to parse from SQS queues
type RdsEventMessage struct {
	EventSource    string `json:"Event Source"`
	EventTime      string `json:"Event Time"`
	IdentifierLink string `json:"Identifier Link"`
	SourceID       string `json:"Source ID"`
	SourceARN      string `json:"Source ARN"`
	EventID        string `json:"Event ID"`
	EventMessage   string `json:"Event Message"`
}
