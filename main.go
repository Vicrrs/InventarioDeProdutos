package main

import (
	"log"
	"net/http"

	"github.com/Vicrrs/InventarioDeProdutos/routes"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load() // Esta função procura o arquivo .env na raiz do diretório atual
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	routes.CarregaRotas()
	// Inicia um servidor HTTP na porta 8000.
	log.Fatal(http.ListenAndServe(":8000", nil))
}
