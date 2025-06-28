package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
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

	if creds.Login != os.Getenv("ADMIN_LOGIN") || creds.Password != os.Getenv("ADMIN_PASSWORD") {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	claims := jwt.MapClaims{
		"login": creds.Login,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	resp := TokenResponse{Token: signed}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
