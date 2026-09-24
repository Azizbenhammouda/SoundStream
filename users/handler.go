package users

import (
	"encoding/json"
	"net/http"
)

type userHandler struct {
	service UserService
}

func NewUserHandler(s UserService) userHandler {
	return userHandler{
		service: s,
	}
}
func (h userHandler) Register(w http.ResponseWriter, req *http.Request) {
	var input RegisterInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
