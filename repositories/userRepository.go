package repositories

import (
	"awesomeProject1/initializers"
	"awesomeProject1/models"
	"context"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func CreateUser() (*models.User, error) {
	user := models.NewUser()
	collection := initializers.CLIENT.Database(initializers.DBNAME).Collection("users")
	res, err := collection.InsertOne(context.TODO(), bson.M{"_id": user.Id})
	if err != nil {
		return nil, err
	}
	fmt.Println("Created user with id: ", res.InsertedID)
	return &user, nil
}

func GetUserById(id string) (*models.User, error) {
	collection := initializers.CLIENT.Database(initializers.DBNAME).Collection("users")
	var user *models.User
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
