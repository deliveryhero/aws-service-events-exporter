package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/urfave/cli"
	"time"
)

type MessageProcessor struct {
	client sqs.SQS
	queueUrl string
}

func Consume(ctx *cli.Context) error {
	QueueUrl := ctx.String("queue-url")
	Region := ctx.String("region")
	CredPath := ctx.String("cred-path")
	CredProfile := ctx.String("profile")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewSharedCredentials(CredPath, CredProfile),
		MaxRetries:  aws.Int(5),
	})
	if err != nil {
		return err
	}
	client := sqs.New(sess)
	messageProcessor := &MessageProcessor{*client, QueueUrl}
	err = messageProcessor.pollQueue() // Blocks
	if err != nil {
		return err
	}
	return nil

}

func (messageProcessor *MessageProcessor) pollQueue() error {
	for {
		fmt.Println("Long polling for a message... (Will wait for 10s)", messageProcessor.queueUrl)

		receiveMessageRequest := &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(messageProcessor.queueUrl),
			MaxNumberOfMessages: aws.Int64(10),
			VisibilityTimeout:   aws.Int64(10),
			WaitTimeSeconds: aws.Int64(10),
		}
		receiveMessageOutput, receiveError := messageProcessor.client.ReceiveMessage(receiveMessageRequest)
		if receiveError != nil{
			return receiveError
		}
		if len(receiveMessageOutput.Messages) < 1 {
			fmt.Println("No messages on queue. Will sleep for 1s, then long poll again.")
			time.Sleep(time.Second)
		}

		for _, message := range receiveMessageOutput.Messages {
			go messageProcessor.processMessage(message)
		}
	}
}

func (messageProcessor *MessageProcessor) processMessage(message *sqs.Message) {
	fmt.Println(&message.Body)
	//deleteMessageRequest := &sqs.DeleteMessageInput{
	//	QueueUrl:            aws.String(QueueUrl),
	//	ReceiptHandle: message.ReceiptHandle,
	//}
	//_, err := messageProcessor.client.DeleteMessage(deleteMessageRequest)
	//fmt.Println(err)
}