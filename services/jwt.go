package services

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"tracky/log"

	"github.com/golang-jwt/jwt/v4"
)

func CreateAccessToken(userID string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	expirationHour, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRATION_HOUR"))
	if err != nil {
		log.SysLog.WithField("err", err).Error("error getting the expiration date")
		return "", err
	}
	expiration := time.Now().Add(time.Hour * time.Duration(expirationHour)).Unix()
	return createToken(userID, secret, false, expiration)
}

func CreateRefreshToken(userID string) (string, error) {
	secret := os.Getenv("JWT_REFRESH_SECRET")
	expirationHour, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRATION_HOUR"))
	if err != nil {
		log.SysLog.WithField("err", err).Error("error getting the expiration date")
		return "", err
	}
	expiration := time.Now().Add(time.Hour * time.Duration(expirationHour)).Unix()
	return createToken(userID, secret, true, expiration)
}

func createToken(userID, secret string, refresh bool, expiration int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      userID,
		"refresh": refresh,
		"exp":     expiration,
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.SysLog.WithField("err", err).Error("error creating token")
		return "", err
	}
	return tokenString, nil
}

func IsTokenValid(token string) bool {
	_, err := parseToken(token)
	return err == nil
}

func GetTokenClaim(token, claim string) (interface{}, error) {
	parsedToken, err := parseToken(token)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error parsing token")
		return nil, err
	}
	return parsedToken.Claims.(jwt.MapClaims)[claim], nil
}

func GetTokenClaims(token string) (interface{}, error) {
	parsedToken, err := parseToken(token)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error parsing token")
		return nil, err
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}

func parseToken(token string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_REFRESH_SECRET")
	return jwt.Parse(token,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.SysLog.WithField("signingMethod", token.Header["alg"]).
					Error("unexpected signing method")
				return nil,
					fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})
}
