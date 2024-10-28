package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// Command line arguments
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// router
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// structured logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("starting server", "addr", *addr)
	logger.Info("started server", "link", "http://localhost:4000/")

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
