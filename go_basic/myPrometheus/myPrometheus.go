/*
 * @Author: your name
 * @Date: 2022-04-27 19:20:04
 * @LastEditTime: 2022-04-27 19:45:44
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /Go/go_basic/prometheus/promethus.go
 */

package myPrometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	http_request_total = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "the total number of processed http requests",
	})
)

func init() {
	prometheus.MustRegister(http_request_total)
}

func RealizePrometheus() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		http_request_total.Inc()
	})

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":12012", nil)
}

func WithDefaultIndex() {

}
