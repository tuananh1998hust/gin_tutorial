package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tuananh1998hust/gin_tutorial/models"
	"github.com/tuananh1998hust/gin_tutorial/utils"

	"go.mongodb.org/mongo-driver/bson"
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

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	name := c.PostForm("name")
	password := c.PostForm("password")
	password, err = utils.HashPassword(password)
	createdAt := time.Now()
	updatedAt := time.Now()

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
