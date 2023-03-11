package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

const logFile = "logs.txt"

var migrateDB = false

func main() {
	// Delete log
	err := os.Remove(logFile)
	if err != nil {
		log.Fatal(err)
	}
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)

	log.Println("Running...")
	flag.BoolVar(&migrateDB, "migrateDB", false, "Migrate the database")
	flag.Parse()

	http.ListenAndServe(":8080", ChiRouter().InitRouter())
}
