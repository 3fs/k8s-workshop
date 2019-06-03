package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := flag.String("addr", ":80", "interface and port to listen on")
	flag.Parse()

	// Initialize logger.
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	// Determine the CODE env variable.
	name, ok := os.LookupEnv("CODE")
	if !ok {
		name = "NOT_FOUND"
	}

	// Set up http server.
	http.HandleFunc("/hello", logRequest(logger, helloHandler(name)))
	http.HandleFunc("/ready", logRequest(logger, readyHandler()))
	http.HandleFunc("/live", logRequest(logger, liveHandler(time.Now())))

	// File server.
	fs := http.FileServer(http.Dir("/tmp"))
	http.HandleFunc("/files/", logRequest(logger, http.StripPrefix("/files/", fs).ServeHTTP))

	// A game.
	game := http.FileServer(http.Dir("/game"))
	http.HandleFunc("/", logRequest(logger, game.ServeHTTP))

	logger.Printf("Server is starting on %s", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		logger.Fatal(err)
	}
}

// logRequest is a middleware that logs some request details
func logRequest(logger *log.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Println(r.Method, r.Host, r.URL.Path, r.RemoteAddr, r.UserAgent())
		next.ServeHTTP(w, r)
	}
}

// helloHandler responds to request with given name
func helloHandler(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s\n", name)
	}
}

// readyHandler is used for readiness probe
func readyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok\n")
	}
}

// liveHandler is used for liveness probe and fakes being dead after 20 seconds
func liveHandler(start time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		duration := time.Since(start)
		if duration.Seconds() > 20 {
			w.WriteHeader(500)
			fmt.Fprintf(w, "error: not alive any more\n")
		} else {
			fmt.Fprintf(w, "ok\n")
		}
	}
}
