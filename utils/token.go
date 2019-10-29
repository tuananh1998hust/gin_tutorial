package utils

import (
	"log"

	jwt "github.com/dgrijalva/jwt-go"
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

	return string(hash), err
}

// ComparePassword :
func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateToken :
func GenerateToken(userID primitive.ObjectID, email, name string) string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), Token{
		ID:    userID,
		Email: email,
		Name:  name,
	})

	secret := config.Key.SecretOrKey

	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}
