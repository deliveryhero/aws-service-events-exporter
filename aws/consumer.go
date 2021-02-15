package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var EventsCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "aws_events",
		Help: "This metric indicates which event happened on rds db instances",
	},
	[]string{"event_id", "event_message", "event_source"},
)

func Consume(ctx *cli.Context) {
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
		log.Fatal(err)
	}
	client := sqs.New(sess)
	messageProcessor := &MessageProcessor{*client, QueueUrl}
	err = messageProcessor.pollQueue() // Blocks
	if err != nil {
		log.Fatal(err)
	}

}
