package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"user-service/data"

	"github.com/AdamShannag/toolkit/v2"
	_ "github.com/sijms/go-ora/v2"
)

const webPort = "80"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
	tools  *toolkit.Tools
}

func main() {
	log.Println("Starting service")
	tools := toolkit.Tools{}
	// Connect to db
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Oracle")
	}

	// set up cofig
	app := Config{
		DB:     conn,
		Models: data.New(conn),
		tools:  &tools,
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

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("oracle", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Oracle not yet ready...")
			counts++
		} else {
			log.Println("Connected to Oracle!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}
