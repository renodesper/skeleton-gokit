package postgre

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

// NewPostgreClient ...
func NewPostgreClient(username, password, host string, port int, dbName string) *pg.DB {
	addr := fmt.Sprintf("%s:%d", host, port)
	db := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     username,
		Password: password,
		Database: dbName,
	})

	return db
}
