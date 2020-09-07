package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/rmar8138/go-api-test/pkg/config"
	"github.com/rmar8138/go-api-test/pkg/handler"

	m "github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	// get current background context
	ctx := context.Background()

	// set up router
	router := chi.NewRouter()

	// middleware
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(m.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	}))

	// routes
	router.Get("/", handler.Index)

	// make channel to subscribe to os signals
	stopChan := make(chan os.Signal, 1)

	// registers the given channel to receive os/unix notifications
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	srv := &http.Server{Addr: ":" + config.Configuration.Port, Handler: router}

	log.Info("Starting server on port " + config.Configuration.Port)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	<-stopChan // wait for SIGINT

	log.Info("Shutting down server...")

	if err := srv.Shutdown(ctx); err != nil {
		log.Error(err)
	}

	log.Info("Server gracefully stopped")
}
