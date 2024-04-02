package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // Importação do driver PostgreSQL
)

func ConectaComBancoDeDados() *sql.DB {
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	sslmode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", user, dbname, password, host, sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}
