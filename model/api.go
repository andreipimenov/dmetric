package model

import "time"

//APIErrors contains all errors responsed by server
type APIErrors struct {
	Errors []*APIMessage `json:"errors"`
}

//APIMessage - common server response with code and message
type APIMessage struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

//APIMetrics - metrics from device
type APIMetrics struct {
	Metric1   interface{} `json:"metric_1"`
	Metric2   interface{} `json:"metric_2"`
	Metric3   interface{} `json:"metric_3"`
	Metric4   interface{} `json:"metric_4"`
	Metric5   interface{} `json:"metric_5"`
	LocalTime *time.Time  `json:"local_time"`
}
