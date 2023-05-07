package main

import (
	"fmt"
	"skillbox-diploma/internal/config"
	"skillbox-diploma/internal/simulator"
)

func main() {
	cfg := config.Get()

	simulator.Start(fmt.Sprintf("%s:%d", cfg.Simulator.Host, cfg.Simulator.Port))
}
