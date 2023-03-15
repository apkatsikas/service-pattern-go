package main

import (
	"net/http"
	"sync"

	"github.com/irahardianto/service-pattern-go/infrastructures/flagutil"
	"github.com/irahardianto/service-pattern-go/infrastructures/logutil"
)

// Setup FlagUtil interface and singleton
type IFlagUtil interface {
	Setup()
	Get() *flagutil.FlagUtil
}

var (
	fu     *flagutil.FlagUtil
	fuOnce sync.Once
)

func FlagUtil() IFlagUtil {
	if fu == nil {
		fuOnce.Do(func() {
			fu = &flagutil.FlagUtil{}
		})
	}
	return fu
}

func main() {
	// Setup logs
	logutil.Setup()
	logutil.Info("Running...")

	// Setup flags
	FlagUtil().Setup()

	http.ListenAndServe(":8080", ChiRouter().InitRouter())
}
