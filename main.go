package main

import (
	"log"
	"net/http"

	"math/rand"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metrics struct {
	dagsterRunsRunning prometheus.Collector
}

func getHandler() *http.Handler {
	reg := prometheus.NewRegistry()
	metrics := &metrics{
		dagsterRunsRunning: nil,
	}

	metrics.dagsterRunsRunning = promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "dagster_runs_running",
		Help: "The number of dagster runs running",
	}, func() float64 {
		return float64(rand.Intn(100))
	})

	reg.MustRegister(metrics.dagsterRunsRunning)

	handler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})

	return &handler
}

func main() {
	handler := getHandler()

	http.Handle("/metrics", *handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
