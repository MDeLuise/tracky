package actions

import (
	"fmt"
	"net/http"
	"tracky/log"
	"tracky/models"
	"tracky/response"
	"tracky/services"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

type ApiKeyResource struct{}

func (a ApiKeyResource) List(c buffalo.Context) error {
	apiKey := &models.ApiKeys{}
	if err := services.GetAllApiKey(c.Value("tx").(*pop.Connection), apiKey); err != nil {
		return response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, apiKey)
}

func (a ApiKeyResource) Show(c buffalo.Context) error {
	id := c.Param("api_key_id")
	apiKey := &models.ApiKey{}
	if err := services.GetApiKeyByID(c.Value("tx").(*pop.Connection), apiKey, id); err != nil {
		return response.SendNotFoundError(c, err)
	}
	return response.SendOKResponse(c, apiKey)
}

func (a ApiKeyResource) Create(c buffalo.Context) error {
	bearer := c.Request().Header.Get("Authorization")
	userID, err := services.GetTokenClaim(bearer[7:], "id", false)
	if err != nil {
		log.SysLog.WithField("err", err).Error("error while getting the userID from the token")
		return err
	}
	id, _ := uuid.FromString(userID.(string))
	apiKey, err := services.CreateApiKey(c.Value("tx").(*pop.Connection), id)
	if err != nil {
		return response.SendGeneralError(c, err)
	}
	return response.SendOKResponse(c, apiKey)
}

func (a ApiKeyResource) Destroy(c buffalo.Context) error {
	id := c.Param("api_key_id")
	if err := services.DestroyApiKey(c.Value("tx").(*pop.Connection), id); err != nil {
		return response.SendGeneralError(c, err)
	}
	return c.Render(http.StatusOK, r.JSON("ok"))
}

func (a ApiKeyResource) Update(c buffalo.Context) error {
	return fmt.Errorf("Cannot update an API KEY, function not implemented.")
}
