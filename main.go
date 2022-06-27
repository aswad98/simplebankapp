package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/minibank/api"
	db "github.com/minibank/db/sqlc"
	"github.com/minibank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connected to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start error", err)
	}
}
