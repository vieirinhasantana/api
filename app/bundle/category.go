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

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	response := middleware.ListAllCategories()
	if len(response) > 0 {
		handler.RespondJSON(w, http.StatusOK, response)
	} else {
		handler.RespondError(w, http.StatusBadRequest, "Erro search categories")
	}
}

func PostCategory(w http.ResponseWriter, r *http.Request) {
	bytesBody, _ := ioutil.ReadAll(r.Body)
	if len(bytesBody) != 0 {
		defer r.Body.Close()
		var category model.Category
		json.Unmarshal(bytesBody, &category)
		response := middleware.PostCategory(category.Descricao, category.Xyw)
		if response == true {
			handler.RespondJSON(w, http.StatusOK, response)
		} else {
			handler.RespondError(w, http.StatusBadRequest, "Invalid insert")
		}

	} else {
		handler.RespondError(w, http.StatusBadRequest, "err insert")
	}
}

func PutCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	bytesBody, _ := ioutil.ReadAll(r.Body)
	if len(bytesBody) != 0 {
		defer r.Body.Close()
		var category model.Category
		json.Unmarshal(bytesBody, &category)
		response := middleware.PutCategory(category.Descricao, category.Xyw, id)
		if response == true {
			handler.RespondJSON(w, http.StatusOK, response)
		} else {
			handler.RespondError(w, http.StatusBadRequest, "Invalid update")
		}

	} else {
		handler.RespondError(w, http.StatusBadRequest, "err update")
	}
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := middleware.DeleteCategory(id)
	if response {
		handler.RespondJSON(w, http.StatusOK, response)
	} else {
		handler.RespondError(w, http.StatusBadRequest, "Erro delete category")
	}
}
