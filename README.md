# Kube-KPIFetch

Kube-KPIFetch, is a go-lang written program to push aggregated metrics to a endpoint such as prometheus. The advantage is to be light and fast at compiling metrics that can potentially slow down applications such as Grafana to compile when environments have grown tramendously in a single/multi-tenant Kubernetes cluster.

Currently collected metrics:

- Pod
- Replicasets
- Deployments
- Services
- Cronjobs
- Jobs
- Horizontal Pod Autoscaler
- Persistent Volumes
- Persistent Volume Claims
- Statefulsets
- Namespaces
