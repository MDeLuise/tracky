package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

type Target struct {
	ID           uuid.UUID     `json:"id" db:"id"`
	Name         string        `json:"name" db:"name"`
	Description  string        `json:"description" db:"description"`
	Observations []Observation `json:"values,omitempty" has_many:"observations" order_by:"time desc"`
	CreatedAt    time.Time     `json:"-" db:"created_at"`
	UpdatedAt    time.Time     `json:"-" db:"updated_at"`
}

type Targets []Target

func (t Targets) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

func (t *Target) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{
			Field: t.Name, Name: "Name", Message: "The name cannot be empty"},
		&validators.StringLengthInRange{Field: t.Name, Name: "Name", Min: 2,
			Max: 15, Message: "the name must be at least 2 characters longer"},
		&validators.FuncValidator{
			Field:   t.Name,
			Name:    "Name",
			Message: "%s has already been registered!",
			Fn: func() bool {
				var b bool
				q := tx.Where("name = ?", t.Name)
				if t.ID != uuid.Nil {
					q = q.Where("id != ?", t.ID)
				}
				b, err = q.Exists(t)
				if err != nil {
					return false
				}
				return !b
			},
		},
	), err
}

func (t *Target) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (t *Target) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (t *Target) BeforeDestroy(tx *pop.Connection) error {
	linkedObservations := &Observations{}
	err := DB.Where("target_id = ?", t.ID).All(linkedObservations)
	if err != nil {
		return err
	}
	return DB.Destroy(linkedObservations)
}
