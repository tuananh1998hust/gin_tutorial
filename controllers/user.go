package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tuananh1998hust/gin_tutorial/config"
	"github.com/tuananh1998hust/gin_tutorial/models"
	"github.com/tuananh1998hust/gin_tutorial/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateUser :
func CreateUser(c *gin.Context) {
	email := c.PostForm("email")

	var findUserByEmail models.User
	err := models.UserCollection.FindOne(context.TODO(), bson.D{bson.E{Key: "email", Value: email}}).Decode(&findUserByEmail)

	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is exist, Please change email",
		})
		return
	}

	if err == mongo.ErrNoDocuments {
		name := c.PostForm("name")
		password := c.PostForm("password")
		password, err = utils.HashPassword(password)
		createdAt := time.Now()
		updatedAt := time.Now()

		if email == "" || name == "" || password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Please fill all the fields",
			})
			return
		}

		if len(password) < 6 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Password is must unless 6 characters",
			})
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		id, err := models.UserCollection.InsertOne(context.TODO(), bson.M{
			"name":      name,
			"email":     email,
			"password":  password,
			"createdAt": createdAt,
			"updatedAt": updatedAt,
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"id":        id.InsertedID,
			"name":      name,
			"email":     email,
			"password":  password,
			"createdAt": createdAt,
			"updatedAt": updatedAt,
		})
	}
}

// GetUser :
func GetUser(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token is not valid",
		})
		return
	}

	token := splitToken[1]

	// at(time.Unix(0, 0), func() {
	// 	decode, _ := jwt.ParseWithClaims(token, &utils.Token{}, func(decode *jwt.Token) (interface{}, error) {
	// 		secret := config.Key.SecretOrKey
	// 		return []byte(secret), nil
	// 	})

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"data": decode,
	// 	})
	// })

	decode, _ := jwt.ParseWithClaims(token, &utils.Token{}, func(decode *jwt.Token) (interface{}, error) {
		secret := config.Key.SecretOrKey
		return []byte(secret), nil
	})

	c.JSON(http.StatusOK, gin.H{
		"data": decode,
	})

	// if claims, ok := decode.Claims.(*utils.Token); ok && decode.Valid {
	// 	log.Print(claims)
	// 	var findUser models.User
	// 	err := models.UserCollection.FindOne(context.TODO(), bson.D{bson.E{Key: "_id", Value: claims.ID}}).Decode(&findUser)

	// 	if err == mongo.ErrNoDocuments {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "User not found",
	// 		})

	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"_id":   findUser.ID,
	// 		"name":  findUser.Name,
	// 		"email": findUser.Email,
	// 	})
	// } else {
	// 	log.Print("HEllo")
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "UnAuthorization",
	// 	})
	// }
}

func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}
