package interfaces

import "gorm.io/gorm"

type IDbHandler interface {
	Connection() *gorm.DB
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
}
