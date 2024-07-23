# go-metrics-k8s


docker build -t ugurbzkrt/goprometheuskeda:v2 .

docker run -d -p 8181:8181 --name goprometheuskeda ugurbzkrt/goprometheuskeda:v2

docker push ugurbzkrt/goprometheuskeda:v2
