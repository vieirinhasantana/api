package bundle

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"developer.patronus.io/mobile_backend/app/handler"
	"developer.patronus.io/mobile_backend/app/middleware"
	"developer.patronus.io/mobile_backend/app/model"
)

func PostLoginUser(w http.ResponseWriter, r *http.Request) {
	bytesBody, _ := ioutil.ReadAll(r.Body)
	if len(bytesBody) != 0 {
		defer r.Body.Close()
		var user model.User
		json.Unmarshal(bytesBody, &user)
		response := middleware.SearchDataMongo(user.Username, user.Password)
		if response == true {
			handler.RespondJSON(w, http.StatusOK, response)
		} else {
			handler.RespondError(w, http.StatusBadRequest, "Credential is invalid for authentication.")
		}

	} else {
		handler.RespondError(w, http.StatusBadRequest, "Credential is invalid for authentication.")
	}
}
