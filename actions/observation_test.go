package actions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gobuffalo/httptest"
	"github.com/gofrs/uuid"
)

const observationBaseURL = "/value/"

func (as *ActionSuite) Test_ObservationEndpointShouldBeAuthenticated() {
	var responses = make([]*httptest.JSONResponse, 0)
	responses = append(responses, as.JSON(observationBaseURL).Get())
	responses = append(responses, as.JSON(observationBaseURL+"42").Get())
	responses = append(responses, as.JSON(observationBaseURL).Post(nil))
	responses = append(responses, as.JSON(observationBaseURL+"42").Delete())
	responses = append(responses, as.JSON(observationBaseURL+"42").Put(nil))

	for _, res := range responses {
		as.Equal(401, res.Code)
	}
}

func (as *ActionSuite) Test_ObservationGet() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(observationBaseURL)
	req.Headers = headers
	res := req.Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_ObservationPostLegal() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(observationBaseURL)
	req.Headers = headers
	targetID, err := getTargetID(as, token)
	if err != nil {
		as.Fail(err.Error())
	}
	res := req.Post(&struct {
		Value     float64
		Time      string
		Target_id string
	}{
		42.42,
		time.Now().Format(time.RFC3339),
		targetID,
	})
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_ObservationPostNotExistingTarget() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(observationBaseURL)
	req.Headers = headers
	fakeID, err := uuid.NewV4()
	if err != nil {
		as.Fail(err.Error())
	}
	res := req.Post(&struct {
		Value    float64
		Time     string
		TargetID uuid.UUID
	}{
		42.42,
		time.Now().Format(time.RFC3339),
		fakeID,
	})
	as.Equal(400, res.Code)
}

func (as *ActionSuite) Test_ObservationDeleteLegal() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(observationBaseURL)
	req.Headers = headers
	targetID, err := getTargetID(as, token)
	if err != nil {
		as.Fail(err.Error())
	}
	res := req.Post(&struct {
		Value     float64
		Time      string
		Target_id string
	}{
		42.42,
		time.Now().Format(time.RFC3339),
		targetID,
	})
	as.Equal(200, res.Code)

	data, err := getResponseData(res)
	if err != nil {
		as.Fail(err.Error())
	}
	observationID := data["id"].(string)
	req = as.JSON(observationBaseURL + observationID)
	req.Headers = headers
	as.Equal(http.StatusOK, req.Delete().Code)
}

func (as *ActionSuite) Test_ObservationDeleteNotExisting() {
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
	req := as.JSON(observationBaseURL + fakeID.String())
	req.Headers = headers
	as.Equal(404, req.Delete().Code)
}

func (as *ActionSuite) Test_ObservationUpdateNonExisting() {
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
	req := as.JSON(observationBaseURL + fakeID.String())
	req.Headers = headers
	res := req.Put(&struct {
		Value     float64
		Time      string
		Target_id uuid.UUID
	}{
		42.42,
		time.Now().Format(time.RFC3339),
		fakeID,
	})
	as.Equal(404, res.Code)
}

func (as *ActionSuite) Test_ObservationUpdateLegal() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	targetID, err := getTargetID(as, token)
	if err != nil {
		as.Fail(err.Error())
	}
	req := as.JSON(observationBaseURL)
	req.Headers = headers
	res := req.Post(&struct {
		Value     float64
		Time      string
		Target_id string
	}{
		42.42,
		time.Now().Format(time.RFC3339),
		targetID,
	})
	as.Equal(http.StatusOK, res.Code)
	data, err := getResponseData(res)
	if err != nil {
		as.Fail(err.Error())
	}
	observationID := data["id"].(string)
	req = as.JSON(observationBaseURL + observationID)
	req.Headers = headers
	res = req.Put(&struct {
		Value     float64
		Target_id string
	}{
		24.24,
		targetID,
	})
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_ObservationUpdateNotExistingTarget() {
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	targetID, err := getTargetID(as, token)
	if err != nil {
		as.Fail(err.Error())
	}
	req := as.JSON(observationBaseURL)
	req.Headers = headers
	res := req.Post(&struct {
		Value     float64
		Time      string
		Target_id string
	}{
		42.42,
		time.Now().Format(time.RFC3339),
		targetID,
	})
	as.Equal(http.StatusOK, res.Code)
	data, err := getResponseData(res)
	if err != nil {
		as.Fail(err.Error())
	}
	observationID := data["id"].(string)
	fakeID, err := uuid.NewV4()
	if err != nil {
		as.Fail(err.Error())
	}
	req = as.JSON(observationBaseURL + observationID)
	req.Headers = headers
	res = req.Put(&struct {
		Value     float64
		Time      string
		Target_id string
	}{
		24.24,
		time.Now().Format(time.RFC3339),
		fakeID.String(),
	})
	as.Equal(http.StatusBadRequest, res.Code)
}

func getTargetID(as *ActionSuite, token string) (string, error) {
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(targetBaseURL)
	req.Headers = headers
	res := req.Get()
	if res.Code != http.StatusOK {
		return "", fmt.Errorf("cannot get targets, status code %v", res.Code)
	}
	data, err := getResponseDataArray(res)
	if err != nil {
		return "", fmt.Errorf("cannot get response data: %s", err)
	}
	return data[0]["id"].(string), nil
}
