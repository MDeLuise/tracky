package services

import (
	"fmt"
	"tracky/log"
	"tracky/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

func ApiKeyAuthenticationIfExists(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if key := c.Param("key"); key != "" {
			err := isApiKeyCorrect(c.Value("tx").(*pop.Connection), key)
			if err != nil {
				log.SysLog.WithField("err", err).Error("error checking API KEY")
				return err
			}
			userID, err := getUserIDFromApiKey(c.Value("tx").(*pop.Connection), key)
			if err != nil {
				log.SysLog.WithField("err", err).Error("error getting user from API KEY")
				return err
			}
			token, err := CreateAccessToken(userID)
			if err != nil {
				log.SysLog.WithField("err", err).Error("error creating token")
				return err
			}
			c.Request().Header.Add("Authorization", "Bearer " + token)
			return next(c)
		} else {
			return next(c)
		}
	}
}

func isApiKeyCorrect(tx *pop.Connection, key string) error {
	exists, err := tx.Q().Where("value = ?", key).Exists(&models.ApiKey{})
	if err != nil {
		return nil
	}
	if !exists {
		log.SysLog.Error("api key does not exists")
		return fmt.Errorf("api key does not exists")
	} else {
		return nil
	}
}

func GetAllApiKey(tx *pop.Connection, ApiKey *models.ApiKeys) error {
	err := tx.All(ApiKey)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error while connecting to DB")
	}
	return err
}

func GetApiKeyByID(tx *pop.Connection, apiKey *models.ApiKey, id string) error {
	err := tx.Find(apiKey, id)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity")
	}
	return err
}

func CreateApiKey(tx *pop.Connection, userID uuid.UUID) (*models.ApiKey, error) {
	toCreate := &models.ApiKey{
		Value: createApiKeyValue(),
		UserID: userID,
	}
	vErr, err := tx.ValidateAndCreate(toCreate)
	if err != nil {
		log.SysLog.WithField("err", err).Error("entity not valid")
		return toCreate, err
	}
	if vErr.HasAny() {
		log.SysLog.WithField("vErr", vErr.Errors).Error("entity not valid")
		return toCreate, vErr
	}
	return toCreate, nil
}

func createApiKeyValue() string {
	id, _ := uuid.NewV4()
	return id.String()
}

func DestroyApiKey(tx *pop.Connection, id string) error {
	fmt.Println("---- to destroy: " + id)
	ApiKeyToDestroy := &models.ApiKey{}
	err := GetApiKeyByID(tx, ApiKeyToDestroy, id)
	if err != nil {
		log.SysLog.WithField("err", err).Error("cannot find entity")
		return err
	}
	if err = tx.Destroy(ApiKeyToDestroy); err != nil {
		log.SysLog.WithField("err", err).Error("error while destroying entity")
		return err
	}
	return nil
}

func getUserIDFromApiKey(tx *pop.Connection, key string) (string, error) {
	apiKeyStruct := &models.ApiKey{}
	if err := tx.Where("value = ?", key).Last(apiKeyStruct); err != nil {
		return "", err
	}
	return apiKeyStruct.UserID.String(), nil
}