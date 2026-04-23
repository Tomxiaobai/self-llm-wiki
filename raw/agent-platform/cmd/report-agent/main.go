package main

import (
	"log"

	"agent-platform/internal/shared/config"
	"agent-platform/internal/shared/server"
)

func main() {
	cfg := config.Load()
	if err := server.ServeHealth(cfg.HTTP.ReportAddr, "report-agent"); err != nil {
		log.Fatal(err)
	}
}
