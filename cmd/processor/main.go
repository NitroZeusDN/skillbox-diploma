package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"skillbox-diploma/internal/config"
	"skillbox-diploma/internal/handler"
	"skillbox-diploma/internal/server"
	"syscall"
)

func main() {
	cfg := config.Get()
	ctx := context.Background()

	// инициализруем обработчик запросов.
	h := handler.New(cfg)

	// инициализируем HTTP сервер и запускаем его в отдельной горутине.
	srv := server.New(cfg.Server, h.Routes())
	go func() {
		if err := srv.Start(ctx); err != nil {
			log.Fatalf("httpServer working failed: %s", err)
		}
	}()
	log.Println("server has been started")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// блокируем дальнейшую обработку до тех пор,
	// пока не придет нужный сигнал завершения.
	<-done

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server.Shutdown failed: %s", err)
	}

	log.Println("server has been stopped")
}
