package bundle

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"developer.patronus.io/mobile_backend/app/handler"
	"developer.patronus.io/mobile_backend/app/middleware"
	"developer.patronus.io/mobile_backend/app/model"
	"github.com/gorilla/mux"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	response := middleware.ListAllProducts()
	if len(response) > 0 {
		handler.RespondJSON(w, http.StatusOK, response)
	} else {
		handler.RespondError(w, http.StatusBadRequest, "Erro search products")
	}
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	bytesBody, _ := ioutil.ReadAll(r.Body)
	if len(bytesBody) != 0 {
		defer r.Body.Close()
		var product model.Product
		json.Unmarshal(bytesBody, &product)
		response := middleware.PostProduct(
			product.NomeProduto,
			product.ValorProduto,
			product.QuantidadeEstoque,
			product.Ativo,
			product.DataExpiracaoAnuncio,
			product.FotoProduto,
			product.CodigoBarras,
			product.Coordenadas,
			product.XywCategoria,
			product.Xyw)
		if response == true {
			handler.RespondJSON(w, http.StatusOK, response)
		} else {
			handler.RespondError(w, http.StatusBadRequest, "Invalid insert")
		}

	} else {
		handler.RespondError(w, http.StatusBadRequest, "err insert")
	}
}

func PutProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	bytesBody, _ := ioutil.ReadAll(r.Body)
	if len(bytesBody) != 0 {
		defer r.Body.Close()
		var product model.Product
		json.Unmarshal(bytesBody, &product)
		response := middleware.PutProduct(
			product.NomeProduto,
			product.ValorProduto,
			product.QuantidadeEstoque,
			product.Ativo,
			product.DataExpiracaoAnuncio,
			product.FotoProduto,
			product.CodigoBarras,
			product.Coordenadas,
			product.XywCategoria,
			product.Xyw,
			id)
		if response == true {
			handler.RespondJSON(w, http.StatusOK, response)
		} else {
			handler.RespondError(w, http.StatusBadRequest, "Invalid update")
		}

	} else {
		handler.RespondError(w, http.StatusBadRequest, "err update")
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := middleware.DeleteProduct(id)
	if response {
		handler.RespondJSON(w, http.StatusOK, response)
	} else {
		handler.RespondError(w, http.StatusBadRequest, "Erro delete product")
	}
}
