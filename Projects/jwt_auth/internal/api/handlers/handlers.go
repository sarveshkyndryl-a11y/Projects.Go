package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
var jwtkey = []byte("my_secret_key")
//why jwtkey is stored as byte slice?
// Because it is more efficient to work with byte slices when dealing with binary data, such as JWTs.
// Additionally, using a byte slice avoids unnecessary string conversions and allocations.

var user = map[string]string{
	"admin": "admin@123",
	"admin1": "admin1@123",
}
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	expectedpassword,ok := user[credentials.Username]
	if !ok || expectedpassword != credentials.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: credentials.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
	})
	w.Write([]byte(tokenString))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Handle home
}

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	// Handle token refresh
}
