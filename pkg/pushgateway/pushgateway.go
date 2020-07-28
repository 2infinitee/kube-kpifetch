package pushgateway

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

type KPIMetric struct {
	Name    string
	Help    string
	Value   float64
	Cluster string
}

// PushMetrics takes in a KPIMetric and pushes the metric to the pushgateway url
// KPIMetric should contain metric name, metric type, help string, and cluster
// Depending on the metric type, the float64 is defaulted to 0.
func PushMetrics(gatewayUrl string, kpiMetrics []KPIMetric) error {

	//gatewayUrl:="k8s-kpi-prometheus-pushgateway.kube-monitoring.svc.cluster.local:9091/"
	//gatewayUrl = "127.0.0.1:9091/"

	fmt.Println("Pushing metrics to pushgateway using url: " + gatewayUrl)

	var allMetrics = make(map[prometheus.Gauge]KPIMetric)

	for _, m := range kpiMetrics {

		var gaugeOpts prometheus.GaugeOpts
		if m.Cluster == "" {
			gaugeOpts = prometheus.GaugeOpts{
				Name: "k8s_kpi_" + m.Name,
				Help: m.Help,
			}
		} else {
			gaugeOpts = prometheus.GaugeOpts{
				Name:        "k8s_kpi_" + m.Name,
				Help:        m.Help,
				ConstLabels: prometheus.Labels{"k8s_cluster_alert": m.Cluster},
			}
		}
		gauge := prometheus.NewGauge(gaugeOpts)
		gauge.Set(m.Value)

		allMetrics[gauge] = m
	}

	registry := prometheus.NewRegistry()
	for gauge, kpiMetric := range allMetrics {
		registerMetric(*registry, gauge, kpiMetric)
	}
	pusher := push.New(gatewayUrl, "KPI Fetcher").Gatherer(registry)
	return pusher.Push()

}

// registerMetric registers a prometheus gauge metric and recovers from panics, especially from duplicate metrics
// collector registration panics. Program does not error but skips the duplicate metric attempting to be registered.
func registerMetric(registry prometheus.Registry, metric prometheus.Gauge, kpiMetric KPIMetric) {
	// Catch panics.
	var r interface{}
	defer func() {
		r = recover()
		if r != nil {
			fmt.Println("Recovered panic:", r)
			fmt.Println("Failed to register metric:", kpiMetric)
		}
	}()
	registry.MustRegister(metric)
}
