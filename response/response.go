package response

import (
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

var r *render.Engine = render.New(render.Options{
	DefaultContentType: "application/json",
})

type Response struct {
	Timestamp     time.Time
	StatusCode    int
	StatusMessage string
	Data          interface{}
}

func SendOKResponse(c buffalo.Context, data interface{}) error {
	return SendResponse(c, http.StatusOK, data)
}

func SendResponse(c buffalo.Context, statusCode int, data interface{}) error {
	return send(c, statusCode, data)
}

func SendGeneralError(c buffalo.Context, err error) error {
	return SendError(c, http.StatusInternalServerError, err)
}

func SendNotFoundError(c buffalo.Context, err error) error {
	return SendError(c, http.StatusNotFound, err)
}

func SendBadRequestError(c buffalo.Context, err error) error {
	return SendError(c, http.StatusBadRequest, err)
}

func SendUnauthorizedError(c buffalo.Context, err error) error {
	return SendError(c, http.StatusUnauthorized, err)
}

func SendError(c buffalo.Context, statusCode int, err error) error {
	return send(c, statusCode, err.Error())
}

func send(c buffalo.Context, statusCode int, data interface{}) error {
	toReturn := Response{
		Timestamp:     time.Now(),
		StatusCode:    statusCode,
		StatusMessage: http.StatusText(statusCode),
		Data:          data,
	}
	return c.Render(statusCode, r.JSON(toReturn))
}
