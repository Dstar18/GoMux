package handler

import (
	"GoMux/config"
	"GoMux/usecase"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
		config.Logger.Error(err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Gets user Error!",
		})
		return
	}

	// return success
	config.Logger.Info("Gets user successfully")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Gets user successfully",
		"data":    users,
	})
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	// Get id from URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// check format ID
	id, err := strconv.Atoi(idStr)
	if err != nil {
		config.Logger.Error("Invalid ID format!")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Invalid ID format!",
		})
		return
	}

	// get data on DB
	user, err := h.userUseCase.GetUserById(uint(id))
	if err != nil {
		config.Logger.Error("Data not found!")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Data not found!",
		})
		return
	}

	// return success
	config.Logger.Info("Get user successfully")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Get user successfully",
		"data":    user,
	})
}
