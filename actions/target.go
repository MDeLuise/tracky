package actions

import (
	"encoding/json"
	"io"

	"tracky_go/log"
	"tracky_go/models"
	"tracky_go/response"
	"tracky_go/services"

	"github.com/gobuffalo/buffalo"
)

type TargetResource struct{}

func (t TargetResource) List(c buffalo.Context) error {
	target := &models.Targets{}
	if err := services.GetAllTargets(target); err != nil {
		return response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, target)
}

func (t TargetResource) Show(c buffalo.Context) error {
	id := c.Param("target_id")
	target := &models.Target{}
	if err := services.GetTargetByID(target, id); err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity")
		return response.SendNotFoundError(c, err)
	}
	return response.SendOKResponse(c, target)
}

func (t TargetResource) Create(c buffalo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.SysLog.WithField("err", err).Error("Error reading body")
		return response.SendBadRequestError(c, err)
	}
	target := &models.Target{}
	json.Unmarshal([]byte(body), target)
	if err = services.CreateTarget(target); err != nil {
		response.SendBadRequestError(c, err)
	}
	return response.SendOKResponse(c, target)
}

func (t TargetResource) Destroy(c buffalo.Context) error {
	id := c.Param("target_id")
	if err := services.DestroyTarget(id); err != nil {
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
	target := &models.Target{}
	json.Unmarshal([]byte(body), target)
	targetToUpdate.Description = target.Description
	targetToUpdate.Name = target.Name
	if err = services.UpdateTarget(id, targetToUpdate); err != nil {
		return response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, targetToUpdate)
}
