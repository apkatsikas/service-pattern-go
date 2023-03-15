package flagutil

import (
	"flag"
)

type FlagUtil struct {
	MigrateDB bool
}

func (fu *FlagUtil) Setup() {
	flag.BoolVar(&fu.MigrateDB, "migrateDB", false, "Migrate the database")
	flag.Parse()
}

func (fu *FlagUtil) Get() *FlagUtil {
	return fu
}
