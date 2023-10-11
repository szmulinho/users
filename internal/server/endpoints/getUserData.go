package endpoints

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/szmulinho/users/internal/model"
	"log"
	"net/http"
	"strings"
)

func (h *handlers) GetUserDataHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	user, err := h.getUserFromToken(token)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(response)
}

func (h *handlers) getUserFromToken(tokenString string) (*model.User, error) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return model.JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Invalid token claims")
	}

	userID := int64(claims["userID"].(float64))

	var user model.User
	if err := h.db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
