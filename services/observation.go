package services

import (
	"tracky/log"
	"tracky/models"
)

func GetAllObservation(Observation *models.Observations) error {
	err := models.DB.All(Observation)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error while connecting to DB")
	}
	return err
}

func GetObservationByID(observation *models.Observation, id string) error {
	err := models.DB.Find(observation, id)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity")
	}
	return err
}

func CreateObservation(toCreate *models.Observation) error {
	vErr, err := models.DB.ValidateAndCreate(toCreate)
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

func UpdateObservation(id string, updated *models.Observation) error {
	ObservationToUpdate := &models.Observation{}
	if err := GetObservationByID(ObservationToUpdate, id); err != nil {
		return err
	}
	ObservationToUpdate.Value = updated.Value
	ObservationToUpdate.Time = updated.Time
	ObservationToUpdate.TargetID = updated.TargetID
	vErr, err := models.DB.ValidateAndUpdate(ObservationToUpdate)
	if vErr.HasAny() {
		log.SysLog.WithField("vErr", vErr).Error("entity not valid")
		return vErr
	} else if err != nil {
		log.SysLog.WithField("err", err).Error("entity not valid")
		return err
	}
	return nil
}

func DestroyObservation(id string) error {
	ObservationToDestroy := &models.Observation{}
	err := GetObservationByID(ObservationToDestroy, id)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity")
		return err
	}
	if err = models.DB.Destroy(ObservationToDestroy); err != nil {
		log.SysLog.WithField("err", err).Error("error while destroying entity")
		return err
	}
	return nil
}
