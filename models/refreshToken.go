package models

type Token struct {
	UserId string `bson:"_id"`
	Token  string `bson:"refresh_token"`
	Expiry int64  `bson:"exp"`
}
