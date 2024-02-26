package main

import (
	"gowebbook/config"
	"gowebbook/server"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting runners app")
	log.Println("init config")
	config := config.InitConfig("runners")
	log.Println("Initializing database")
	dbHandler := server.InitDatabase(config)
	log.Println("Initializing HTTP server")
	httpServer := server.InitHttpServer(config, dbHandler)
	httpServer.Start()
}
