package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gordonBusyman/lpsh_api/api"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}

	dbPath := os.Getenv("LPSH_DB")
	if dbPath == "" {
		panic("LPSH_DB environment variable not set")
	}

	port := os.Getenv("LPSH_API_PORT")
	if port == "" {
		panic("LPSH_API_PORT environment variable not set")
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	a := api.NewAPI(db)

	r.Get("/settings/{code}", a.Settings)

	fmt.Println("server listening on 8080")
	if err := http.ListenAndServe(":"+port, r); err != nil {
		panic(err)
	}
}
