package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

var (
	database *pgx.Conn
)

func GetDatabase() *pgx.Conn {
	if database == nil {
		Connect()
	}
	return database
}

func Connect() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	database = conn
}
