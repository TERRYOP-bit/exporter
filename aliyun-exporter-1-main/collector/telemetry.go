package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestsStats = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "aliyun_exporter_requests_total",
			Help: "The total number of cloudmonitor requests sent to Aliyun.",
		},
	)

	responseError = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "aliyun_exporter_error_response_total",
			Help: "The total number response from Aliyun returned with error.",
		},
	)

	responseFormatError = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "aliyun_exporter_error_response_format_total",
			Help: "The total number response from Aliyun with format error.",
		},
	)
)

func init() {
	prometheus.MustRegister(requestsStats)
	prometheus.MustRegister(responseError)
	prometheus.MustRegister(responseFormatError)
}
