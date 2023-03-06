package interfaces

type IDbHandler interface {
	Query(statement string) (IRow, error)
	// Generics? Could return an entity of type T
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
}
