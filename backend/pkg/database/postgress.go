package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(dbURL string) *pgxpool.Pool {
	pool, err := pgxpool.New(
		context.Background(),
		dbURL,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = pool.Ping(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database Connected")

	return pool
}
