package actions

import (
	"net/http"

	"github.com/gobuffalo/httptest"
	"github.com/gofrs/uuid"
)

const apiKeyBaseURL = "/key"

func (as *ActionSuite) Test_ApiKeyEndpointShouldBeAuthenticated() {
	var responses = make([]*httptest.JSONResponse, 0)
	responses = append(responses, as.JSON(apiKeyBaseURL).Get())
	responses = append(responses, as.JSON(apiKeyBaseURL+"/42").Get())
	responses = append(responses, as.JSON(apiKeyBaseURL).Post(nil))
	responses = append(responses, as.JSON(apiKeyBaseURL+"/42").Delete())
	responses = append(responses, as.JSON(apiKeyBaseURL+"/42").Put(nil))

	for _, res := range responses {
		as.Equal(http.StatusUnauthorized, res.Code)
	}
}

func (as *ActionSuite) Test_ApiKeyGet() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(apiKeyBaseURL)
	req.Headers = headers
	res := req.Get()
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_ApiKeyPost() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(apiKeyBaseURL)
	req.Headers = headers
	res := req.Post(&struct {}{})
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_ApiKeyDestroyLegal() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(appendAtBaseURL(apiKeyBaseURL, fixtureApiKeyID1))
	req.Headers = headers
	as.Equal(http.StatusOK, req.Delete().Code)
}

func (as *ActionSuite) Test_ApiKeyDestroyNotExisting() {
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
	req := as.JSON(appendAtBaseURL(apiKeyBaseURL, fakeID.String()))
	req.Headers = headers
	as.Equal(http.StatusInternalServerError, req.Delete().Code)
}

func (as *ActionSuite) Test_ApiKeyUpdateShouldNotBePossible() {
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
	req := as.JSON(appendAtBaseURL(apiKeyBaseURL, fakeID.String()))
	req.Headers = headers
	res := req.Put(&struct {}{})
	as.Equal(http.StatusInternalServerError, res.Code)
}