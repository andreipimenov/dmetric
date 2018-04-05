package main

type Config struct {
	Port      int    `json:"port"`
	DBURL     string `json:"db_url"`
	RedisHost string `json:"redis_host"`
	RedisPort int    `json:"redis_port"`
	RedisDB   int    `json:"redis_db"`
}
