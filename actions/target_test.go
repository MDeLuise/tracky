package actions

import (
	"net/http"
	"time"

	"github.com/gobuffalo/httptest"
	"github.com/gofrs/uuid"
)

const targetBaseURL = "/target"

func (as *ActionSuite) Test_TargetEndpointShouldBeAuthenticated() {
	var responses = make([]*httptest.JSONResponse, 0)
	responses = append(responses, as.JSON(targetBaseURL).Get())
	responses = append(responses, as.JSON(targetBaseURL+"/42").Get())
	responses = append(responses, as.JSON(targetBaseURL).Post(nil))
	responses = append(responses, as.JSON(targetBaseURL+"/42").Delete())
	responses = append(responses, as.JSON(targetBaseURL+"/42").Put(nil))

	for _, res := range responses {
		as.Equal(http.StatusUnauthorized, res.Code)
	}
}

func (as *ActionSuite) Test_TargetGet() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(targetBaseURL)
	req.Headers = headers
	res := req.Get()
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_TargetPostLegal() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(targetBaseURL)
	req.Headers = headers
	res := req.Post(&struct {
		Name        string
		Description string
	}{
		"foo",
		"",
	})
	as.Equal(http.StatusOK, res.Code)

	resData, err := getResponseData(res)
	if err != nil {
		as.Fail(err.Error())
	}

	req = as.JSON(appendAtBaseURL(targetBaseURL, resData["id"].(string)))
	req.Headers = headers
	as.Equal(http.StatusOK, req.Get().Code)
}

func (as *ActionSuite) Test_TargetPostWithEmptyName() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(targetBaseURL)
	req.Headers = headers
	res := req.Post(&struct {
		Name        string
		Description string
	}{
		"",
		"foo...bar...",
	})
	as.Equal(http.StatusBadRequest, res.Code)
}

func (as *ActionSuite) Test_TargetDestroyLegal() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(targetBaseURL)
	req.Headers = headers
	res := req.Post(&struct {
		Name        string
		Description string
	}{
		"foo",
		"foo...bar...",
	})
	as.Equal(http.StatusOK, res.Code)

	resData, err := getResponseData(res)
	if err != nil {
		as.Fail(err.Error())
	}

	req = as.JSON(appendAtBaseURL(targetBaseURL, resData["id"].(string)))
	req.Headers = headers
	as.Equal(http.StatusOK, req.Delete().Code)
}

func (as *ActionSuite) Test_TargetDestroyNotExisting() {
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
	req := as.JSON(appendAtBaseURL(targetBaseURL, fakeID.String()))
	req.Headers = headers
	as.Equal(http.StatusInternalServerError, req.Delete().Code)
}

func (as *ActionSuite) Test_TargetDestroyAlsoDeleteLinkedObservations() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(targetBaseURL)
	req.Headers = headers
	res := req.Post(&struct {
		Name string
	}{
		"foo",
	})
	as.Equal(http.StatusOK, res.Code)

	resData, err := getResponseData(res)
	if err != nil {
		as.Fail(err.Error())
	}

	targetID := resData["id"].(string)
	req = as.JSON(observationBaseURL)
	req.Headers = headers
	res = req.Post(&struct {
		Value     float64
		Time      string
		Target_id string
	}{
		42.42,
		time.Now().Format(time.RFC3339),
		targetID,
	})
	as.Equal(http.StatusOK, res.Code)
	resData, err = getResponseData(res)
	if err != nil {
		as.Fail(err.Error())
	}
	observationID := resData["id"].(string)
	req = as.JSON(appendAtBaseURL(targetBaseURL, targetID))
	req.Headers = headers
	as.Equal(http.StatusOK, req.Delete().Code)

	req = as.JSON(appendAtBaseURL(observationBaseURL, observationID))
	req.Headers = headers

	as.Equal(http.StatusNotFound, req.Get().Code)
}

func (as *ActionSuite) Test_TargetUpdateLegal() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(targetBaseURL)
	req.Headers = headers
	res := req.Post(&struct {
		Name        string
		Description string
	}{
		"foo",
		"foo...bar...",
	})
	as.Equal(http.StatusOK, res.Code)

	resData, err := getResponseData(res)
	if err != nil {
		as.Fail(err.Error())
	}

	req = as.JSON(appendAtBaseURL(targetBaseURL, resData["id"].(string)))
	req.Headers = headers
	res = req.Put(&struct {
		Name        string
		Description string
	}{
		"BAR",
		"",
	})
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_TargetUpdateWithEmptyName() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(targetBaseURL)
	req.Headers = headers
	res := req.Post(&struct {
		Name        string
		Description string
	}{
		"foo",
		"foo...bar...",
	})
	as.Equal(http.StatusOK, res.Code)

	resData, err := getResponseData(res)
	if err != nil {
		as.Fail(err.Error())
	}

	req = as.JSON(appendAtBaseURL(targetBaseURL, resData["id"].(string)))
	req.Headers = headers
	res = req.Put(&struct {
		Name        string
		Description string
	}{
		"",
		"",
	})
	as.Equal(http.StatusInternalServerError, res.Code)
}

func (as *ActionSuite) Test_TargetUpdateNonExisting() {
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
	req := as.JSON(appendAtBaseURL(targetBaseURL, fakeID.String()))
	req.Headers = headers
	res := req.Put(&struct {
		Name        string
		Description string
	}{
		"",
		"",
	})
	as.Equal(http.StatusInternalServerError, res.Code)
}
