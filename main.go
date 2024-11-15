package main

import (
	"context"
	"log"

	"github.com/takumin/go-pkl-playground/config"
)

func main() {
	cfg, err := config.LoadFromPath(context.Background(), "config.pkl")
	if err != nil {
		panic(err)
	}
	log.Println("Host:", cfg.Host)
	log.Println("Port:", cfg.Port)
}
