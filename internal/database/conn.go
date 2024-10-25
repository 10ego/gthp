package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// DB represents a database connection pool
type DB struct {
	*pgxpool.Pool
	log *zap.SugaredLogger
}

// Connect establishes a connection to the database and returns a DB instance
func Connect(databaseURL string, log *zap.SugaredLogger) (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database URL: %v", err)
	}

	// You can set additional pool configuration here if needed
	// For example:
	// config.MaxConns = 10
	// config.MinConns = 2

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	// Verify the connection
	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	return &DB{Pool: pool, log: log}, nil
}

// Close closes the database connection pool
func (db *DB) Close() {
	db.Pool.Close()
}

// Example query method
func (db *DB) GetUserByID(ctx context.Context, id int) (string, error) {
	var name string
	err := db.QueryRow(ctx, "SELECT name FROM users WHERE id = $1", id).Scan(&name)
	if err != nil {
		db.log.Errorw("Error querying user", "error", err, "id", id)
		return "", fmt.Errorf("error querying user: %v", err)
	}
	return name, nil
}
