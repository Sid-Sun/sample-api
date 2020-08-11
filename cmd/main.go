package main

import (
	"github.com/sid-sun/sample-api/cmd/config"
	"github.com/sid-sun/sample-api/pkg/api"
)

func main() {
	cfg := config.Load()
	initLogger(cfg.GetEnv())
	api.StartServer(cfg, logger)
}
