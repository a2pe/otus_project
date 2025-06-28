package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"otus_project/internal/config"

	"time"
)

type Credentials struct {
	Login    string `json:"login" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=4,max=64"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

// LoginHandler godoc
// @Summary Авторизация
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body Credentials true "Login and password"
// @Success 200 {object} TokenResponse
// @Failure 401 {string} string "unauthorized"
// @Router /api/login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(creds); err != nil {
		http.Error(w, "validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	authConfig := config.LoadAuthConfig()

	if creds.Login != authConfig.AdminLogin || creds.Password != authConfig.AdminPassword {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	claims := jwt.MapClaims{
		"login": creds.Login,
		"exp":   time.Now().Add(time.Hour * 8).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, _ := token.SignedString([]byte(authConfig.JWTSecret))

	resp := TokenResponse{Token: signed}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
