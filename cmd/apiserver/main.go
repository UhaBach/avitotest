package main

import (
	"avitotest/internal/app/apiserver"
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	cfg := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, cfg)
	if err != nil{
		log.Fatal(err)
	}

	s := apiserver.New(cfg)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
