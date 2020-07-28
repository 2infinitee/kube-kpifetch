package metrics

import (
	"testing"
)

func TestPushGatewayMetrics(t *testing.T) {
	gatewayUrl := "127.0.0.1:9091/"

	fakeMetric := KPIMetric{
		Name:       "pagerduty_incidents_past_24_hours",
		Help:       "PagerDuty incident count in the past 24 hours",
		Value:      0,
		Cluster: 	"bravo.k8s.ava.theplatform.com",
	}

	fakeMetric2 := KPIMetric{
		Name:       "pagerduty_incidents_past_24_hours",
		Help:       "PagerDuty incident count in the past 24 hours",
		Value:      0,
		Cluster: 	"bravo.k8s.asg.theplatform.com",
	}

	fakeMetric3 := KPIMetric{
		Name:       "pagerduty_incidents_past_24_hours",
		Help:       "PagerDuty incident count in the past 24 hours",
		Value:      0,
		Cluster: 	"bravo.k8s.ade.theplatform.com",
	}

	fakeMetric4 := KPIMetric{
		Name:       "pagerduty_incidents_past_24_hours",
		Help:       "PagerDuty incident count in the past 24 hours",
		Value:      0,
		Cluster: 	"bravo.k8s.aor.theplatform.com",
	}

	fakeMetric5 := KPIMetric{
		Name:       "pagerduty_incidents_past_24_hours",
		Help:       "PagerDuty incident count in the past 24 hours",
		Value:      0,
		Cluster: 	"bravo.k8s.aort.theplatform.com",
	}

	fakeMetric6 := KPIMetric{
		Name:       "pagerduty_incidents_past_24_hours",
		Help:       "PagerDuty incident count in the past 24 hours",
		Value:      0,
		Cluster: 	"bravo.k8s.aie.theplatform.com",
	}

	var kpiMetrics []KPIMetric
	kpiMetrics = append(kpiMetrics, fakeMetric)
	kpiMetrics = append(kpiMetrics, fakeMetric2)
	kpiMetrics = append(kpiMetrics, fakeMetric3)
	kpiMetrics = append(kpiMetrics, fakeMetric4)
	kpiMetrics = append(kpiMetrics, fakeMetric5)
	kpiMetrics = append(kpiMetrics, fakeMetric6)

	err := PushMetrics(gatewayUrl, kpiMetrics)
	if err != nil {
		t.Fatal(err)
	}

}

func TestPushGatewayMetricsRecover(t *testing.T) {
	gatewayUrl := "127.0.0.1:9091/"

	fakeMetric := KPIMetric{
		Name:       "pagerduty_incidents_past_24_hours",
		Help:       "PagerDuty incident count in the past 24 hours",
		Value:      0,
		Cluster: 	"bravo.k8s.ava.theplatform.com",
	}

	fakeMetric2 := KPIMetric{
		Name:       "pagerduty_incidents_past_24_hours",
		Help:       "PagerDuty incident count in the past 24 hours",
		Value:      0,
		Cluster: 	"bravo.k8s.ava.theplatform.com",
	}

	fakeMetric3 := KPIMetric{
		Name:       "pagerduty_incidents_past_24_hours",
		Help:       "PagerDuty incident count in the past 24 hours",
		Value:      0,
		Cluster: 	"bravo.k8s.ade.theplatform.com",
	}

	fakeMetric4 := KPIMetric{
		Name:       "pagerduty_incidents_past_24_hours",
		Help:       "PagerDuty incident count in the past 24 hours",
		Value:      0,
		Cluster: 	"bravo.k8s.aor.theplatform.com",
	}

	fakeMetric5 := KPIMetric{
		Name:       "pagerduty_incidents_past_24_hours",
		Help:       "PagerDuty incident count in the past 24 hours",
		Value:      0,
		Cluster: 	"bravo.k8s.aort.theplatform.com",
	}

	fakeMetric6 := KPIMetric{
		Name:       "pagerduty_incidents_past_24_hours",
		Help:       "PagerDuty incident count in the past 24 hours",
		Value:      0,
		Cluster: 	"bravo.k8s.aie.theplatform.com",
	}

	var kpiMetrics []KPIMetric
	kpiMetrics = append(kpiMetrics, fakeMetric)
	kpiMetrics = append(kpiMetrics, fakeMetric2)
	kpiMetrics = append(kpiMetrics, fakeMetric3)
	kpiMetrics = append(kpiMetrics, fakeMetric4)
	kpiMetrics = append(kpiMetrics, fakeMetric5)
	kpiMetrics = append(kpiMetrics, fakeMetric6)

	err := PushMetrics(gatewayUrl, kpiMetrics)
	if err != nil {
		t.Fatal(err)
	}

}