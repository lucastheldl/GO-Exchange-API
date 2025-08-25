package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go-api/product/infra"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var Conn *pgx.Conn


func runMigrations(databaseURL string) error {
    m, err := migrate.New("file://migrations", databaseURL)
    if err != nil {
        return err
    }
    defer m.Close()

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return err
    }

    fmt.Println("Migrations applied successfully")
    return nil
}

// função principal
func main() {

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	connString := os.Getenv("DB_URL")
	if connString == "" {
		fmt.Fprintf(os.Stderr, "DB_URL not set\n")
		os.Exit(1)
	}


	err = runMigrations(connString)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Migration failed: %v\n", err)
        os.Exit(1)
    }
	
	ctx := context.Background()

	Conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}


	defer Conn.Close(ctx)
	fmt.Println("Connection established")
	
    router := mux.NewRouter()

	infra.RegisterRoutes(router, Conn)
	
	fmt.Println("🚀 Server running at http://localhost:8000")

	log.Fatal(http.ListenAndServe(":8000", router))
}