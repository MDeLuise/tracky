package models

import (
	"encoding/json"
	"fmt"
	"time"
	"tracky/log"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type ApiKey struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Value     string    `json:"value" db:"value"`
	User      *User     `json:"-" belongs_to:"user"`
	UserID    uuid.UUID `json:"user_id,omitempty" db:"user_id"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

func (a ApiKey) String() string {
	jo, _ := json.Marshal(a)
	return string(jo)
}

type ApiKeys []ApiKey

func (a *ApiKey) Validate(tx *pop.Connection) (*validate.Errors, error) {
	exist, err := DB.Where("id = ?", a.UserID).Exists(&User{})
	if err != nil {
		log.SysLog.WithField("err", err).Error("error while connecting to DB")
		return nil, err
	}
	if !exist {
		log.SysLog.Error("linked user does not exist")
		return nil, fmt.Errorf("linked user does not exist")
	}
	if len(a.Value) < 5 {
		return nil, fmt.Errorf("value cannot be less then 5 characters")
	}
	return nil, nil
}

func (a *ApiKey) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (a *ApiKey) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
