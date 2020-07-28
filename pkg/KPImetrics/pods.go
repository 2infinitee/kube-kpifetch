package KPIMetrics

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// test
func playPod() {
	fmt.Println("hi")
}

// getPodCounts uses the context and kube client mods and is supposed to return a number of how many pods currently exist in the cluster
func getPodCounts(ctx context.Context, kubeClient *kubernetes.Clientset, namespace string) (int, error) {

	// allPods is a variable that searches for pods and lists them
	allPods, err := kubeClient.CoreV1().Pods().List(ctx, metav1.ListOptions{})
	if err != nil {
		return 0, err
	}

	return len(allPods.Items), nil

}

func getPods(ctx context.Context, kubeClient *kubernetes.Clientset) (map[string]int, error) {

	// look at Joshulynes example
	allPods := make(map[string]int)

	// must grab kube context
	ctx := context.Background()

	namespaces, err := getAllNamespaces(ctx, kubeClient)
	if err != nil {
		return allPods, err
	}

	for _, ns := range namespaces {
		podCount, err := getPodCounts(ctx, kubeClient, ns)
		if err != nil {
			return allPods, err
		}

		allPods[ns] = podCount
	}

	return allPods, nil
}
