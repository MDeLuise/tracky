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
	expirationSeconds, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRATION_SECONDS"))
	if err != nil {
		log.SysLog.WithField("err", err).Error("error getting the expiration date")
		return "", err
	}
	expiration := time.Now().Add(time.Second * time.Duration(expirationSeconds)).Unix()
	return createToken(userID, secret, false, expiration)
}

func CreateRefreshToken(userID string) (string, error) {
	secret := os.Getenv("JWT_REFRESH_SECRET")
	expirationSeconds, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRATION_SECONDS"))
	if err != nil {
		log.SysLog.WithField("err", err).Error("error getting the expiration date")
		return "", err
	}
	expiration := time.Now().Add(time.Second * time.Duration(expirationSeconds)).Unix()
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

func IsTokenValid(token string, refresh bool) bool {
	_, err := parseToken(token, refresh)
	return err == nil
}

func GetTokenClaim(token, claim string, refresh bool) (interface{}, error) {
	parsedToken, err := parseToken(token, refresh)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error parsing token")
		return nil, err
	}
	return parsedToken.Claims.(jwt.MapClaims)[claim], nil
}

func GetTokenClaims(token string, refresh bool) (interface{}, error) {
	parsedToken, err := parseToken(token, refresh)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error parsing token")
		return nil, err
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}

func GetExpiration(token string, refresh bool) (time.Time, error) {
	parsedToken, err := parseToken(token, refresh)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error parsing token")
		return time.Now(), err
	}
	return time.Unix(int64(parsedToken.Claims.(jwt.MapClaims)["exp"].(float64)), 0), nil
}

func parseToken(token string, refresh bool) (*jwt.Token, error) {
	var secret string
	if !refresh {
		secret = os.Getenv("JWT_SECRET")
	} else {
		secret = os.Getenv("JWT_REFRESH_SECRET")
	}
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
