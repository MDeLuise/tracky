package actions

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"testing"
	"tracky_go/response"

	"github.com/gobuffalo/httptest"
	"github.com/gobuffalo/suite/v4"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	action, err := suite.NewActionWithFixtures(App(), os.DirFS("../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &ActionSuite{
		Action: action,
	}
	suite.Run(t, as)
}

func getLoginToken(as *ActionSuite) (string, error) {
	as.LoadFixture("load test data")
	res := as.JSON(authBaseURL).Post(&loginRequest{
		Username: "admin",
		Password: "admin",
	})
	if res.Code != 200 {
		return "", errors.New(
			"response code not 200: " + strconv.FormatInt(int64(res.Code), 10))
	}
	data, err := getResponseData(res)
	return data["token"].(string), err
}

func getResponseData(res *httptest.JSONResponse) (map[string]interface{}, error) {
	jsonRes := &response.Response{}
	err := json.Unmarshal(res.Body.Bytes(), jsonRes)
	if err != nil {
		return nil, err
	}
	return (jsonRes.Data.(map[string]interface{})), nil
}

func getResponseDataArray(res *httptest.JSONResponse) ([]map[string]interface{}, error) {
	jsonRes := &response.Response{}
	err := json.Unmarshal(res.Body.Bytes(), jsonRes)
	if err != nil {
		return nil, err
	}
	var toReturn = make([]map[string]interface{}, 0)
	for _, el := range jsonRes.Data.([]interface{}) {
		toReturn = append(toReturn, el.(map[string]interface{}))
	}
	return toReturn, nil
}

func appendAtBaseURL(baseURL, param string) string {
	return baseURL + "/" + param
}
