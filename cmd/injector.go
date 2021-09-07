//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/kazune-br/go-orm-challange/internal/gateway/sqlboilerorm"
)

func InitRegistry() *sqlboilerorm.ComponentRegistry {
	wire.Build(
		sqlboilerorm.NewComponentRegistry,
		sqlboilerorm.NewMySQLConnector,
		sqlboilerorm.NewBoilerRepository)
	return &sqlboilerorm.ComponentRegistry{}
}
