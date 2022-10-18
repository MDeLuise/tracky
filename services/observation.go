package services

import (
	"tracky/log"
	"tracky/models"

	"github.com/gobuffalo/pop/v6"
)

func GetAllObservation(tx *pop.Connection, Observation *models.Observations) error {
	err := tx.All(Observation)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error while connecting to DB")
	}
	return err
}

func GetObservationByID(tx *pop.Connection, observation *models.Observation, id string) error {
	err := tx.Find(observation, id)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity")
	}
	return err
}

func CreateObservation(tx *pop.Connection, toCreate *models.Observation) error {
	vErr, err := tx.ValidateAndCreate(toCreate)
	if err != nil {
		log.SysLog.WithField("err", err).Error("entity not valid")
		return err
	}
	if vErr.HasAny() {
		log.SysLog.WithField("vErr", vErr.Errors).Error("entity not valid")
		return vErr
	}
	return nil
}

func UpdateObservation(tx *pop.Connection, id string, updated *models.Observation) error {
	ObservationToUpdate := &models.Observation{}
	if err := GetObservationByID(tx, ObservationToUpdate, id); err != nil {
		return err
	}
	ObservationToUpdate.Value = updated.Value
	ObservationToUpdate.Time = updated.Time
	ObservationToUpdate.TargetID = updated.TargetID
	vErr, err := tx.ValidateAndUpdate(ObservationToUpdate)
	if vErr.HasAny() {
		log.SysLog.WithField("vErr", vErr).Error("entity not valid")
		return vErr
	} else if err != nil {
		log.SysLog.WithField("err", err).Error("entity not valid")
		return err
	}
	return nil
}

func DestroyObservation(tx *pop.Connection, id string) error {
	ObservationToDestroy := &models.Observation{}
	err := GetObservationByID(tx, ObservationToDestroy, id)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity")
		return err
	}
	if err = tx.Destroy(ObservationToDestroy); err != nil {
		log.SysLog.WithField("err", err).Error("error while destroying entity")
		return err
	}
	return nil
}
