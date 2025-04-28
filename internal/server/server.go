package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Pavel-Sergeev-ekb/first_http_server/internal/handlers"
)

type AppServer struct {
	Logger *log.Logger
	Server *http.Server
}

func NewServer(logger *log.Logger) *AppServer {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.BaseHandle)

	mux.HandleFunc("/upload", handlers.UploadHandle)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux, //роутер
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &AppServer{
		Logger: logger,
		Server: server,
	}
}
