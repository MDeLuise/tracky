package actions

import (
	"strconv"
	"tracky/log"
	"tracky/models"
	"tracky/response"
	"tracky/services"

	"github.com/gobuffalo/buffalo"
)

func TargetMean(c buffalo.Context) error {
	targetID := c.Param("target_id")
	target := &models.Target{}
	if err := services.GetTargetByID(target, targetID); err != nil {
		return response.SendNotFoundError(c, err)
	}
	return response.SendOKResponse(c, services.CalcMean(target))
}

func TargetMeanAt(c buffalo.Context) error {
	targetID := c.Param("target_id")
	numOflastValues, err := strconv.Atoi(c.Param("at"))
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot parse value")
		return response.SendBadRequestError(c, err)
	}
	target := &models.Target{}
	if err := services.GetTargetByID(target, targetID); err != nil {
		return response.SendNotFoundError(c, err)
	}
	return response.SendOKResponse(c, services.CalcMeanAt(target, numOflastValues))
}

func TargetLastIncrement(c buffalo.Context) error {
	targetID := c.Param("target_id")
	target := &models.Target{}
	if err := services.GetTargetByID(target, targetID); err != nil {
		return response.SendNotFoundError(c, err)
	}
	return response.SendOKResponse(c, services.CalcLastIncr(target))
}
