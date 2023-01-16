package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func hello(w http.ResponseWriter, req *http.Request) {

	httpRequestTotal.With(prometheus.Labels{"method": "GET"}).Inc()
	fmt.Fprintf(w, "hello world!\n")

}

var (
	httpRequestTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "The total number of http request.",
		},
		[]string{"method"},
	)
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
