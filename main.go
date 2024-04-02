package main

import (
	"net/http"
	"text/template"
)

// Colecao de variaveis que formam um novo tipo
type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 59, Quantidade: 5},
		{"Tênis", "Confortável", 89.9, 3},
		{"Fone", "Muito bom", 129.0, 2},
		{"Notebook", "Acer", 5999.0, 4},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
