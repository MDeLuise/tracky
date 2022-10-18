package services

import (
	"tracky/log"
	"tracky/models"

	"github.com/gobuffalo/pop/v6"
)

func GetAllTargets(tx *pop.Connection, targets *models.Targets) error {
	err := tx.All(targets) 
	if err != nil {
		log.SysLog.WithField("err", err).Error("error while connecting to DB")
	}
	return err
}

func GetTargetByID(tx *pop.Connection, target *models.Target, id string) error {
	err := tx.Eager().Find(target, id)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity")
	}
	return err
}

func CreateTarget(tx *pop.Connection, toCreate *models.Target) error {
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

func UpdateTarget(tx *pop.Connection, id string, updated *models.Target) error {
	targetToUpdate := &models.Target{}
	if err := GetTargetByID(tx, targetToUpdate, id); err != nil {
		return err
	}
	targetToUpdate.Name = updated.Name
	targetToUpdate.Description = updated.Description
	vErr, err := tx.ValidateAndUpdate(targetToUpdate)
	if vErr.HasAny() {
		log.SysLog.WithField("vErr", vErr).Error("entity not valid")
		return vErr
	} else if err != nil {
		log.SysLog.WithField("err", err).Error("entity not valid")
		return err
	}
	return nil
}

func DestroyTarget(tx *pop.Connection, id string) error {
	targetToDestroy := &models.Target{}
	err := GetTargetByID(tx, targetToDestroy, id)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity")
		return err
	}
	if err = tx.Destroy(targetToDestroy); err != nil {
		log.SysLog.WithField("err", err).Error("error while destroying entity")
		return err
	}
	return nil
}

func GetLikedObservations(tx *pop.Connection, t *models.Target) (*models.Observations, error) {
	linkedObservations := &models.Observations{}
	if err := tx.Where("target_id = ?", t.ID).All(linkedObservations); err != nil {
		return nil, err
	}
	return linkedObservations, nil
}
