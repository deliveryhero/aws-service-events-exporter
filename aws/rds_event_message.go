package aws

type RdsEventMessage struct {
	EventSource    string `json:"Event Source"`
	EventTime      string `json:"Event Time"`
	IdentifierLink string `json:"Identifier Link"`
	SourceId       string `json:"Source ID"`
	SourceARN      string `json:"Source ARN"`
	EventID        string `json:"Event ID"`
	EventMessage   string `json:"Event Message"`
}
