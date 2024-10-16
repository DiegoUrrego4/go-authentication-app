package main

import (
	"log"
	"login-app/cmd/bootstrap"
	"login-app/internal/platform/storage/newsql"
)

func main() {
	db, err := newsql.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	srv := bootstrap.NewServer(db)
	srv.Run()
}
