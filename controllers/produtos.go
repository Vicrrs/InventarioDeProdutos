package controllers

import (
	"html/template"
	"net/http"

	"github.com/Vicrrs/InventarioDeProdutos/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html")) // Carrega e compila todos os templates HTML do diretório templates. Panic em caso de erro.

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()  // Chama a função para buscar todos os produtos do banco de dados.
	temp.ExecuteTemplate(w, "Index", todosOsProdutos) // Renderiza o template Index, passando a lista de produtos como dados.
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}
