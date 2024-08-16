package main

import (
	"log"
	"server/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatal("problem in intializing conn: s", err)
	}
}
