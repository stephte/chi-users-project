package testhelper

import (
	"chi-users-project/app/services/dtos"
	"encoding/json"
	"net/http"
)

// ------ User specific helpers -----

func(this TestHelper) GetUserDTO(res *http.Response) dtos.UserDTO {
	user := dtos.UserDTO{}
	jsonErr := json.Unmarshal(this.GetResponseBody(res), &user)
	if jsonErr != nil {
		this.t.Error(jsonErr)
	}

	return user
}
