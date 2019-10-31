package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuananh1998hust/gin_tutorial/models"
	"github.com/tuananh1998hust/gin_tutorial/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Login :
func Login(c *gin.Context) {
	email := c.PostForm("email")

	var findUserByEmail models.User
	err := models.UserCollection.FindOne(context.TODO(), bson.D{bson.E{Key: "email", Value: email}}).Decode(&findUserByEmail)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User is not exist",
		})
	}

	if err == nil {
		password := c.PostForm("password")

		compare := utils.ComparePassword(password, findUserByEmail.Password)

		if compare != true {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Wrong password",
			})
			return
		}

		token, _ := utils.GenerateToken(findUserByEmail.ID, findUserByEmail.Email, findUserByEmail.Name)

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
