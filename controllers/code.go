package controllers

import (
	"math"

	"github.com/astaxie/beego"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

var (
	codeCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_request_total_code",
			Help: "total request code controller",
		},
	)

	temps = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "pond_temperature_celsius",
		Help:       "The temperature of the frog pond.",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})
)

type CodeController struct {
	beego.Controller
}

func init() {
	prometheus.MustRegister(codeCounter)
	prometheus.MustRegister(temps)

}

func (c *CodeController) Get() {
	codeCounter.Inc()

	// Simulate some observations.
	for i := 0; i < 1000; i++ {
		temps.Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10)
	}
	metric := &dto.Metric{}
	temps.Write(metric)
	c.TplName = "code.tpl"

}
