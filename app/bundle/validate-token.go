package bundle

import (
	"net/http"
)

func validationTokenUser(token, userEmail string) bool {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "http://35.198.48.166:/verify-token", nil)
	request.Header.Set("Patronus-Token", "Bearer "+token)
	request.Header.Set("Patronus-LoggedIn", userEmail)
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		return true
	}
	return false
}
