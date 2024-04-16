package repositories

import (
	"awesomeProject1/initializers"
	"awesomeProject1/models"
	"context"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func InsertRefreshToken(encryptedToken *string, user *models.User) error {
	collection := initializers.CLIENT.Database(initializers.DBNAME).Collection("refresh_tokens")
	//encryptedToken, _ := bcrypt.GenerateFromPassword([]byte(*token), bcrypt.DefaultCost)
	//encryptedToken = []byte(base64.StdEncoding.EncodeToString(encryptedToken))
	_, err := collection.InsertOne(context.TODO(), bson.M{"_id": user.Id, "refresh_token": encryptedToken, "exp": time.Now().Add(time.Hour * 24).Unix()})
	if err != nil {
		return err
	}
	fmt.Println("Refresh token created: ", encryptedToken)
	return nil
}

func GetRefreshTokenByToken(token *string) (*models.Token, error) {
	collection := initializers.CLIENT.Database(initializers.DBNAME).Collection("refresh_tokens")
	var refreshToken *models.Token
	err := collection.FindOne(context.TODO(), bson.M{"refresh_token": token}).Decode(&token)
	if err != nil {
		return nil, err
	}
	return refreshToken, nil
}

func GetRefreshTokenById(id *string) (*models.Token, error) {
	collection := initializers.CLIENT.Database(initializers.DBNAME).Collection("refresh_tokens")
	var refreshToken *models.Token
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&refreshToken)
	if err != nil {
		return nil, err
	}
	return refreshToken, nil
}

func DeleteRefreshTokenById(id *string) error {
	collection := initializers.CLIENT.Database(initializers.DBNAME).Collection("refresh_tokens")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
