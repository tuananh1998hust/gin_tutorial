package controllers

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tuananh1998hust/gin_tutorial/models"
	"github.com/tuananh1998hust/gin_tutorial/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

		if len(password) < 6 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Password is must unless 6 characters",
			})

			return
		}

		password, _ = utils.HashPassword(password)

		if email == "" || name == "" || password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Please fill all the fields",
			})
			return
		}

		var newUser models.User

		id, err := newUser.CreateUser(name, email, password)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"id":       id.InsertedID,
			"name":     name,
			"email":    email,
			"password": password,
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

	decode := utils.DecodeToken(c)

	if claims, ok := decode.Claims.(*utils.Token); ok && decode.Valid {
		var findUser models.User
		findUser, err := findUser.FindUserByID(claims.ID)

		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "User not found",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"_id":   findUser.ID,
			"name":  findUser.Name,
			"email": findUser.Email,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token is not valid",
		})
	}
}

// UpdateUser :
func UpdateUser(c *gin.Context) {
	name := c.PostForm("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please fill all the fields",
		})

		return
	}

	bearerToken := c.GetHeader("Authorization")

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token is not valid",
		})
		return
	}

	decode := utils.DecodeToken(c)

	if claims, ok := decode.Claims.(*utils.Token); ok && decode.Valid {
		var findUser models.User
		findUser, err := findUser.FindUserByID(claims.ID)

		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "User not found",
			})

			return
		}

		paramID := c.Param("id")
		id, _ := primitive.ObjectIDFromHex(paramID)

		var updateUser models.User

		updateUser, err = updateUser.UpdateUser(id, name)

		log.Println(updateUser)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":    updateUser.ID,
			"name":  name,
			"email": updateUser.Email,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token is not valid",
		})
	}
}
