package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// ConectaComBancoDeDados cria e retorna uma conexão com o banco de dados configurado através de variáveis de ambiente.
func ConectaComBancoDeDados() *sql.DB {
	// Recupera as variáveis de ambiente necessárias para a conexão com o banco de dados.
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	sslmode := os.Getenv("DB_SSLMODE")

	// Formata a string de conexão com os valores das variáveis de ambiente.
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", user, dbname, password, host, sslmode)

	// Tenta abrir uma conexão com o banco de dados usando a string de conexão formatada.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error()) // Em caso de erro, encerra a execução com um panic, exibindo a mensagem de erro.
	}
	return db // Retorna o objeto de conexão com o banco de dados se a conexão for bem-sucedida.
}
