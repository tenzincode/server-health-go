package health

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	CPUUsageGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_usage",
		Help: "Current CPU usage",
	})
	MemoryUsageGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memory_usage",
		Help: "Current memory usage",
	})
)

func init() {
	prometheus.MustRegister(CPUUsageGauge)
	prometheus.MustRegister(MemoryUsageGauge)
}
