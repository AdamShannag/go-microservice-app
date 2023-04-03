package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AdamShannag/toolkit/v2"
)

const webPort = "80"

type Config struct {
	tools *toolkit.Tools
}

func main() {
	tools := toolkit.Tools{}
	app := Config{&tools}

	log.Printf("Starting broker service on port %s\n", webPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
