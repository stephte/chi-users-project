package testhelper

import (
	"chi-users-project/app/services/dtos"
	"encoding/json"
	"net/http"
)

// ------ User specific helpers -----

func(this TestHelper) GetUserOutDTO(res *http.Response) dtos.UserOutDTO {
	user := dtos.UserOutDTO{}
	jsonErr := json.Unmarshal(this.GetResponseBody(res), &user)
	if jsonErr != nil {
		this.t.Error(jsonErr)
	}

	return user
}
