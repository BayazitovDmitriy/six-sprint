package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Serv struct {
	lg       *log.Logger
	httpServ *http.Server
}

func (s Serv) ServStart() error {
	return s.httpServ.ListenAndServe()
}
func CreateRouter(l *log.Logger) *Serv {
	rt := newRouter()

	servStr := &http.Server{
		Addr:         ":8080",
		Handler:      rt,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Serv{
		lg:       l,
		httpServ: servStr,
	}
}

func newRouter() *http.ServeMux {
	rt := http.NewServeMux()
	rt.HandleFunc("/", handlers.HtmlReturn)
	rt.HandleFunc("/upload", handlers.HtmlParse)
	return rt
}
