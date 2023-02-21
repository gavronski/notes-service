package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"notes-service/data"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	webPort string = "80"
	rpcPort string = "5001"
)

type Config struct {
	Models data.Models
}

var app Config

func main() {
	// connect to db
	db := connectToDB()

	app = Config{
		Models: data.NewModels(db),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// register RPC Server
	err := rpc.Register(new(RPCServer))

	go app.rpcListen()

	err = srv.ListenAndServe()

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

// openDB creates postgres connection
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

// rpcListen runs RPC server
func (app *Config) rpcListen() error {
	log.Println("Starting listening RPC server on port ", rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))

	if err != nil {
		return err
	}

	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(rpcConn)
	}
}
