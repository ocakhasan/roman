package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ocakhasan/roman/pkg/app"
	"github.com/ocakhasan/roman/pkg/handler"

	"github.com/sirupsen/logrus"
)

func main() {
	var (
		l   = logrus.New() // Initialize the logger
		ctx = context.Background()
	)

	httpPort := os.Getenv("PORT") // read the env
	if httpPort == "" {
		httpPort = "8080"
	}

	app := app.New(l)
	router := handler.CreateHandler(app)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", httpPort),
		Handler: router,
	}

	// Implement the graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			l.Fatalf("server error %s", err)
		}
	}()

	l.Infof("running on http://localhost:%v", httpPort)
	<-done
	l.Infof("server stopped")
	if err := server.Shutdown(ctx); err != nil {
		l.Fatalf("Server shutdown failed %+v", err)
	}

	l.Info("server shutdowned gracefully")
}
