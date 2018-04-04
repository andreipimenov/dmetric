package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

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

	a := &Application{
		Config: c,
	}

	r := NewRouter(a)

	log.Printf("Start listening on port %d", c.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", c.Port), r)
	if err != nil {
		log.Fatal(err)
	}
}
