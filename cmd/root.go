package cmd

import (
	"fmt"
	"os"

	"github.com/deliveryhero/aws-service-events-exporter/aws"
	"github.com/deliveryhero/aws-service-events-exporter/metrics"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const bannerMsg = `
    ___        _  _
   /   \  ___ | |(_)__   __  ___  _ __  _   _    /\  /\  ___  _ __   ___
  / /\ / / _ \| || |\ \ / / / _ \| '__|| | | |  / /_/ / / _ \| '__| / _ \
 / /_// |  __/| || | \ V / |  __/| |   | |_| | / __  / |  __/| |   | (_) |
/___,'   \___||_||_|  \_/   \___||_|    \__, | \/ /_/   \___||_|    \___/
                                        |___/
`

func RootCmd() {
	app := cli.NewApp()
	app.Name = "RDS events exporter"
	app.Usage = "This app exports your RDS events as metric by sqs to /metrics"
	app.Description = "This app exports your RDS events as metric by sqs to /metrics"
	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "listen-address",
			Value:  "0.0.0.0",
			Usage:  "address to listen on for metrics",
			EnvVar: "LISTEN_ADDRESS",
		},
		cli.IntFlag{
			Name:   "port",
			Value:  9090,
			Usage:  "port to listen on for metrics",
			EnvVar: "PORT",
		},
		cli.StringFlag{
			Name:   "queue-url",
			Value:  "",
			Usage:  "queue url that you need to consume rds events",
			EnvVar: "QUEUE_URL",
		},
		cli.StringFlag{
			Name:   "region",
			Value:  "eu-west-1",
			Usage:  "AWS region",
			EnvVar: "AWS_REGION",
		},
		cli.BoolFlag{
			Name:   "verbose",
			Usage:  "runs in verbose mode",
			EnvVar: "VERBOSE",
		},
	}
	app.Flags = myFlags
	app.Action = run

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	fmt.Print(bannerMsg)
	fmt.Println()
	log.Info(fmt.Sprintf("Prometheus exporter starting to listen on %s:%d/metrics", c.String("listen-address"), c.Int("port")))
	go aws.Consume(c)
	metrics.Serve(c)
	return nil
}
