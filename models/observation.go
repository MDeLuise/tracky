package models

import (
	"encoding/json"
	"fmt"
	"time"

	"tracky_go/log"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type Observation struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Value     float64   `json:"value" db:"value"`
	Time      time.Time `json:"time" db:"time"`
	Target    *Target   `json:"-" belongs_to:"target"`
	TargetID  uuid.UUID `json:"target_id,omitempty" db:"target_id"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

func (o Observation) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

func (u *Observation) MarshalJSON() ([]byte, error) {
	type Alias Observation
	formattedTime := u.Time.Format(time.RFC3339)
	return json.Marshal(&struct {
		Time string `json:"time"`
		*Alias
	}{
		Time:  formattedTime,
		Alias: (*Alias)(u),
	})
}

type Observations []Observation

func (o Observations) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

func (o *Observation) Validate(tx *pop.Connection) (*validate.Errors, error) {
	exist, err := DB.Where("id = ?", o.TargetID).Exists(&Target{})
	if err != nil {
		log.SysLog.WithField("err", err).Error("error while connecting to DB")
		return nil, err
	}
	if !exist {
		log.SysLog.Error("linked target does not exist")
		return nil, fmt.Errorf("linked target does not exist")
	}
	return nil, nil
}

func (o *Observation) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (o *Observation) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
