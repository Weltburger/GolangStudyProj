package main

import (
	"forRoma/pkg/server"
	"log"
	"runtime"
)

func main() {
	server, err := server.New("./config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
