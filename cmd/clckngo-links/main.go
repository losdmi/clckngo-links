package main

import (
	portsHttp "clckngo/links/internal/ports/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
)

func main() {
	var server portsHttp.Server

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	handler := portsHttp.HandlerFromMux(server, router)

	port := os.Getenv("PORT")
	err := http.ListenAndServe(
		"localhost:"+port,
		handler,
	)
	if err != nil {
		return
	}
}