package actions

import (
	"fmt"
	"net/http"
	"os"
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
	secret := os.Getenv("JWT_SECRET")

	if err := c.Bind(user); err != nil {
		return err
	}

	username := user.Username
	password := user.Password

	if username == "" || password == "" {
		log.SysLog.Error("username or password empty")
		return response.SendError(
			c,
			http.StatusBadRequest,
			fmt.Errorf("username and password cannot be empty"),
		)
	}

	q := models.DB.Select("id, username, password").Where("username = ?", username)
	err := q.First(user)
	if err != nil {
		return response.SendGeneralError(c, err)
	}

	PasswordHash := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(PasswordHash), []byte(password))
	if err != nil {
		log.SysLog.Error("invalid username or password")
		return response.SendError(
			c,
			http.StatusUnauthorized,
			fmt.Errorf("invalid username or password"),
		)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.SysLog.Error("error creating the jwt token")
		return response.SendGeneralError(c, err)
	}

	return response.SendOKResponse(c, map[string]string{"token": tokenString})
}
