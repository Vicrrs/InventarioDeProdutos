package models

import "github.com/Vicrrs/InventarioDeProdutos/db"

// Produto define a estrutura de dados para um produto, com campos para id, nome, descrição, preço e quantidade.
type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// BuscaTodosOsProdutos recupera todos os produtos do banco de dados e retorna uma slice de Produto.
func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados() // Chama a função para conectar com o banco de dados.
	defer db.Close()                  // Assegura que a conexão com o banco será fechada ao final da função.

	selecDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos") // Executa uma query SQL para selecionar todos os produtos.
	if err != nil {
		panic(err.Error()) // Em caso de erro na query, encerra a execução com um panic.
	}

	p := Produto{}          // Cria uma instância vazia de Produto.
	produtos := []Produto{} // Cria uma slice vazia de Produto para armazenar os produtos recuperados.

	for selecDeTodosOsProdutos.Next() { // Itera sobre os resultados da query.
		var id, quantidade int
		var nome, descricao string
		var preco float64

		// Preenche as variáveis com os dados de um produto, usando o método Scan.
		err := selecDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error()) // Em caso de erro ao escanear, encerra com um panic.
		}

		// Atribui os valores escaneados às propriedades do produto p.
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p) // Adiciona o produto p à slice de produtos.
	}
	return produtos // Retorna a slice de produtos.
}
