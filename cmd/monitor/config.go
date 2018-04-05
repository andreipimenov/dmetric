package main

type Config struct {
	MetricLimits struct {
		Metric1 MetricLimit `json:"metric_1"`
		Metric2 MetricLimit `json:"metric_2"`
		Metric3 MetricLimit `json:"metric_3"`
		Metric4 MetricLimit `json:"metric_4"`
		Metric5 MetricLimit `json:"metric_5"`
	} `json:"metric_limits"`
	Port         int    `json:"port"`
	DBURL        string `json:"db_url"`
	RedisHost    string `json:"redis_host"`
	RedisPort    int    `json:"redis_port"`
	RedisDB      int    `json:"redis_db"`
	SMTPHost     string `json:"smtp_host"`
	SMTPPort     int    `json:"smtp_port"`
	SMTPLogin    string `json:"smtp_login"`
	SMTPPassword string `json:"smtp_password"`
	MailTo       string `json:"mail_to"`
}

type MetricLimit struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
}
