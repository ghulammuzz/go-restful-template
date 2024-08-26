package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ghulammuzz/go-restful-template/internal/errors"
	"github.com/ghulammuzz/go-restful-template/internal/service"
	"github.com/ghulammuzz/go-restful-template/pkg/logger"
	"github.com/ghulammuzz/go-restful-template/pkg/utils"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode request body", "error", err)
		utils.SendError(w, errors.NewAppError(errors.ErrPayload, "Invalid request payload"))
		return
	}

	if req.Username == "" {
		utils.SendError(w, errors.NewAppError(errors.ErrUsernameRequired, "Username is required"))
		return
	}
	if req.Password == "" {
		utils.SendError(w, errors.NewAppError(errors.ErrPasswordRequired, "Password is required"))
		return
	}

	err := h.service.RegisterUser(req.Username, req.Password)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			logger.Error("Failed to register user", "error_code", appErr.Code, "error_message", appErr.Message)
			utils.SendError(w, appErr)
			return
		}
		logger.Error("Failed to register user", "error", err)
		utils.SendError(w, errors.NewAppError(errors.ErrInternal, "Internal server error"))
		return
	}

	utils.SendResponse(w, http.StatusCreated, map[string]string{"message": "User registered successfully"})
}
