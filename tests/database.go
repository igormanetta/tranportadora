package tests

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/avast/retry-go"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
)

func waitForDB(dbpool *pgxpool.Pool) error {
	return retry.Do(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := dbpool.Ping(ctx)
		if err != nil {
			return err
		}
		return nil
	}, retry.Attempts(5), retry.Delay(2*time.Second))
}

func SetupTestDB(c testcontainers.Container) *pgxpool.Pool {
	ctx := context.Background()
	host, err := c.Host(ctx)
	if err != nil {
		log.Fatalf("Failed to get container host: %v", err)
	}

	port, err := c.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("Failed to get container port: %v", err)
	}

	dsn := "postgres://test:test@%s:%s/testdb?sslmode=disable"
	dbpool, err := pgxpool.New(ctx, fmt.Sprintf(dsn, host, port.Port()))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := waitForDB(dbpool); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return dbpool
}
