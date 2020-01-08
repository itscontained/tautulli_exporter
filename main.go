package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	//Create a new instance of the foocollector and
	//register it with the prometheus client.
	foo := newFooCollector()
	prometheus.MustRegister(foo)
	//This section will start the HTTP server and expose
	//any metrics on the /metrics endpoint.
	http.Handle("/metrics", promhttp.Handler())
	log.Info("Beginning to serve on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
