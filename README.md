# go-metrics-k8s


docker build -t ugurbzkrt/goprometheuskeda:v2 .

docker run -d -p 8181:8181 --name goprometheuskeda ugurbzkrt/goprometheuskeda:v2

docker push ugurbzkrt/goprometheuskeda:v2


vi /etc/config/prometheus.yaml || /etc/prometheus/prometheus.yaml



```
1) helm get values prometheus -n prometheus --all > values.yaml

vim values.yaml

/additionalScrapeConfigs => arama yapmak için bunun altına

2)
extraScrapeConfigs: |
  - job_name: 'goprometheus'
    scrape_interval: 15s
    static_configs:
      - targets: ['goprometheus-service.default.svc.cluster.local:8181']


3) helm upgrade --install prometheus prometheus-community/prometheus -f values.yaml -n prometheus
```

## Using Keda
```
apiVersion: keda.sh/v1alpha1
kind: ScaledObject
metadata:
  name: goprometheus-scaledobject
  namespace: default
  labels:
    deploymentName: goprometheus-deployment
spec:
  scaleTargetRef:
    name: goprometheus-deployment
  minReplicaCount: 1
  maxReplicaCount: 3
  triggers:
  - type: prometheus
    metadata:
      serverAddress: http://prometheus-server.prometheus.svc.cluster.local:80
      metricName: product_order_total
      threshold: '20'
      query: sum(product_order_total)
```

