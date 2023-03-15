package main

import (
	"net/http"

	fu "github.com/irahardianto/service-pattern-go/infrastructures/flagutil"
	"github.com/irahardianto/service-pattern-go/infrastructures/logutil"
)

var flagutil fu.FlagUtil

func main() {
	// Setup logs
	logutil.Setup()
	logutil.Info("Running...")

	// Setup flags
	flagutil = fu.FlagUtil{}
	flagutil.Setup()

	http.ListenAndServe(":8080", ChiRouter().InitRouter())
}
