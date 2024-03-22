package configs

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitPostgreSQL(databaseUrl string) *pgxpool.Conn {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	config.MaxConnIdleTime = 5 * time.Minute
	config.MaxConnLifetime = 5 * time.Minute
	config.MaxConns = 10

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		for i := 0; i < 10; i++ {
			conn, err = pool.Acquire(context.Background())
			if err == nil {
				break
			}
			time.Sleep(3 * time.Second)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
	}

	return conn
}
