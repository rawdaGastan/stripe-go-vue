// Package app for c4s backend app
package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/rawdaGastan/stripe-go-vue/config"
	"github.com/rawdaGastan/stripe-go-vue/middlewares"
	"github.com/rawdaGastan/stripe-go-vue/models"
	"github.com/rs/zerolog/log"
	"github.com/stripe/stripe-go/v74"
)

// App for all dependencies of backend server
type App struct {
	config config.Configuration
	db     models.DB
}

// NewApp creates new server app all configurations
func NewApp(ctx context.Context, configFile string) (app *App, err error) {
	config, err := config.ReadConfFile(configFile)
	if err != nil {
		return
	}

	db := models.NewDB()
	err = db.Connect(config.Database.File)
	if err != nil {
		return
	}
	err = db.Migrate()
	if err != nil {
		return
	}

	return &App{
		config: config,
		db:     db,
	}, nil
}

// Start starts the app
func (a *App) Start(ctx context.Context) (err error) {
	stripe.Key = a.config.StripeKeys.Secret

	a.registerHandlers()

	log.Info().Msgf("Server is listening on port %s", a.config.Port)

	srv := &http.Server{
		Addr: a.config.Port,
	}

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("HTTP server error")
		}
		log.Info().Msg("Stopped serving new connections")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal().Err(err).Msg("HTTP shutdown error")
	}
	log.Info().Msg("Graceful shutdown complete")

	return nil
}

func (a *App) registerHandlers() {
	r := mux.NewRouter()

	versionRouter := r.PathPrefix("/" + a.config.Version).Subrouter()

	versionRouter.HandleFunc("/checkout", WrapFunc(a.createCheckoutSession)).Methods("POST", "OPTIONS")
	versionRouter.HandleFunc("/sell", WrapFunc(a.sellProduct)).Methods("PUT", "OPTIONS")
	versionRouter.HandleFunc("/create", WrapFunc(a.createProduct)).Methods("POST", "OPTIONS")
	versionRouter.HandleFunc("/products", WrapFunc(a.getProducts)).Methods("GET", "OPTIONS")

	// middlewares
	r.Use(middlewares.EnableCors)
	http.Handle("/", r)
}
