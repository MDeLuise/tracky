package actions

import (
	"net/http"
	"os"
	"time"
)

const authBaseURL = "/auth/login"

type loginRequest struct {
	Username string
	Password string
}

func (as *ActionSuite) Test_WrongCredentialsShouldNotAuthenticate() {
	as.LoadFixture("load test data")
	res := as.JSON(authBaseURL).Post(&loginRequest{
		Username: "admin",
		Password: "wrong",
	})
	as.Equal(http.StatusUnauthorized, res.Code)
}

func (as *ActionSuite) Test_CorrectCredentialsShouldAuthenticate() {
	as.LoadFixture("load test data")
	res := as.JSON(authBaseURL).Post(&loginRequest{
		Username: "admin",
		Password: "admin",
	})
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_ExpiredTokenShouldNotWork() {
	oldExpiration := os.Getenv("ACCESS_TOKEN_EXPIRATION_SECONDS")
	if err := os.Setenv("ACCESS_TOKEN_EXPIRATION_SECONDS", "2"); err != nil {
		as.Fail(err.Error())
	}
	token, err := getLoginToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(targetBaseURL)
	req.Headers = headers
	as.Equal(http.StatusOK, req.Get().Code)
	time.Sleep(2 * time.Second)
	os.Setenv("ACCESS_TOKEN_EXPIRATION_SECONDS", oldExpiration)
	as.Equal(http.StatusUnauthorized, req.Get().Code)
}

func (as *ActionSuite) Test_RefreshTokenWork() {
	oldExpiration := os.Getenv("ACCESS_TOKEN_EXPIRATION_SECONDS")
	if err := os.Setenv("ACCESS_TOKEN_EXPIRATION_SECONDS", "2"); err != nil {
		as.Fail(err.Error())
	}
	token, err := getRefreshToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	req := as.JSON(targetBaseURL)
	req.Headers = headers
	as.Equal(http.StatusUnauthorized, req.Get().Code)
	time.Sleep(2 * time.Second)
	res := as.JSON("/auth/refresh").Post(&struct {
		Refresh_token string
	}{
		Refresh_token: token,
	})
	os.Setenv("ACCESS_TOKEN_EXPIRATION_SECONDS", oldExpiration)
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_ExpiredRefreshTokenShouldNotWork() {
	oldExpiration := os.Getenv("REFRESH_TOKEN_EXPIRATION_SECONDS")
	if err := os.Setenv("REFRESH_TOKEN_EXPIRATION_SECONDS", "2"); err != nil {
		as.Fail(err.Error())
	}
	token, err := getRefreshToken(as)
	if err != nil {
		as.Fail(err.Error())
	}
	time.Sleep(2 * time.Second)
	res := as.JSON("/auth/refresh").Post(&struct {
		Refresh_token string
	}{
		Refresh_token: token,
	})
	as.Equal(http.StatusBadRequest, res.Code)
	os.Setenv("REFRESH_TOKEN_EXPIRATION_SECONDS", oldExpiration)
}
