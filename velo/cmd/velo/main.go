package main

import (
	"fmt"
	"log"
	"os"

	"github.com/velo-api/velo/internal/gateway"
	"github.com/velo-api/velo/pkg/config"
)

func main() {
	configPath := "configs/velo.yaml"
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}

	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Printf("🚀 Velo API Gateway starting on %s:%d\n", cfg.Server.Host, cfg.Server.Port)

	gw, err := gateway.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create gateway: %v", err)
	}

	if err := gw.Start(); err != nil {
		log.Fatalf("Gateway failed: %v", err)
	}
}
