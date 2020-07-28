package main

import (
	"context"
	"github.com/2infinitee/kube-kpifetch/pkg/KPIMetrics"
	kubeClient "github.com/2infinitee/kube-kpifetch/pkg/kubeClient"
	"github.com/2infinitee/kube-kpifetch/pkg/pushgateway"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

// Global variables for setting global configurations to push metrics to prometheus
// cfg is a short variable for Config
var cfg Config

// configPath by default is set to a string value called "config.yaml"
var configPath = "config.yml"

// clusterName is a variable to identify what cluster metrics is running from, set to hostname for local test.
var clusterName = os.Getenv("CLUSTER_FQDN")

// init is used to locate where the configurations exist on the container
// you can configure this in the configmap.yaml
func init() {
	// setting option to config a user's config path to something else than configPath
	userConfigPath := os.Getenv("KPI_METRIC_CONFIG_PATH")

	// if userConfigPath does not equal a blank value change configPath to equal to userConfigPath
	if userConfigPath != "" {
		configPath = userConfigPath
	}

	// error equals to loaded config path
	err := cfg.Load(configPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(play.Play)

}

func main() {
	log.Println("Starting Kubernetes metric fetch...")

	// local variable, in the form of a list, that collects all the metrics to be sent
	// to the main function and creates a list out of it
	var collectedMetrics []pushgateway.KPIMetric

	// collect all pods in the last 1 hour
	// if error does not equal to a zero, then log the error, or else append
	// everything collected thus far including these new metrics collected.
	kubeKPI, err := getAggrKubeKPIs(1)
	if err != nil {
		log.Errorln("Error getting Kubernetes KPIs: ", err)
	} else {
		collectedMetrics = append(collectedMetrics, kubeKPI...)
	}

	// this is where the code pushes gathered KPI to a endpoint, typically Prometheus pushgateway
	log.Println("About to push: ", len(collectedMetrics), "KPIs")
	err = pushgateway.PushMetrics(cfg.PushGatewayURL, collectedMetrics)
	if err != nil {
		log.Errorln("Error pushing KPI metrics to endpoint", err)
	}

	// c is to create a kube client
	c, err := kubeClient.Create("")
	if err != nil {
		log.Fatalln("Could not create kubeClient, ", err)
	}

	log.Println(getPodsMetrics(), err)

}

func getPodsMetrics(ctx context.Context, kubeClient *kubernetes.Clientset) (int, error) {

	// collectPods grabs all pods and give a number back
	collectPods, err := kubeClient.CoreV1().Pods().List(ctx, metav1.ListOptions{})
	if err != nil {
		return 0, err
	}

	return len(collectPods.Items), nil
}
