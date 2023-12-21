package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const address = ":1970"

func main() {
	http.HandleFunc("/", logging(handler))
	http.HandleFunc("/health", health)
	server := http.Server{Addr: address}
	go func() {
		log.Println("Server starting up...")
		log.Fatal(server.ListenAndServe())
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	log.Println("Server shutting down...")
	server.Shutdown(context.Background())
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from a rainy Utrecht")
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		f(w, r)
	}
}
