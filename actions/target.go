package actions

import (
	"encoding/json"
	"io"

	"tracky_go/log"
	"tracky_go/models"
	"tracky_go/response"

	"github.com/gobuffalo/buffalo"
)

type TargetResource struct{}

func (t TargetResource) List(c buffalo.Context) error {
	target := []models.Target{}
	err := models.DB.All(&target)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error while connecting to DB")
		return response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, target)
}

func (t TargetResource) Show(c buffalo.Context) error {
	id := c.Param("target_id")
	target := models.Target{}
	err := models.DB.Find(&target, id)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity with id")
		return response.SendNotFoundError(c, err)
	}
	return response.SendOKResponse(c, &target)
}

func (t TargetResource) Create(c buffalo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.SysLog.WithField("err", err).Error("Error reading body")
		return response.SendBadRequestError(c, err)
	}
	target := &models.Target{}
	json.Unmarshal([]byte(body), target)
	vErr, err := models.DB.ValidateAndCreate(target)
	if err != nil {
		log.SysLog.WithField("err", err).Error("entity not valid")
		return response.SendBadRequestError(c, err)
	}
	if vErr.HasAny() {
		log.SysLog.WithField("vErr", vErr.Errors).Error("entity not valid")
		return response.SendBadRequestError(c, vErr)
	}
	return response.SendOKResponse(c, target)
}

func (t TargetResource) Destroy(c buffalo.Context) error {
	id := c.Param("target_id")
	target := &models.Target{}
	if err := models.DB.Find(target, id); err != nil {
		return response.SendNotFoundError(c, err)
	}
	if err := models.DB.Destroy(target); err != nil {
		response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, nil)
}

func (t TargetResource) Update(c buffalo.Context) error {
	id := c.Param("target_id")
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot read body")
		return response.SendBadRequestError(c, err)
	}
	targetToUpdate := &models.Target{}
	err = models.DB.Find(targetToUpdate, id)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity with id")
		return response.SendNotFoundError(c, err)
	}
	target := &models.Target{}
	json.Unmarshal([]byte(body), target)
	targetToUpdate.Description = target.Description
	targetToUpdate.Name = target.Name
	vErr, err := models.DB.ValidateAndUpdate(targetToUpdate)
	if vErr.HasAny() {
		log.SysLog.WithField("err", err).Error("entity not valid")
		return response.SendBadRequestError(c, vErr)
	} else if err != nil {
		log.SysLog.WithField("err", err).Error("entity not valid")
		return response.SendBadRequestError(c, err)
	}
	return response.SendOKResponse(c, targetToUpdate)
}
