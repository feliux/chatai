package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/feliux/chatai/handler"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/subosito/gotenv"
)

const (
	LISTEN_ADDR string = "LISTEN_ADDR"
)

func init() {
	if err := gotenv.Load(); err != nil {
		log.Fatalf("error ocurred reading .env file %s", err)
	}
}

func main() {
	router := chi.NewMux()
	router.Use(middleware.Logger)
	// Serve the public filesystem
	router.Handle("/*", public())
	router.Handle("/ws", handler.Make(handler.HandleChat))
	// GET
	router.Get("/", handler.Make(handler.HandleHomeIndex))
	router.Get("/chat", handler.Make(handler.HandleChatIndex))

	listenAddr := os.Getenv(LISTEN_ADDR)
	slog.Info("application running", "addr", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
