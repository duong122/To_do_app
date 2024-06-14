package config

import (
	"database/sql"
	"fmt"
	"server/helper"

	_ "github.com/lib/pq" // Postgres golang driver
	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbName   = "test"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("host = %s port = %s password = %s dbname = %s sslmode = disable",
		host, port, user, password, dbName)

	db, err := sql.Open("postgres", sqlInfo)
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	log.Info().Msg("Connected to database !!")
	return db
}
