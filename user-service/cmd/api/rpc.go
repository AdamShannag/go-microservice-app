package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"user-service/data"
)

type RPCServer struct {
	DB     *sql.DB
	Models data.Models
}

type RPCUserPayload struct {
	ID      string
	Name    string
	Address string
}

func (r *RPCServer) CreateUser(payload RPCUserPayload, resp *string) error {
	user := r.Models.User
	user.ID = payload.ID
	user.Name = payload.Name
	user.Address = payload.Address
	log.Println("here ----------------- 1 ok")
	_, err := user.Insert(user)
	if err != nil {
		log.Println("error creating a user", err)
		return err
	}
	log.Println("here ----------------- 2 ok")

	*resp = "User created via RPC:" + payload.Name
	return nil
}

func (app *Config) rpcListen() error {
	log.Println("Starting RPC server on port ", rpcPort)
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
