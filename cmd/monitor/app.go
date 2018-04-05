package main

import "github.com/andreipimenov/dmetric/model"

type Application struct {
	Config *Config
	Cache  model.Cache
	DB     model.DB
}
