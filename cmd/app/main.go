package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"products-cdc/domain"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := pgxpool.New(ctx, "postgres://postgres:postgres@postgres:5432/products")
	if err != nil {
		panic(err)
	}

	router := http.NewServeMux()
	router.HandleFunc("POST /products", createProductHandler(logger, db))

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	go func() {
		logger.Info("Running server", slog.String("addr", server.Addr))
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Server error", slog.Any("error", err))
		}

		logger.Info("Server stopped")
	}()

	<-ctx.Done()

	shudownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shudownCtx); err != nil {
		logger.Error("Server shutdown error", slog.Any("error", err))
	}

	logger.Info("Server shutdown")
}

func createProductHandler(logger *slog.Logger, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.NewV7()
		if err != nil {
			logger.Error("Failed to generate UUID", slog.Any("error", err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		product := domain.Product{
			ID:       id,
			Name:     gofakeit.ProductName(),
			UPC:      gofakeit.ProductUPC(),
			Price:    decimal.NewFromFloat(gofakeit.Price(0, 1000)),
			Quantity: gofakeit.UintN(100),
		}

		result, err := db.Exec(r.Context(), `INSERT INTO products (id, name, upc, price, quantity) VALUES ($1, $2, $3, $4, $5)`, product.ID, product.Name, product.UPC, product.Price, product.Quantity)
		if err != nil {
			logger.Error("Failed to insert product", slog.Any("error", err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if result.RowsAffected() != 1 {
			logger.Error("Failed to insert product", slog.Any("error", err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			logger.Error("Failed to encode product", slog.Any("error", err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
