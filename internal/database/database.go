package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/crisnet-dev/fastfood/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func SetUpDB() error {

	db, err := sql.Open("pgx", config.GetEnv().DATABASE_URL)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	_, err = db.ExecContext(
		ctx,
		`
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			product_name TEXT NOT NULL,
			price REAL NOT NULL,
			image_url TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS pending_orders (
			id SERIAL PRIMARY KEY,
			orders JSONB
		);
	`,
	)
	if err != nil {
		return err
	}

	DB = db

	return nil
}
