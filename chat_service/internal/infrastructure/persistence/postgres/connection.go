package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func GetPgConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	return conn
}
