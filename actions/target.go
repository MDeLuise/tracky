package actions

import (
	"encoding/json"
	"io"

	"tracky/log"
	"tracky/models"
	"tracky/response"
	"tracky/services"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

type TargetResource struct{}

func (t TargetResource) List(c buffalo.Context) error {
	target := &models.Targets{}
	if err := services.GetAllTargets(c.Value("tx").(*pop.Connection), target); err != nil {
		return response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, target)
}

func (t TargetResource) Show(c buffalo.Context) error {
	id := c.Param("target_id")
	target := &models.Target{}
	if err := services.GetTargetByID(c.Value("tx").(*pop.Connection), target, id); err != nil {
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
	if err = services.CreateTarget(c.Value("tx").(*pop.Connection), target); err != nil {
		return response.SendBadRequestError(c, err)
	}
	return response.SendOKResponse(c, target)
}

func (t TargetResource) Destroy(c buffalo.Context) error {
	id := c.Param("target_id")
	if err := services.DestroyTarget(c.Value("tx").(*pop.Connection), id); err != nil {
		return response.SendGeneralError(c, err)
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
	targetToUpdate.Unit = target.Unit
	if err = services.UpdateTarget(c.Value("tx").(*pop.Connection), id, targetToUpdate); err != nil {
		return response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, targetToUpdate)
}
