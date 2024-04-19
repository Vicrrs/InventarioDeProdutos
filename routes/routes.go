package routes

import (
	"net/http"

	"github.com/Vicrrs/InventarioDeProdutos/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
