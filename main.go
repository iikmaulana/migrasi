package main

import (
	"github.com/iikmaulana/migrasi/config"
	"github.com/joho/godotenv"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	config.Init()

}
