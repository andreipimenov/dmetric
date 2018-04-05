package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/andreipimenov/dmetric/driver/mail"
	"github.com/andreipimenov/dmetric/driver/postgres"
	"github.com/andreipimenov/dmetric/driver/redis"
	"github.com/andreipimenov/goto/config"
)

func main() {

	cfgFile := flag.String("config", "data/config.json", "configuration file")

	cfgDriver := config.NewFileConfig(*cfgFile)
	c := &Config{}
	err := cfgDriver.GetJSON(c)
	if err != nil {
		log.Fatal(err)
	}

	redis, err := redis.NewRedis(fmt.Sprintf("%s:%d", c.RedisHost, c.RedisPort), c.RedisDB)
	if err != nil {
		log.Fatal(err)
	}

	postgres, err := postgres.NewPostgres(c.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	mail := mail.NewMail(c.SMTPHost, c.SMTPPort, c.SMTPLogin, c.SMTPPassword)

	a := &Application{
		Config:   c,
		Cache:    redis,
		DB:       postgres,
		Notifier: mail,
	}

	r := NewRouter(a)

	log.Printf("Start listening on port %d", c.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", c.Port), r)
	if err != nil {
		log.Fatal(err)
	}
}
