package model

import "gopkg.in/mgo.v2/bson"

type Product struct {
	ID                   bson.ObjectId `json:"_id,omitempty"bson:"_id,omitempty"`
	NomeProduto          string        `json:"nome_produto"`
	ValorProduto         string        `json:"valor_produto"`
	QuantidadeEstoque    string        `json:"quantidade_estoque"`
	Ativo                string        `json:"ativo"`
	DataExpiracaoAnuncio string        `json:"data_expiracao_anuncio"`
	FotoProduto          string        `json:"foto_produto"`
	CodigoBarras         string        `json:"codigo_barra"`
	Coordenadas          string        `json:"coodernadas"`
	XywCategoria         string        `json:"xyw_categoria"`
	Xyw                  string        `json:"xyw"`
}
