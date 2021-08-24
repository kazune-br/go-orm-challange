package main

import (
	"github.com/kazune-br/go-orm-challange/internal/gateway/gormorm"
	"github.com/kazune-br/go-orm-challange/internal/gateway/sqlboilerorm"
)

func main() {
	gormorm.RunGorm()
	sqlboilerorm.RunSqlboiler()
}
