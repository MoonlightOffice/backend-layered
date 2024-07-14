package internal

import (
	"net/http"

	"giants/pkg/usecase/user"
)

func userRegister(w http.ResponseWriter, r *http.Request) {
	reqData := struct {
		Email string `json:"email"`
	}{}
	if !bindReqData(r, &reqData) {
		writeResponse(w, 400, nil)
		return
	}

	uObj, err := user.UserService.Register(reqData.Email)
	if err != nil {
		writeResponse(w, 400, H{"log": err.Error()})
		return
	}

	writeResponse(w, 200, H{"userId": uObj.UserId})
}

func userGet(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, 200, nil)
}
