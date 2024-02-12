package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	_ "image/png"

	"github.com/dgrijalva/jwt-go"
	"github.com/gosimple/slug"
	"github.com/nfnt/resize"
	"github.com/sayedatif/tigerhall/config"
	"github.com/sayedatif/tigerhall/db"
	"gorm.io/gorm"
)

func IsUsernameExists(database *gorm.DB, username string) bool {
	var user db.User
	if err := database.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
	}
	if user.ID > 0 {
		return true
	}
	return false
}

func GenerateUsername(firstName string, lastName string) string {
	database := db.GetDB()
	originalUsername := strings.ToLower(firstName + "_" + lastName)

	initialSlug := slug.Make(originalUsername)

	newUsername := initialSlug
	counter := 1
	for IsUsernameExists(database, newUsername) {
		newSlug := fmt.Sprintf("%s_%d", initialSlug, counter)
		newUsername = newSlug
		counter++
	}

	return newUsername
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

func GetParsedTime(dateString string) (time.Time, error) {
	layout := "2006-01-02 15:04:05.999"

	parsedLastSeenAt, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Now(), err
	}
	return parsedLastSeenAt, nil
}

func HandleImageUpload(file multipart.File, userId int, tigerId string) (string, error) {
	_, _, err := image.DecodeConfig(file)
	if err != nil {
		return "", err
	}

	file.Seek(0, 0)

	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}
	resizedImg := resize.Resize(250, 200, img, resize.Lanczos3)

	stringUserId := strconv.Itoa(userId)
	userDir := filepath.Join("./images", stringUserId)
	err = os.MkdirAll(userDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	now := time.Now()
	formattedTime := now.Format("2006-01-02T15:04:05")
	filename := filepath.Base(fmt.Sprintf("tiger_%s_%s.jpg", tigerId, formattedTime))
	filepath := filepath.Join(fmt.Sprintf("./images/%s", stringUserId), filename)
	out, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	err = jpeg.Encode(out, resizedImg, nil)
	if err != nil {
		return "", err
	}

	return filepath, nil
}

func StringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}
