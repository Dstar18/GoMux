package handler

import (
	"GoMux/config"
	"GoMux/entity"
	"GoMux/usecase"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
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
	return
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
	return
}

type UserValidate struct {
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	Role     string `json:"role" validate:"required,min=2,max=20"`
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// initize validate
	var user UserValidate

	// check request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		config.Logger.Error("Invalid request body")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid request body",
		})
		return
	}

	// validation
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "This field is" + " " + err.Tag() + " " + err.Param()
		}
		config.Logger.Error(errors)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": errors,
		})
		return
	}

	// validate password
	if err := config.ValidatePassword(user.Password); err != nil {
		config.Logger.Error(err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	// check email is ready
	_, err := h.userUseCase.GetUserByField("email", user.Email)
	if err == nil {
		config.Logger.Warn("email " + user.Email + " is already")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "email " + user.Email + " is already",
		})
		return
	}

	// hash password
	hashedPassword, err := config.HashPassword(user.Password)
	if err != nil {
		config.Logger.Error("Failed to hash password")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Failed to hash password",
		})
		return
	}

	param := entity.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  hashedPassword,
		Role:      user.Role,
		CreatedAt: time.Now(),
	}
	// create to db
	if err := h.userUseCase.CreateUser(&param); err != nil {
		config.Logger.Error(err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	// return success
	config.Logger.Info("Create user successfully")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Create user successfully",
		"data":    nil,
	})
	return
}
