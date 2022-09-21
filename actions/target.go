package actions

import (
	"encoding/json"
	"io"
	"net/http"

	"tracky_go/log"
	"tracky_go/models"
	"tracky_go/response"

	"github.com/gobuffalo/buffalo"
)

func TargetIndex(c buffalo.Context) error {
	target := []models.Target{}
	err := models.DB.All(&target)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error while connecting to DB")
		return response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, target)
}

func TargetShow(c buffalo.Context) error {
	id := c.Param("id")
	target := models.Target{}
	err := models.DB.Find(&target, id)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity with id")
		return response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, &target)
}

func TargetAdd(c buffalo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.SysLog.WithField("err", err).Error("Error reading body")
		return response.SendGeneralError(c, err)
	}
	target := &models.Target{}
	json.Unmarshal([]byte(body), target)
	vErr, err := models.DB.ValidateAndCreate(target)
	if err != nil {
		log.SysLog.WithField("err", err).Error("entity not valid")
		return response.SendError(c, http.StatusBadRequest , err)
	}
	if vErr.HasAny() {
		log.SysLog.WithField("vErr", vErr.Errors).Error("entity not valid")
		return response.SendError(c, http.StatusBadRequest , vErr)
	}
	return response.SendOKResponse(c, target)
}

func TargetDelete(c buffalo.Context) error {
	id := c.Param("id")
	target := &models.Target{}
	if err := models.DB.Find(target, id); err != nil {
		return response.SendError(c, http.StatusNotFound, err)
	}
	if err := models.DB.Destroy(target); err != nil {
		response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, nil)
}

func TargetUpdate(c buffalo.Context) error {
	id := c.Param("id")
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot read body")
		return response.SendGeneralError(c, err)
	}
	targetToUpdate := &models.Target{}
	err = models.DB.Find(targetToUpdate, id)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity with id")
		return response.SendError(c, http.StatusNotFound, err)
	}
	target := &models.Target{}
	json.Unmarshal([]byte(body), target)
	targetToUpdate.Description = target.Description
	targetToUpdate.Name = target.Name
	vErr, err := models.DB.ValidateAndUpdate(targetToUpdate)
	if vErr.HasAny() {
		log.SysLog.WithField("err", err).Error("entity not valid")
		return response.SendError(c, http.StatusBadRequest, vErr)
	} else if err != nil {
		log.SysLog.WithField("err", err).Error("entity not valid")
		return response.SendError(c, http.StatusBadRequest, err)
	}
	return response.SendOKResponse(c, targetToUpdate)
}
