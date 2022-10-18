package actions

import (
	"encoding/json"
	"io"
	"net/http"
	"tracky/log"
	"tracky/models"
	"tracky/response"
	"tracky/services"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

type ObservationResource struct{}

func (o ObservationResource) List(c buffalo.Context) error {
	observation := &models.Observations{}
	if err := services.GetAllObservation(c.Value("tx").(*pop.Connection), observation); err != nil {
		return response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, observation)
}

func (o ObservationResource) Show(c buffalo.Context) error {
	id := c.Param("observation_id")
	observation := &models.Observation{}
	if err := services.GetObservationByID(c.Value("tx").(*pop.Connection), observation, id); err != nil {
		return response.SendNotFoundError(c, err)
	}
	return response.SendOKResponse(c, observation)
}

func (o ObservationResource) Create(c buffalo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot read body")
		return response.SendBadRequestError(c, err)
	}
	observation := &models.Observation{}
	err = json.Unmarshal([]byte(body), observation)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot unmarshal entity")
		return response.SendBadRequestError(c, err)
	}
	if err = services.CreateObservation(c.Value("tx").(*pop.Connection), observation); err != nil {
		return response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, observation)
}

func (o ObservationResource) Destroy(c buffalo.Context) error {
	id := c.Param("observation_id")
	if err := services.DestroyObservation(c.Value("tx").(*pop.Connection), id); err != nil {
		return response.SendGeneralError(c, err)
	}
	return c.Render(http.StatusOK, r.JSON("ok"))
}

func (o ObservationResource) Update(c buffalo.Context) error {
	id := c.Param("observation_id")
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot read body")
		return response.SendBadRequestError(c, err)
	}
	observation := &models.Observation{}
	if err = json.Unmarshal([]byte(body), observation); err != nil {
		log.SysLog.WithField("err", err).Error("cannot unmarshal json")
		return response.SendBadRequestError(c, err)
	}
	if err := services.UpdateObservation(c.Value("tx").(*pop.Connection), id, observation); err != nil {
		return response.SendGeneralError(c, err)
	}
	updatedObservation := &models.Observation{}
	if err := services.GetObservationByID(c.Value("tx").(*pop.Connection), updatedObservation, id); err != nil {
		log.SysLog.WithField("err", err).Error("error getting the updated entity")
		return response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, updatedObservation)
}
