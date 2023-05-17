package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "test"
)

func NewDB(ctx context.Context) (*Database, error) {
	dsn := generateDsn()
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return newDataBase(pool), nil
}

func generateDsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}
