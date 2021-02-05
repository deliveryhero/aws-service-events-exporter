package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"net/http"
)

func Serve(c *cli.Context) {
	http.Handle("/metrics", promhttp.Handler())
	log.Infof("Beginning to serve on port %s:%s", c.String("listen-address"), c.Int("port"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", c.String("listen-address"), c.Int("port")), nil))
}
