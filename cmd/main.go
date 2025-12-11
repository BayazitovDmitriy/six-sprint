package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	lg := log.New(log.Writer(), "serv:", log.Ldate|log.Ltime)
	serv := server.CreateRouter(lg)
	if err := serv.ServStart(); err != nil {
		lg.Fatal("Server error:", err)
	}
}
