package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create the handlers
	ph := handlers.NewProducts(l)
	

	// create a new serve mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)
	

	// create a new server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,                     // set the default handler
		IdleTimeout:  120 * time.Second,     // max time for connections using TCP Keep-Alive
		ReadTimeout:  1 * time.Second,        // max time to read request from the client
		WriteTimeout: 1 * time.Second,       // max time to write response to the client
	}

	// start the server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)


	// Block until a signal is received.
	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

	err := http.ListenAndServe(":9090", sm)
	log.Fatal(err)
}
