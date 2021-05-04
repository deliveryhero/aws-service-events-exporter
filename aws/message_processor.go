package aws

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	log "github.com/sirupsen/logrus"
)

type MessageProcessor struct {
	client   sqs.SQS
	queueUrl string
}

func (messageProcessor *MessageProcessor) pollQueue() error {
	for {
		log.Debug("Long polling for a message... (Will wait for 10s)", messageProcessor.queueUrl)
		receiveMessageRequest := &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(messageProcessor.queueUrl),
			MaxNumberOfMessages: aws.Int64(10),
			VisibilityTimeout:   aws.Int64(10),
			WaitTimeSeconds:     aws.Int64(10),
		}
		receiveMessageOutput, receiveError := messageProcessor.client.ReceiveMessage(receiveMessageRequest)
		if receiveError != nil {
			return receiveError
		}
		if len(receiveMessageOutput.Messages) < 1 {
			log.Info("No messages on queue. Will sleep for 1s, then long poll again.")
			time.Sleep(time.Second)
		}

		for _, message := range receiveMessageOutput.Messages {
			go messageProcessor.processMessage(message)
		}
	}
}

func (messageProcessor *MessageProcessor) processMessage(message *sqs.Message) {
	msg := Message{}
	err := json.Unmarshal([]byte(aws.StringValue(message.Body)), &msg)
	if err != nil {
		log.Error(err)
		return
	}
	events, err := strconv.Unquote(string(msg.Message))

	if err != nil {
		log.Error(err)
	}

	event := RdsEventMessage{}
	err = json.Unmarshal([]byte(events), &event)
	if err != nil {
		log.Error(err)
		return
	}
	eventId := strings.Split(event.EventID, "#")
	if len(eventId) == 1 {
		EventsCounter.WithLabelValues("none", event.EventMessage, event.SourceId).Inc()
		return
	}

	EventsCounter.WithLabelValues(eventId[1], event.EventMessage, event.SourceId).Inc()
}
