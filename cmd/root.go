package cmd

import (
	"fmt"
	"github.com/deliveryhero/log-rds-events-exporter/aws"
	"github.com/urfave/cli"
	"log"
	"os"
)
const bannerMsg = `
    ___        _  _                                                       
   /   \  ___ | |(_)__   __  ___  _ __  _   _    /\  /\  ___  _ __   ___  
  / /\ / / _ \| || |\ \ / / / _ \| '__|| | | |  / /_/ / / _ \| '__| / _ \ 
 / /_// |  __/| || | \ V / |  __/| |   | |_| | / __  / |  __/| |   | (_) |
/___,'   \___||_||_|  \_/   \___||_|    \__, | \/ /_/   \___||_|    \___/ 
                                        |___/                             
`

func RootCmd(){
	app := cli.NewApp()
	app.Name = "RDS events exporter"
	app.Usage = "This app exports your RDS events as metric by sqs to /metrics"
	app.Description = "This app exports your RDS events as metric by sqs to /metrics"
	myFlags := []cli.Flag{
		cli.StringFlag{
			Name: "listen-address",
			Value: "0.0.0.0",
			EnvVar: "LISTEN_ADDRESS",
		},
		cli.IntFlag{
			Name: "port",
			Value: 9090,
			EnvVar: "PORT",
		},
		cli.StringFlag{
			Name: "queue-url",
			Value: "",
			EnvVar: "QUEUE_URL",
		},
		cli.StringFlag{
			Name: "region",
			Value: "eu-west-1",
			EnvVar: "AWS_REGION",
		},
		cli.StringFlag{
			Name: "cred-path",
			Value: "",
			EnvVar: "AWS_CREDENTIALS_PATH",
		},
		cli.StringFlag{
			Name: "profile",
			Value: "default",
			EnvVar: "AWS_PROFILE",
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
	log.Println(fmt.Sprintf("Prometheus exporter starting to listen on %s:%d/metrics",c.String("listen-address"), c.Int("port")))
	err := aws.Consume(c)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}