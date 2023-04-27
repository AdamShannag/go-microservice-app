package main

import (
	"context"
	"github.com/AdamShannag/go-microservice-app/auth-service/api"
	"github.com/go-rel/postgres"
	"github.com/go-rel/rel"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

const port = "80"

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "main")))
	shutdowns []func() error
)

func main() {

	var (
		ctx        = context.Background()
		repository = initRepository(ctx)
		mux        = api.NewMux(repository)
		server     = http.Server{
			Addr:    ":" + port,
			Handler: mux,
		}
		shutdown = make(chan struct{})
	)

	go gracefulShutdown(ctx, &server, shutdown)

	logger.Info("server starting: http://localhost" + server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatal("server error", zap.Error(err))
	}

	<-shutdown
}

func initRepository(ctx context.Context) rel.Repository {
	var (
		logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "repository")))
		dsn       = os.Getenv("DSN")
	)

	adapter, err := postgres.Open(dsn)
	if err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
		log.Panic(err)
	}
	err = adapter.Ping(ctx)
	if err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
		log.Panic(err)
	}
	// add to graceful shutdown list.
	shutdowns = append(shutdowns, adapter.Close)

	repository := rel.New(adapter)
	repository.Instrumentation(func(ctx context.Context, op string, message string, args ...interface{}) func(err error) {
		// no op for rel functions.
		if strings.HasPrefix(op, "rel-") {
			return func(error) {}
		}

		t := time.Now()

		return func(err error) {
			duration := time.Since(t)
			if err != nil {
				logger.Error(message, zap.Error(err), zap.Duration("duration", duration), zap.String("operation", op))
			} else {
				logger.Info(message, zap.Duration("duration", duration), zap.String("operation", op))
			}
		}
	})

	return repository
}

func gracefulShutdown(ctx context.Context, server *http.Server, shutdown chan struct{}) {
	var (
		sigint = make(chan os.Signal, 1)
	)

	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	<-sigint

	logger.Info("shutting down server gracefully")

	// stop receiving any request.
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("shutdown error", zap.Error(err))
	}

	// close any other modules.
	for i := range shutdowns {
		shutdowns[i]()
	}

	close(shutdown)
}
