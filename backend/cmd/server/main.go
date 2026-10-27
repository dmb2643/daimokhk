package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/dmb2643/daimokhk/backend/internal/listing"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)
	return mux
}

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "critical error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	r := newRouter()

	dsn := os.Getenv("POSTGRESQL_STRING")
	if dsn == "" {
		return fmt.Errorf("POSTGRESQL_STRING is required")
	}

	connectCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	pool, err := openDB(connectCtx, dsn)

	cancel()
	if err != nil {
		return err
	}
	defer pool.Close()

	listingRepo := listing.NewPostgreSQLRepo(pool)
	listingService := listing.NewService(*listingRepo)
	listingHandler := listing.NewHandler(*listingService)

	r.HandleFunc("GET /listings", listingHandler.GetListings)
	r.HandleFunc("GET /listings/{id}", listingHandler.GetListingByID)

	slog.Info("runing")
	if err := http.ListenAndServe("127.0.0.1:8080", r); err != nil {
		return err
	}

	return nil
}
