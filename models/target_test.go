package models

import "time"

func (ms *ModelSuite) Test_TargetWithoutNameIllegal() {
	toTest := Target{
		Description: "foo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	vErr, err := toTest.Validate(ms.DB)
	ms.NoError(err)
	ms.Assert().True(vErr.HasAny())
}

func (ms *ModelSuite) Test_TargetWithoutDescriptionLegal() {
	toTest := Target{
		Name:      "foo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	vErr, err := toTest.Validate(ms.DB)
	ms.NoError(err)
	ms.Assert().False(vErr.HasAny())
}

func (ms *ModelSuite) Test_TargetWithDuplicatedNameIllegal() {
	ms.LoadFixture("load test data")
	target := Target{}
	if err := ms.DB.First(&target); err != nil {
		ms.Fail("error while retrieving target: %s", err)
	}
	toTest := Target{
		Name:      target.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	vErr, err := toTest.Validate(ms.DB)
	ms.NoError(err)
	ms.Assert().True(vErr.HasAny())
}
