package server

import (
	config "github.com/arudji1211/ecommerce/config"
)

func serve() {
	pgConn := config.NewPostgresGormConnection()
}
