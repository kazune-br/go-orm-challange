package main

import (
	"github.com/kazune-br/go-orm-challange/internal/gateway/sqlboilerorm"
)

func main() {
	// gormorm.RunGorm()
	registry := InitRegistry()
	defer registry.DBConn.Close()
	sqlboilerorm.Run(registry)
}