package users

import (
	"encoding/json"
	"errors"
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

	user, err := h.service.Register(input)
	if err != nil {
		if errors.Is(err, ErrEmailTaken) {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	//might change this doesnt handle error well
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
