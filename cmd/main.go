package main

import (
	"log"
	"os"

	"github.com/Pavel-Sergeev-ekb/first_http_server/internal/server"
)

func main() {
	logger := log.New(
		os.Stdout,
		"[%s] ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC,
	)

	srv := server.NewServer(logger)

	if err := srv.Server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
