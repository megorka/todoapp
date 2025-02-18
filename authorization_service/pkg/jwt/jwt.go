package jwt

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("a99aca380883ecc5438bc3c6afa91a3bbdc9b921cd350677e1d23971eabf35661d3d1752890dab9644ef6bed173600b235d97c82b76a68a3162fdf147995cf6f6d212125da8bc1c0ad3340ea212b5d7dddd5d8a55c4a0b125bf83e1a53aa5dd4dcad7a553f80bfaecfc7085225aaf5a831e69e531f072334d58106629345a501")
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateToken(userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"user_id": strconv.Itoa(userID),
		"exp":     expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (int, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, "", err
	}

	userID, ok := claims["user_id"].(int)
	if !ok {
		return 0, "", err
	}

	userRole, ok := claims["userRole"].(string)
	if !ok {
		return 0, "", err
	}

	return userID, userRole, nil
}