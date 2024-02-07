package utils

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sayedatif/tigerhall/config"
)

func generateRandomSlug(length int) string {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(randomBytes)[:length]
}

func GenerateUsername(firstName string, lastName string) string {
	randomSlug := generateRandomSlug(6)
	return strings.ToLower(firstName + "_" + lastName + "_" + randomSlug)
}

func GenerateToken(user_id int64, secret []byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetEnv(key string) any {
	config := config.GetConfig()
	env := config.Get(key)
	return env
}
