package aws

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

type MessageProcessor struct {
	client   sqs.SQS
	queueUrl string
}

func (messageProcessor *MessageProcessor) pollQueue() error {
	for {
		log.Info("Long polling for a message... (Will wait for 10s)", messageProcessor.queueUrl)

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
		log.Info(err)
		return
	}
	events, err := strconv.Unquote(string(msg.Message))
	if err != nil {
		log.Info(err)
	}
	event := RdsEventMessage{}
	err = json.Unmarshal([]byte(events), &event)
	if err != nil {
		log.Info(err)
		return
	}
	eventId := strings.Split(event.EventID, "#")
	if len(eventId) == 1 {
		eventsCounter.WithLabelValues("none", event.EventMessage, event.EventSource).Inc()
		return
	}
	eventsCounter.WithLabelValues(eventId[1], event.EventMessage, event.EventSource).Inc()
	//deleteMessageRequest := &sqs.DeleteMessageInput{
	//	QueueUrl:            aws.String(QueueUrl),
	//	ReceiptHandle: message.ReceiptHandle,
	//}
	//_, err := messageProcessor.client.DeleteMessage(deleteMessageRequest)
	//fmt.Println(err)
}
