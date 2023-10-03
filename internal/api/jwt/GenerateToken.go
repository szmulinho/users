package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/szmulinho/users/internal/model"
	"net/http"
	"time"
)

func GenerateToken(w http.ResponseWriter, r *http.Request, userID int64, isDoctor bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   userID,
		"isDoctor": isDoctor,
		"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})

	tokenString, err := token.SignedString(model.JwtKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	}

	return tokenString, nil
}
