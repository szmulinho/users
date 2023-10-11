package endpoints

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/szmulinho/users/internal/model"
	"net/http"
	"time"
)

func (h *handlers) GenerateToken(w http.ResponseWriter, r *http.Request, ID int64, isUser bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": ID,
		"isUser": isUser,
		"exp":    time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})

	tokenString, err := token.SignedString(model.JwtKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	} else {
		fmt.Println("token generated")
	}

	return tokenString, nil
}
