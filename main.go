package main

import (
	"net/http"

	"github.com/Vicrrs/InventarioDeProdutos/routes"
	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()
	// Registra a função index como manipuladora da rota raiz.
	http.ListenAndServe(":8000", nil) // Inicia um servidor HTTP na porta 8000.
}
