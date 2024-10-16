package newsql

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

type ConfigPostgresDB struct {
	conn *pgx.Conn
}

func LoadEnvConfig() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("error loading .env file: %w", err)
	}

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	return connStr, nil
}

func NewPostgresDB() (*ConfigPostgresDB, error) {
	connStr, err := LoadEnvConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading env config: %w", err)
	}
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	return &ConfigPostgresDB{conn: conn}, nil
}

func (db *ConfigPostgresDB) Close() {
	db.conn.Close(context.Background())
}
