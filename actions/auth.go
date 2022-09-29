package actions

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"tracky_go/log"
	"tracky_go/models"
	"tracky_go/response"

	"github.com/gobuffalo/buffalo"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func AuthLogin(c buffalo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		log.SysLog.WithField("err", err).Error("error processing request")
		return response.SendBadRequestError(c, err)
	}

	username := user.Username
	password := user.Password
	if username == "" || password == "" {
		log.SysLog.Error("username or password empty")
		return response.SendBadRequestError(
			c,
			fmt.Errorf("username and password cannot be empty"),
		)
	}

	q := models.DB.Select("id, username, password").Where("username = ?", username)
	err := q.First(user)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error while searching in DB")
		return response.SendGeneralError(c, err)
	}

	PasswordHash := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(PasswordHash), []byte(password))
	if err != nil {
		log.SysLog.Error("invalid username or password")
		return response.SendUnauthorizedError(
			c,
			fmt.Errorf("invalid username or password"),
		)
	}

	token, err := createAccessToken(user.ID.String())
	if err != nil {
		log.SysLog.WithField("err", err).Error("error creating the jwt token")
		return response.SendGeneralError(c, err)
	}

	refreshToken, err := createRefreshToken(user.ID.String())
	if err != nil {
		log.SysLog.WithField("err", err).Error("error creating the refresh jwt token")
		return response.SendGeneralError(c, err)
	}

	return response.SendOKResponse(c, map[string]string{
		"token": token, "refreshToken": refreshToken})
}

func AuthRefresh(c buffalo.Context) error {
	usedRefreshToken := &struct{ RefreshToken string }{}
	refreshSecret := os.Getenv("JWT_REFRESH_SECRET")
	if err := c.Bind(usedRefreshToken); err != nil {
		log.SysLog.WithField("err", err).Error("error processing request")
		return response.SendBadRequestError(c, err)
	}

	token, err := jwt.Parse(usedRefreshToken.RefreshToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.SysLog.WithField("signingMethod", token.Header["alg"]).
					Error("unexpected signing method")
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(refreshSecret), nil
		})
	if err != nil {
		log.SysLog.WithField("err", err).Error("error parsing token")
		return response.SendBadRequestError(c, err)
	} else if !token.Valid {
		log.SysLog.Error("token not valid")
		return response.SendBadRequestError(c, fmt.Errorf("token not valid"))
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.SysLog.Error("error parsing token claims")
		return response.SendGeneralError(c, fmt.Errorf("error parsing token claims"))
	} else if !claims["refresh"].(bool) {
		log.SysLog.Error("token is not a refresh token")
		return response.SendBadRequestError(c, fmt.Errorf("token is not a refresh token"))
	}
	userID := claims["id"].(string)
	newToken, err := createAccessToken(userID)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error creating the jwt token")
		return response.SendGeneralError(c, err)
	}
	newRefreshToken, err := createRefreshToken(userID)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error creating the refresh token")
		return response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, map[string]string{
		"token": newToken, "refreshToken": newRefreshToken})
}

func createAccessToken(userID string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	expirationHour, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRATION_HOUR"))
	if err != nil {
		return "", err
	}
	expiration := time.Now().Add(time.Hour * time.Duration(expirationHour)).Unix()
	return createToken(userID, secret, false, expiration)
}

func createRefreshToken(userID string) (string, error) {
	secret := os.Getenv("JWT_REFRESH_SECRET")
	expirationHour, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRATION_HOUR"))
	if err != nil {
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
		return "", err
	}
	return tokenString, nil
}
