package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"user-service/data"

	"github.com/AdamShannag/toolkit/v2"
	_ "github.com/sijms/go-ora/v2"
)

const (
	webPort = "80"
	rpcPort = "5001"
)

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

	models := data.New(conn)
	// set up cofig
	app := Config{
		DB:     conn,
		Models: models,
		tools:  &tools,
	}

	// set up rpc server
	rpcServer := &RPCServer{
		DB:     conn,
		Models: models,
	}

	err := rpc.Register(rpcServer)

	if err != nil {
		log.Fatal("failed to register RPCServer")
	}

	go app.rpcListen()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
