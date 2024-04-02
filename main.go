package main

import (
	"net/http"
	"text/template"

	"github.com/Vicrrs/InventarioDeProdutos/models"
	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("templates/*.html")) // Carrega e compila todos os templates HTML do diretório templates. Panic em caso de erro.

func main() {
	http.HandleFunc("/", index)       // Registra a função index como manipuladora da rota raiz.
	http.ListenAndServe(":8000", nil) // Inicia um servidor HTTP na porta 8000.
}

func index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()  // Chama a função para buscar todos os produtos do banco de dados.
	temp.ExecuteTemplate(w, "Index", todosOsProdutos) // Renderiza o template Index, passando a lista de produtos como dados.
}
