package metrics

import (
	"fmt"
	"net/http"

	"github.com/deliveryhero/aws-service-events-exporter/aws"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func Serve(c *cli.Context) {
	registry := prometheus.NewRegistry()
	registry.MustRegister(aws.EventsCounter)
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", c.String("listen-address"), c.Int("port")), nil))
}
