package actions

import (
	"encoding/json"
	"io"
	"net/http"
	"tracky_go/log"
	"tracky_go/models"
	"tracky_go/response"

	"github.com/gobuffalo/buffalo"
	"github.com/gofrs/uuid"
)

func ObservationIndex(c buffalo.Context) error {
	observation := []models.Observation{}
	err := models.DB.All(&observation)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error while connecting to DB")
		return response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, observation)
}

func ObservationShow(c buffalo.Context) error {
	id := c.Param("id")
	observation := models.Observation{}
	err := models.DB.Find(&observation, id)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity")
		return response.SendError(c, http.StatusNotFound, err)
	}
	return response.SendOKResponse(c, &observation)
}

func ObservationCreate(c buffalo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot read body")
		return response.SendGeneralError(c, err)
	}
	observation := &models.Observation{}
	err = json.Unmarshal([]byte(body), observation)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot unmarshal entity")
		return response.SendGeneralError(c, err)
	}
	vErr, err := models.DB.Eager().ValidateAndCreate(observation)
	if err != nil {
		log.SysLog.WithField("err", err).Error("entity not valid")
		return response.SendError(c, http.StatusBadRequest , err)
	}
	if vErr.HasAny() {
		log.SysLog.WithField("vErr", vErr).Error("entity not valid")
		return response.SendError(c, http.StatusBadRequest , vErr)
	}
	return response.SendOKResponse(c, observation)
}

func ObservationDelete(c buffalo.Context) error {
	id := c.Param("id")
	observation := &models.Observation{}
	if err := models.DB.Find(observation, id); err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity with id")
		return response.SendError(c, http.StatusNotFound, err)
	}
	if err := models.DB.Destroy(observation); err != nil {
		log.SysLog.WithField("err", err).Error("error while destroying entity")
		return response.SendGeneralError(c, err)
	}
	return c.Render(http.StatusOK, r.JSON("ok"))
}

func ObservationUpdate(c buffalo.Context) error {
	id := c.Param("id")
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot read body")
		return response.SendGeneralError(c, err)
	}
	observationToUpdate := &models.Observation{}
	if err := models.DB.Find(observationToUpdate, id); err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity")
		return response.SendError(c, http.StatusNotFound, err)
	}
	observation := &models.Observation{}
	if err = json.Unmarshal([]byte(body), observation); err != nil {
		log.SysLog.WithField("err", err).Error("cannot unmarshal json")
		return response.SendError(c, http.StatusNotFound, err)
	}
	observationToUpdate.Value = observation.Value
	if observation.TargetID != uuid.Nil {
		observationToUpdate.TargetID = observation.TargetID
	}
	vErr, err := models.DB.ValidateAndUpdate(observationToUpdate)
	if vErr.HasAny() {
		log.SysLog.WithField("err", err).Error("entity not valid")
		return response.SendError(c, http.StatusBadRequest, vErr)
	} else if err != nil {
		log.SysLog.WithField("vErr", vErr).Error("entity not valid")
		return response.SendError(c, http.StatusBadRequest, err)
	}
	return response.SendOKResponse(c, observationToUpdate)
}
