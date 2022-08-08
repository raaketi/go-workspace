package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	rtr := http.NewServeMux()
	rtr.HandleFunc("/metrics", promhttp.Handler().ServeHTTP)
	rtr.HandleFunc("/api/v1/leagues", league)

	if err := http.ListenAndServe(":3000", rtr); err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}

func league(w http.ResponseWriter, _ *http.Request) {
	counter.Inc()
	_, _ = w.Write([]byte("FIFA"))
}

// Auto register counter metrics.
var counter = promauto.NewCounter(prometheus.CounterOpts{
	Namespace: "football",
	Name:      "leagues_request_counter",
	Help:      "Number of requests",
})
