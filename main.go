package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	PROMETHEUS_COROUTINE_DURATION = time.Second * 2
)

var (
	metricV = &metrics{
		dagsterRuns: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "dagster_runs",
			Help: "The number of dagster runs defined by status and code location and jobs",
		}, []string{"status"}),
	}
)

type metrics struct {
	dagsterRuns *prometheus.GaugeVec
}

func getRunsMetrics() map[string]float64 {
	GRAPHQL_ENDPOINT := os.Getenv("GRAPHQL_ENDPOINT")

	RunStatuses := []RunStatus{
		RunStatusQueued,
		RunStatusNotStarted,
		RunStatusManaged,
		RunStatusStarting,
		RunStatusStarted,
		RunStatusSuccess,
		RunStatusFailure,
		RunStatusCanceling,
		RunStatusCanceled,
	}

	result := map[string]float64{}

	for _, status := range RunStatuses {
		client := graphql.NewClient(GRAPHQL_ENDPOINT, http.DefaultClient)

		data, err := GetRunsOrError(context.Background(), client, status)

		if err != nil {
			println(err)
			os.Exit(1)
		}

		cast, ok := data.RunsOrError.(*GetRunsOrErrorRunsOrErrorRuns)

		if !ok {
			// format print with data typename
			fmt.Printf("data.RunsOrError is not of type *GetRunsOrErrorRunsOrErrorRuns, it is of type %T", data.RunsOrError)

			// TODO handle this better
			os.Exit(1)
		}
		result[string(status)] = float64(cast.Count)
	}
	return result
}

func recordMetrics() {
	go func() {
		for {
			for k, v := range getRunsMetrics() {
				metricV.dagsterRuns.WithLabelValues(k).Set(v)
			}
			time.Sleep(PROMETHEUS_COROUTINE_DURATION)
		}
	}()
}

func getHandler() *http.Handler {
	reg := prometheus.NewRegistry()
	reg.MustRegister(metricV.dagsterRuns)
	handler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})

	return &handler
}

func main() {
	recordMetrics()
	handler := getHandler()

	http.Handle("/metrics", *handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
