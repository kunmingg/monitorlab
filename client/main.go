package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net"
	"net/http"
	"time"
)

var (
	// Counter metrics
	// num of requests counter vec
	// status field has values: {"OK", "UNKNOWN", "INTERNAL", "INVALID_ARGUMENT"}
	heartbeatCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "heartbeat_counter",
			Help: "Number of requests for deployments",
		},
		[]string{"status"},
	)

	// Gauge metrics
	heartbeatGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "heartbeat_gauge",
		Help: "Number of requests for deployments",
	})
)

func init() {
	// Register prometheus counters
	prometheus.MustRegister(heartbeatCounter)
	prometheus.MustRegister(heartbeatGauge)
}

// Add heartbeat every 10 seconds
func countHeartbeat() {
	for {
		time.Sleep(5 * time.Second)
		heartbeatCounter.WithLabelValues("tag").Inc()
		randNum := rand.Intn(10)
		if randNum > 3 {
			heartbeatGauge.Add(float64(randNum))
		} else {
			heartbeatGauge.Sub(float64(randNum))
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	// add an http handler for prometheus metrics
	http.Handle("/metrics", promhttp.Handler())

	go countHeartbeat()

	if err = http.Serve(listener, nil); err != nil {
		panic(err)
	}
}