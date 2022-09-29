package services

import (
	"tracky/log"
	"tracky/models"
)

func GetAllTargets(targets *models.Targets) error {
	err := models.DB.All(targets)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error while connecting to DB")
	}
	return err
}

func GetTargetByID(target *models.Target, id string) error {
	err := models.DB.Find(target, id)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity")
	}
	return err
}

func CreateTarget(toCreate *models.Target) error {
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

func UpdateTarget(id string, updated *models.Target) error {
	targetToUpdate := &models.Target{}
	if err := GetTargetByID(targetToUpdate, id); err != nil {
		return err
	}
	targetToUpdate.Name = updated.Name
	targetToUpdate.Description = updated.Description
	vErr, err := models.DB.ValidateAndUpdate(targetToUpdate)
	if vErr.HasAny() {
		log.SysLog.WithField("vErr", vErr).Error("entity not valid")
		return vErr
	} else if err != nil {
		log.SysLog.WithField("err", err).Error("entity not valid")
		return err
	}
	return nil
}

func DestroyTarget(id string) error {
	targetToDestroy := &models.Target{}
	err := GetTargetByID(targetToDestroy, id)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity")
		return err
	}
	if err = models.DB.Destroy(targetToDestroy); err != nil {
		log.SysLog.WithField("err", err).Error("error while destroying entity")
		return err
	}
	return nil
}
