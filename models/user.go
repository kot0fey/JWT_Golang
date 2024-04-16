package models

import "github.com/google/uuid"

type User struct {
	Id string `bson:"_id"`
}

func NewUser() User {
	user := User{}
	user.Id = uuid.New().String()
	return user
}
