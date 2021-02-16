package aws

import (
	"github.com/aws/aws-sdk-go/aws"
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
	sess, err := session.NewSession(
		&aws.Config{
			Region:                        aws.String(ctx.String("region")),
			CredentialsChainVerboseErrors: aws.Bool(ctx.Bool("verbose")),
			MaxRetries:                    aws.Int(5),
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	messageProcessor := &MessageProcessor{
		*sqs.New(sess),
		ctx.String("queue-url"),
	}

	err = messageProcessor.pollQueue() // Blocks
	if err != nil {
		log.Fatal(err)
	}

}
