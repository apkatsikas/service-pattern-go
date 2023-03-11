package main

import (
	"flag"
	"net/http"

	"github.com/irahardianto/service-pattern-go/logutil"
)

const logFile = "logs.txt"

var migrateDB = false

func main() {
	logutil.Setup()
	logutil.Info("Running...")
	flag.BoolVar(&migrateDB, "migrateDB", false, "Migrate the database")
	flag.Parse()

	http.ListenAndServe(":8080", ChiRouter().InitRouter())
}
