package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Monitor struct {
}

var (
	CodeCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "total_code_request",
			Help: "total code request",
		})
)

func RegisterMonitor() {
	prometheus.MustRegister(CodeCounter)
}
