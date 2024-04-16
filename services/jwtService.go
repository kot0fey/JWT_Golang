package services

import (
	"awesomeProject1/models"
	"awesomeProject1/repositories"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func CreateToken(user *models.User, ttl time.Duration) (*string, error) {
	key := os.Getenv("TOKEN_SECRET")
	claims := createClaims(user, ttl)
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString([]byte(key))
	if err != nil {
		return nil, err
	}
	return &token, nil
}
func createClaims(user *models.User, ttl time.Duration) jwt.Claims {
	now := time.Now().UTC()
	claims := make(jwt.MapClaims)
	claims["sub"] = user.Id
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	return claims
}

func CreateRefreshToken(user *models.User) (*string, error) {
	refreshToken, err := CreateToken(user, time.Hour*24)
	if err != nil {
		return nil, err
	}
	responseRefreshToken := base64.URLEncoding.EncodeToString([]byte(*refreshToken))
	dbRefreshToken, err := bcrypt.GenerateFromPassword(hashToken(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	dbRefreshTokenString := string(dbRefreshToken)
	err = repositories.InsertRefreshToken(&dbRefreshTokenString, user)
	if err != nil {
		return nil, err
	}
	return &responseRefreshToken, nil
}

func hashToken(token *string) []byte {
	hasher := sha1.New()
	hasher.Write([]byte(*token))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return []byte(sha)
}

func ValidateRefreshToken(requestRefreshToken string) (*models.User, error) {
	refreshToken, _ := base64.URLEncoding.DecodeString(requestRefreshToken)
	refreshTokenString := string(refreshToken)
	jwtToken, _ := jwt.ParseWithClaims(refreshTokenString, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	claims, ok := jwtToken.Claims.(*models.TokenClaims)
	if !(ok && jwtToken.Valid) {
		return nil, errors.New("Invalid token")
	}
	userId := claims.Sub
	dbToken, err := repositories.GetRefreshTokenById(&userId)
	if dbToken == nil {
		return nil, err
	}
	hashedRefreshToken := hashToken(&refreshTokenString)
	err = bcrypt.CompareHashAndPassword([]byte(dbToken.Token), hashedRefreshToken)
	if err != nil {
		return nil, err
	}
	if dbToken.Expiry <= time.Now().Unix() {
		return nil, errors.New("Token expired")
	}
	repositories.DeleteRefreshTokenById(&userId)
	user, err := repositories.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateTokenPair(user *models.User) (*string, *string, error) {
	access_token, err := CreateToken(user, time.Minute*15)
	if err != nil {
		return nil, nil, err
	}
	refresh_token, err := CreateRefreshToken(user)
	if err != nil {
		return nil, nil, err
	}
	return access_token, refresh_token, nil

}
