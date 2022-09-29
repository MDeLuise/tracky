package models

import (
	"encoding/json"
	"time"

	"tracky/log"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password,omitempty" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

type Users []User

func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Username, Name: "Username"},
		&validators.StringIsPresent{Field: u.Password, Name: "Password"},
		&validators.StringLengthInRange{Field: u.Username, Name: "Username",
			Min: 5, Max: 50,
			Message: "The username must be greater than 4 characters",
		},
		&validators.StringLengthInRange{Field: u.Password, Name: "Password",
			Min: 5, Max: 50,
			Message: "The password must be greater than 4 characters",
		},
		&validators.FuncValidator{
			Field:   u.Username,
			Name:    "Username",
			Message: "%s has already been registered!",
			Fn: func() bool {
				var b bool
				q := tx.Where("username = ?", u.Username)
				if u.ID != uuid.Nil {
					q = q.Where("id != ?", u.ID)
				}
				b, err = q.Exists(u)
				if err != nil {
					return false
				}
				return !b
			},
		},
	), err
}

func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (u *User) BeforeCreate(tx *pop.Connection) error {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error while computing hash")
		return err
	}
	u.Password = string(hash)
	return nil
}
