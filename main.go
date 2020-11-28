package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	termPprof := make(chan bool)
	pprofDone := make(chan bool)
	done := make(chan bool)

	InitPprof(termPprof, pprofDone)

	newRouter := mux.NewRouter()

	server := &http.Server{
		Handler:      newRouter,
		Addr:         "127.0.0.1:15120",
		WriteTimeout: 100 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("=============================================")
		log.Printf("==== Starting Server at: %s =====\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v", server.Addr, err)
		}
	}()

	go func() {
		<-quit
		log.Println("Server is shutting down...")
		termPprof <- true
		<-pprofDone

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Could not gracefully shutdown the server: %v", err)
		}
		close(done)
	}()

	<-done
	log.Println("Server stopped")

}