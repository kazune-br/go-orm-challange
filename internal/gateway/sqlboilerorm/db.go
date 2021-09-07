package sqlboilerorm

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var _ RDBConnector = (*rdbConnector)(nil)

type (
	RDBConnector interface {
		GetDB() *sql.DB
		GetContext() context.Context
		Close()
	}

	rdbConnector struct {
		DB      *sql.DB
		Context context.Context
	}
)

func NewMySQLConnector() RDBConnector {
	dns := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?parseTime=true",
		"root",
		"password",
		"localhost",
		"4402",
		"sqlboilerdb",
	)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}

	return &rdbConnector{
		DB:      db,
		Context: context.Background(),
	}
}

func (dbc *rdbConnector) GetDB() *sql.DB {
	return dbc.DB
}

func (dbc *rdbConnector) GetContext() context.Context {
	return dbc.Context
}

func (dbc *rdbConnector) Close() {
	dbc.DB.Close()
}
