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
		Name: "aws_service_events",
		Help: "This metric indicates on whats happening on various aws services, e.g RDS",
	},
	[]string{"event_id", "event_message", "event_source", "event_time"},
)

func Consume(ctx *cli.Context) {
	sess, err := session.NewSession(&aws.Config{
		Region:                        aws.String(ctx.String("region")),
		CredentialsChainVerboseErrors: aws.Bool(ctx.Bool("verbose")),
		MaxRetries:                    aws.Int(5),
	})

	if err != nil {
		log.Fatal(err)
	}

	messageProcessor := &MessageProcessor{*sqs.New(sess), ctx.String("queue-url")}

	err = messageProcessor.pollQueue() // Blocks
	if err != nil {
		log.Fatal(err)
	}

}
