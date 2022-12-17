package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func PostgresInit(username, password, host, name, port string, env ...string) *sqlx.DB {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, name)
	if len(env) == 0 {
		log.Println("env not stated, will be use local env by default")
		connStr += "?sslmode=disable"
	} else {
		if env[0] == "production" {
			connStr += "?sslmode=enable"
		} else {
			connStr += "?sslmode=disable"
		}
	}

	// Connect to database
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal("cannot connect to postgres: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("cannot connect to postgres: ", err)
	}

	log.Println("connection to postgres successfully")
	return db
}
