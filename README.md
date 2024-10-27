# dagster-exporter

Prometheus exporter for Dagster using dagster graphql api
Inspired by [celery-exporter](https://github.com/danihodovic/celery-exporter)

## Features checklist

- [ ] dagster_runs_running # need to have job related labels, have user code labels
- [ ] dagster_runs_pending
- [ ] dagster_runs_scheduled
- [ ] Dagster runs completed # maybe use sqlite db or something local for this
- [ ] Dagster sensor runs completed

- [ ] Health check

## References

https://prometheus.io/docs/instrumenting/writing_exporters/
