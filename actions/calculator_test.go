package actions

import (
	"net/http"

	"github.com/gobuffalo/httptest"
	"github.com/gofrs/uuid"
)

const statsBaseURL = "/stats"

func (as *ActionSuite) Test_StatsEndpointShouldBeAuthenticated() {
	var responses = make([]*httptest.JSONResponse, 0)
	responses = append(responses, as.JSON(statsBaseURL+"/mean/42").Get())
	responses = append(responses, as.JSON(statsBaseURL+"/mean/42/42").Get())
	responses = append(responses, as.JSON(statsBaseURL+"/increment/42").Get())

	for _, res := range responses {
		as.Equal(http.StatusUnauthorized, res.Code)
	}
}

func (as *ActionSuite) Test_MeanOnExistingTarget() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(statsBaseURL + "/mean/" + fixtureTargetID1)
	req.Headers = headers
	res := req.Get()
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_MeanOnNotExistingTarget() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	fakeID, err := uuid.NewV4()
	if err != nil {
		as.Fail(err.Error())
	}
	req := as.JSON(statsBaseURL + "/mean/" + fakeID.String())
	req.Headers = headers
	res := req.Get()
	as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_MeeanAtOnExistingTarget() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(statsBaseURL + "/mean/" + fixtureTargetID1 + "/42")
	req.Headers = headers
	res := req.Get()
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_MeeanAtOnNotExistingTarget() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	fakeID, err := uuid.NewV4()
	if err != nil {
		as.Fail(err.Error())
	}
	req := as.JSON(statsBaseURL + "/mean/" + fakeID.String() + "/42")
	req.Headers = headers
	res := req.Get()
	as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_LastIncrOnExistingTarget() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(statsBaseURL + "/increment/" + fixtureTargetID1)
	req.Headers = headers
	res := req.Get()
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_LastIncrOnNotExistingTarget() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	fakeID, err := uuid.NewV4()
	if err != nil {
		as.Fail(err.Error())
	}
	req := as.JSON(statsBaseURL + "/increment/" + fakeID.String())
	req.Headers = headers
	res := req.Get()
	as.Equal(http.StatusNotFound, res.Code)
}
