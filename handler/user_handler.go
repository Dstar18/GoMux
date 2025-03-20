package handler

import (
	"GoMux/config"
	"GoMux/usecase"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUseCase.GetUsers()
	if err != nil {
		config.Logger.Error(http.StatusInternalServerError, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Gets user Error!",
		})
		return
	}
	config.Logger.Info("Gets user successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Gets user successfully",
		"data":    users,
	})
}
