package main

import (
	"github.com/alexbt/go-consul-example/go-service-one/pkg/public/server"
)

func main() {
	println("Starting server")
	server.StartServer();
}

