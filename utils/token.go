package utils

import (
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	config "github.com/tuananh1998hust/gin_tutorial/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Token :
type Token struct {
	ID    primitive.ObjectID `json:"id"`
	Email string             `json:"email"`
	Name  string             `json:"name"`
	jwt.StandardClaims
}

// HashPassword :
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		log.Print(err)
	}

	return string(hash), nil
}

// ComparePassword :
func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateToken :
func GenerateToken(userID primitive.ObjectID, email, name string) (string, error) {
	claims := Token{
		userID,
		"email",
		"name",
		jwt.StandardClaims{
			Issuer:    "test",
			ExpiresAt: time.Now().Unix() + 60*60*24*30*1000,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := config.Key.SecretOrKey

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Print(err)
	}

	return tokenString, nil
}

// DecodeToken :
func DecodeToken(c *gin.Context) (decode *jwt.Token) {
	bearerToken := c.GetHeader("Authorization")

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		return decode
	}

	token := splitToken[1]

	decode, err := jwt.ParseWithClaims(token, &Token{}, func(decode *jwt.Token) (interface{}, error) {
		secret := config.Key.SecretOrKey
		return []byte(secret), nil
	})

	if err != nil {
		log.Println(err)
	}

	return decode
}
