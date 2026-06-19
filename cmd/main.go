package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "MORSE-SERVER: ", log.LstdFlags)

	srv := server.NewServer(logger)

	if err := srv.Start(); err != nil {
		logger.Fatalf("Критическая ошибка при работе сервера: %v", err)
	}
}
