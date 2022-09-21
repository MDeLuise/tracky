package actions


const authBaseURL = "/auth/login"

type loginRequest struct{
	Username string
	Password string
} 

func (as *ActionSuite) Test_WrongCredentialsShouldNotAuthenticate() {
	as.LoadFixture("load test data")
	res := as.JSON(authBaseURL).Post(&loginRequest{
		Username: "admin",
		Password: "wrong",
	})
	as.Equal(401, res.Code)
}


func (as *ActionSuite) Test_CorrectCredentialsShouldAuthenticate() {
	as.LoadFixture("load test data")
	res := as.JSON(authBaseURL).Post(&loginRequest{
		Username: "admin",
		Password: "admin",
	})
	as.Equal(200, res.Code)
}