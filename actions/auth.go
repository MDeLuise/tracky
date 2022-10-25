package actions

import (
	"fmt"
	"tracky/log"
	"tracky/models"
	"tracky/response"
	"tracky/services"

	"github.com/gobuffalo/buffalo"
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
		return response.SendUnauthorizedError(c, fmt.Errorf("invalid username or password"))
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

	token, err := services.CreateAccessToken(user.ID.String())
	if err != nil {
		log.SysLog.WithField("err", err).Error("error creating the jwt token")
		return response.SendGeneralError(c, err)
	}

	refreshToken, err := services.CreateRefreshToken(user.ID.String())
	if err != nil {
		log.SysLog.WithField("err", err).Error("error creating the refresh jwt token")
		return response.SendGeneralError(c, err)
	}

	accessExp, err := services.GetExpiration(token, false)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error retrieving the expiration of the access token")
		return response.SendGeneralError(c, err)
	}
	refreshExp, err := services.GetExpiration(refreshToken, true)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error retrieving the expiration of the refresh token")
		return response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, map[string]string{
		"token":                    token,
		"refresh_token":            refreshToken,
		"token_expires_on":         accessExp.String(),
		"refresh_token_expires_on": refreshExp.String(),
	})
}

func AuthRefresh(c buffalo.Context) error {
	usedRefreshToken := &struct{ Refresh_token string }{}
	if err := c.Bind(usedRefreshToken); err != nil {
		log.SysLog.WithField("err", err).Error("error processing request")
		return response.SendBadRequestError(c, err)
	}

	if !services.IsTokenValid(usedRefreshToken.Refresh_token, true) {
		log.SysLog.Error("token not valid")
		return response.SendBadRequestError(c, fmt.Errorf("token not valid"))
	}

	if isRefresh, _ := services.GetTokenClaim(usedRefreshToken.Refresh_token, "refresh", true); !isRefresh.(bool) {
		log.SysLog.Error("token is not a refresh token")
		return response.SendBadRequestError(c, fmt.Errorf("token is not a refresh token"))
	}
	userID, _ := services.GetTokenClaim(usedRefreshToken.Refresh_token, "id", true)
	newToken, err := services.CreateAccessToken(userID.(string))
	if err != nil {
		log.SysLog.WithField("err", err).Error("error creating the jwt token")
		return response.SendGeneralError(c, err)
	}
	newRefreshToken, err := services.CreateRefreshToken(userID.(string))
	if err != nil {
		log.SysLog.WithField("err", err).Error("error creating the refresh token")
		return response.SendGeneralError(c, err)
	}
	accessExp, err := services.GetExpiration(newToken, false)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error retrieving the expiration of the access token")
		return response.SendGeneralError(c, err)
	}
	refreshExp, err := services.GetExpiration(newRefreshToken, true)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error retrieving the expiration of the refresh token")
		return response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, map[string]string{
		"token":              newToken,
		"refresh_token":      newRefreshToken,
		"token_expires_on":   accessExp.String(),
		"refresh_expires_on": refreshExp.String(),
	})
}
