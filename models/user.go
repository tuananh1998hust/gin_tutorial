package models

import (
	"context"
	"log"
	"time"

	"github.com/tuananh1998hust/gin_tutorial/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserCollection :
var UserCollection *mongo.Collection = config.MongoClient.Database("GinTutorial").Collection("user")

// User :
type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updatedAt"`
}

// CreateUser :
func (u *User) CreateUser(name, email, password string) (id *mongo.InsertOneResult, err error) {
	createdAt, updatedAt := time.Now(), time.Now()

	// Save To DB
	newID, err := UserCollection.InsertOne(context.TODO(), bson.M{
		"name":      name,
		"email":     email,
		"password":  password,
		"createdAt": createdAt,
		"updatedAt": updatedAt,
	})

	return newID, err
}

// FindUserByID :
func (u *User) FindUserByID(id primitive.ObjectID) (user User, err error) {
	var findUser User

	err = UserCollection.FindOne(context.TODO(), bson.D{bson.E{Key: "_id", Value: id}}).Decode(&findUser)

	return findUser, err
}

// UpdateUser :
func (u *User) UpdateUser(id primitive.ObjectID, name string) (user User, err error) {
	var updateUser User

	err = UserCollection.FindOneAndUpdate(
		context.TODO(),
		bson.D{bson.E{Key: "_id", Value: id}},
		bson.M{"$set": bson.D{bson.E{Key: "name", Value: name}, bson.E{Key: "updated_at", Value: time.Now()}}},
	).Decode(&updateUser)

	log.Println(updateUser)

	return updateUser, err
}
