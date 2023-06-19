package handler

import (
	"auth/repository"
	"auth/response"
	"auth/service"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	//cfg      *config.Config
	userRepo     repository.IUserRepository
	tokenService *service.TokenService
}

func NewUserHandler(userRepo repository.IUserRepository, tokenService *service.TokenService) *UserHandler {
	return &UserHandler{
		userRepo:     userRepo,
		tokenService: tokenService,
	}
}

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//tokenService := service.NewTokenService(h.cfg)
		claims, err := h.tokenService.ValidateAccessToken(h.tokenService.GetTokenFromBearerString(r.Header.Get("Authorization")))
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		user, err := h.userRepo.GetUserByID(claims.ID)
		if err != nil {
			http.Error(w, "User does not exist", http.StatusBadRequest)
			return
		}

		resp := response.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}
