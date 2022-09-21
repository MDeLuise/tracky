package models

import (
	"time"

	"github.com/gofrs/uuid"
)

func (ms *ModelSuite) Test_ObservationWithNotExistingTargetIllegal() {
	nonExistingTargetID, err := uuid.NewV7()
	if err != nil {
		ms.Fail("error while creating uuid %s", err)
	}
	toTest := Observation{
		Value:     42,
		Time:      time.Now(),
		TargetID:  nonExistingTargetID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	vErr, err := toTest.Validate(ms.DB)
	ms.Error(err)
	ms.Assert().False(vErr.HasAny())
}
