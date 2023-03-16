package flagutil

import (
	"flag"
	"sync"
)

type FlagUtil struct {
	MigrateDB bool
}

func (fu *FlagUtil) Setup() {
	flag.BoolVar(&fu.MigrateDB, "migrateDB", false, "Migrate the database")
	flag.Parse()
}

var (
	fu     *FlagUtil
	fuOnce sync.Once
)

func Get() *FlagUtil {
	if fu == nil {
		fuOnce.Do(func() {
			fu = &FlagUtil{}
		})
	}
	return fu
}
