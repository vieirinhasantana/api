package middleware

import (
	"fmt"

	"developer.patronus.io/mobile_backend/app/model"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func SearchDataMongo(username, password string) bool {
	clientMongo, err := mgo.Dial("localhost:27017/mobile")
	if err != err {
		fmt.Print(err)
	}
	defer clientMongo.Close()
	clientMongo.SetMode(mgo.Monotonic, true)
	peopleCollection := clientMongo.DB("patronus-mgo").C("users")
	result := model.User{}
	err = peopleCollection.Find(bson.M{"username": username}).One(&result)
	if err != nil {
		fmt.Print("No data found")
		return false
	}

	err_password := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	return err_password == nil

}

func ListAllCategories() []model.Category {

	clientMongo, err := mgo.Dial("localhost:27017/mobile")
	if err != err {
		fmt.Print(err)
	}
	defer clientMongo.Close()
	clientMongo.SetMode(mgo.Monotonic, true)
	categoryCollection := clientMongo.DB("patronus-mgo").C("categories")
	var result []model.Category
	err = categoryCollection.Find(nil).All(&result)
	return result
}

func PostCategory(descricao, xyw string) bool {
	clientMongo, err := mgo.Dial("localhost:27017/mobile")
	if err != err {
		fmt.Print(err)
	}
	defer clientMongo.Close()
	clientMongo.SetMode(mgo.Monotonic, true)
	categoryCollection := clientMongo.DB("patronus-mgo").C("categories")
	err = categoryCollection.Insert(bson.M{"descricao": descricao, "xyw": xyw})
	if err != nil {
		return false
	}
	return true
}

func PutCategory(descricao, xyw, id string) bool {
	clientMongo, err := mgo.Dial("localhost:27017/mobile")
	if err != err {
		fmt.Print(err)
	}
	defer clientMongo.Close()
	clientMongo.SetMode(mgo.Monotonic, true)
	categoryCollection := clientMongo.DB("patronus-mgo").C("categories")
	err = categoryCollection.Update(bson.M{"xyw": id}, bson.M{"descricao": descricao, "xyw": xyw})
	if err != nil {
		return false
	}
	return true
}

func DeleteCategory(id string) bool {
	clientMongo, err := mgo.Dial("localhost:27017/mobile")
	if err != err {
		fmt.Print(err)
	}
	defer clientMongo.Close()
	clientMongo.SetMode(mgo.Monotonic, true)
	categoryCollection := clientMongo.DB("patronus-mgo").C("categories")
	err = categoryCollection.Remove(bson.M{"xyw": id})
	if err != nil {
		return false
	}
	return true
}

func ListAllProducts() []model.Product {

	clientMongo, err := mgo.Dial("localhost:27017/mobile")
	if err != err {
		fmt.Print(err)
	}
	defer clientMongo.Close()
	clientMongo.SetMode(mgo.Monotonic, true)
	productCollection := clientMongo.DB("patronus-mgo").C("products")
	var result []model.Product
	err = productCollection.Find(nil).All(&result)
	return result
}

func PostProduct(nomeProduto, valorProduto, quantidadeEstoque, ativo, dataExpiracaoAnuncio, fotoProduto, codigoBarras, coordenadas, xywCategoria, xyw string) bool {
	clientMongo, err := mgo.Dial("localhost:27017/mobile")
	if err != err {
		fmt.Print(err)
	}
	defer clientMongo.Close()
	clientMongo.SetMode(mgo.Monotonic, true)
	productCollection := clientMongo.DB("patronus-mgo").C("products")
	err = productCollection.Insert(bson.M{
		"nome_produto":           nomeProduto,
		"valor_produto":          valorProduto,
		"quantidade_estoque":     quantidadeEstoque,
		"ativo":                  ativo,
		"data_expiracao_anuncio": dataExpiracaoAnuncio,
		"foto_produto":           fotoProduto,
		"codigo_barra":           codigoBarras,
		"coodernadas":            coordenadas,
		"xyw_categoria":          xywCategoria,
		"xyw":                    xyw})
	if err != nil {
		return false
	}
	return true
}

func PutProduct(nomeProduto, valorProduto, quantidadeEstoque, ativo, dataExpiracaoAnuncio, fotoProduto, codigoBarras, coordenadas, xywCategoria, xyw, id string) bool {
	clientMongo, err := mgo.Dial("localhost:27017/mobile")
	if err != err {
		fmt.Print(err)
	}
	defer clientMongo.Close()
	clientMongo.SetMode(mgo.Monotonic, true)
	productCollection := clientMongo.DB("patronus-mgo").C("products")
	err = productCollection.Update(bson.M{"xyw": id}, bson.M{
		"nome_produto":           nomeProduto,
		"valor_produto":          valorProduto,
		"quantidade_estoque":     quantidadeEstoque,
		"ativo":                  ativo,
		"data_expiracao_anuncio": dataExpiracaoAnuncio,
		"foto_produto":           fotoProduto,
		"codigo_barra":           codigoBarras,
		"coodernadas":            coordenadas,
		"xyw_categoria":          xywCategoria,
		"xyw":                    xyw})
	if err != nil {
		return false
	}
	return true
}

func DeleteProduct(id string) bool {
	clientMongo, err := mgo.Dial("localhost:27017/mobile")
	if err != err {
		fmt.Print(err)
	}
	defer clientMongo.Close()
	clientMongo.SetMode(mgo.Monotonic, true)
	productCollection := clientMongo.DB("patronus-mgo").C("products")
	err = productCollection.Remove(bson.M{"xyw": id})
	if err != nil {
		return false
	}
	return true
}
