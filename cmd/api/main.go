package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"notes-service/data"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var webPort string = "80"

type Config struct {
	Models data.Models
}

func main() {
	// connect to db
	db := connectToDB()

	app := Config{
		Models: data.NewModels(db),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	attempts := 0

	for {
		conn, err := openDB(dsn)

		if err != nil {
			log.Println("Waiting for postgres.")
			attempts++
		} else {
			log.Println("Successfully connected to postgres.")
			return conn
		}

		if attempts > 20 {
			log.Println(err)
			return nil
		}

		time.Sleep(2 * time.Second)
		continue
	}

}

func openDB(dsn string) (*sql.DB, error) {
	// open db connection
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	// check if connection is alive
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
