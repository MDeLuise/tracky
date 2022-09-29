package actions

import (
	"encoding/json"
	"io"

	"tracky/log"
	"tracky/models"
	"tracky/response"

	"github.com/gobuffalo/buffalo"
	"github.com/gofrs/uuid"
)

func UsersCreate(c buffalo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot read body")
		return err
	}
	user := &models.User{}
	json.Unmarshal([]byte(body), user)
	user.ID = uuid.Must(uuid.NewV4())
	vErr, err := models.DB.ValidateAndCreate(user)
	if err != nil {
		log.SysLog.Error("entity not valid")
		response.SendBadRequestError(c, err)
	}
	if vErr.HasAny() {
		log.SysLog.WithField("err", err).Error("entity not valid")
		return response.SendBadRequestError(c, vErr)
	}

	user.Password = ""
	return response.SendOKResponse(c, user)
}

func UsersRead(c buffalo.Context) error {
	users := []models.User{}
	err := models.DB.Select("id, username, created_at, updated_at").All(&users)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error while connecting to DB")
		response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, users)
}

func UsersReadByID(c buffalo.Context) error {
	id := c.Param("id")
	user := models.User{}
	err := models.DB.Select("id, username, created_at, updated_at").Find(&user, id)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity")
		response.SendNotFoundError(c, err)
	}
	return response.SendOKResponse(c, &user)
}
