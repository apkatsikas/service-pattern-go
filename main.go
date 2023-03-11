package main

import (
	"flag"
	"fmt"
	"net/http"
)

var migrateDB = false

func main() {
	fmt.Println("Running...")
	flag.BoolVar(&migrateDB, "migrateDB", false, "Migrate the database")
	flag.Parse()

	http.ListenAndServe(":8080", ChiRouter().InitRouter())
}
