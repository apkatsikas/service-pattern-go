package main

import (
	"net/http"

	"github.com/irahardianto/service-pattern-go/infrastructures/flagutil"
	"github.com/irahardianto/service-pattern-go/infrastructures/logutil"
)

func main() {
	// Setup logs
	logutil.Setup()
	logutil.Info("Running...")

	// Setup flags
	flagutil.Get().Setup()

	http.ListenAndServe(":8080", ChiRouter().InitRouter())
}
