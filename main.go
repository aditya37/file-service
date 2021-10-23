package main

import (
	"log"

	"github.com/aditya37/file-service/server"
)

func main() {
	serve, err := server.NewHttpServer()
	if err != nil {
		log.Println(err)
	}
	serve.Start()
}
