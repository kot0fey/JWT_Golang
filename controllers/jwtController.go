package controllers

import (
	"awesomeProject1/models"
	"awesomeProject1/repositories"
	"awesomeProject1/services"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	r.ParseForm()
	guid := r.Form.Get("guid")
	user, err := repositories.GetUserById(guid)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
	}
	sendCookiesTokens(w, user)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}
	refreshTokenCookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "Can not find token", http.StatusBadRequest)
	}
	user, err := services.ValidateRefreshToken(refreshTokenCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	sendCookiesTokens(w, user)
}

func sendCookiesTokens(w http.ResponseWriter, user *models.User) {
	access_token, refresh_token, err := services.CreateTokenPair(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    *access_token,
		HttpOnly: true,
		Secure:   true,
	}
	cookie2 := &http.Cookie{
		Name:     "refresh_token",
		Value:    *refresh_token,
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cookie)
	http.SetCookie(w, cookie2)
	w.WriteHeader(http.StatusOK)
}
