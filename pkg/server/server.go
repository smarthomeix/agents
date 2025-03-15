package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/smarthomeix/agents/pkg/director"
	"github.com/smarthomeix/agents/pkg/http/router"
	"github.com/smarthomeix/agents/pkg/service"
	"github.com/smarthomeix/pkg/mqtt/broker"
)

func NewServer(svc service.ServiceInterface) {
	port := flag.String("port", "8001", "API server port")

	brokerHost := flag.String("broker", "tcp://localhost:1883", "mqtt broker host")

	flag.Parse()

	// Initialize MQTT broker connection
	// We'll handle broker shutdown in the graceful shutdown process.
	clientID := fmt.Sprintf("smarthomeix:agent:%s", *port)

	bkr := broker.NewBroker(*brokerHost, clientID)

	// Create a context that cancels on SIGINT/SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	dir := director.NewDirector(svc, bkr)

	// Start API server in a goroutine and capture the server for graceful shutdown.
	server := router.NewServer(*port, dir)

	go func() {
		log.Printf("Starting API server on port %s...", *port)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("API server failed on port %s: %v", *port, err)
		}
	}()

	// Block until a signal is received.
	<-ctx.Done()
	log.Println("Shutting down gracefully...")

	// Create a timeout context for shutdown procedures.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully shutdown the HTTP server.
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}

	bkr.Disconnect()

	log.Println("Shutdown complete.")
}
