package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {
	conn, err := pgx.Connect(
		context.Background(),
		"postgresql://postgres:1@143.198.222.47:54322/snagtick_be?sslmode=disable",
	)
	if err != nil {
		fmt.Println(err)
	}
	return conn
}
