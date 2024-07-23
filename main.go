package main

import (
 "log"
 "net/http"
 "time"

 "github.com/prometheus/client_golang/prometheus"
 "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
 HttpRequestCountWithPath = prometheus.NewCounterVec(
  prometheus.CounterOpts{
   Name: "http_requests_total_with_path",
   Help: "Number of HTTP requests by path.",
  },
  []string{"url"},
 )

 HttpRequestDuration = prometheus.NewHistogramVec(
  prometheus.HistogramOpts{
   Name: "http_request_duration_seconds",
   Help: "Response time of HTTP request.",
  },
  []string{"path"},
 )

 orderBooksCounter = prometheus.NewCounter(
  prometheus.CounterOpts{
   Name: "product_order_total",
   Help: "Total number of product",
  },
 )
)

func init() {
 prometheus.MustRegister(orderBooksCounter)
 prometheus.MustRegister(HttpRequestCountWithPath)
 prometheus.MustRegister(HttpRequestDuration)
}

func main() {
 http.HandleFunc("/product", orderHandler)
 http.Handle("/metrics", promhttp.Handler())

 log.Println("Starting server on :8181")
 log.Fatal(http.ListenAndServe(":8181", nil))
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
 start := time.Now()
 orderBooksCounter.Inc()
 HttpRequestCountWithPath.WithLabelValues(r.URL.Path).Inc()

 w.Write([]byte("Order placed!"))

 duration := time.Since(start).Seconds()
 HttpRequestDuration.WithLabelValues(r.URL.Path).Observe(duration)
}
