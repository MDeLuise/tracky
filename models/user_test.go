package models

import "time"

func (ms *ModelSuite) Test_UserWithoutUsernameIllegal() {
	toTest := User{
		Password:  "foo12345",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	vErr, err := toTest.Validate(ms.DB)
	ms.NoError(err)
	ms.Assert().True(vErr.HasAny())
}

func (ms *ModelSuite) Test_UserWithoutPasswordIllegal() {
	toTest := User{
		Username:  "foo12345",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	vErr, err := toTest.Validate(ms.DB)
	ms.NoError(err)
	ms.Assert().True(vErr.HasAny())
}
