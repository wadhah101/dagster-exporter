package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Khan/genqlient/graphql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metrics struct {
	dagsterRuns *prometheus.GaugeVec
}

func getHandler() *http.Handler {
	GRAPHQL_ENDPOINT := os.Getenv("GRAPHQL_ENDPOINT")
	reg := prometheus.NewRegistry()
	metrics := &metrics{
		dagsterRuns: nil,
	}

	metrics.dagsterRuns = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "dagster_runs_running",
		Help: "The number of dagster runs running",
	}, []string{"status"})

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

		metrics.dagsterRuns.WithLabelValues(string(status)).Set(float64(cast.Count))
	}

	reg.MustRegister(metrics.dagsterRuns)

	handler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})

	return &handler
}

func main() {

	handler := getHandler()

	http.Handle("/metrics", *handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
