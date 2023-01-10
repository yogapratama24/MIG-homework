package controller

import (
	"fmt"
	"homework_mitramas/helpers"
	"homework_mitramas/model"
	"homework_mitramas/service"
	"strconv"

	"github.com/labstack/echo"
)

type ClientController struct {
	service service.ClientService
}

func NewClientController(service service.ClientService) *ClientController {
	return &ClientController{service}
}

func (clientController *ClientController) GetClientController(c echo.Context) error {
	auth, _ := helpers.ParseJWT(c)
	fmt.Println(auth)
	data, err := clientController.service.GetClient()
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return err
	}

	helpers.NewHandlerResponse("Successfully get clients", data).Success(c)
	return nil
}

func (clientController *ClientController) CreateClientController(c echo.Context) error {
	var request model.UserCreateRequest
	if err := c.Bind(&request); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return err
	}

	if err := helpers.DoValidation(&request); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return nil
	}

	if err := clientController.service.CreateClient(request); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return err
	}

	helpers.NewHandlerResponse("Successfully create client", nil).SuccessCreate(c)
	return nil
}

func (clientController *ClientController) UpdateClientController(c echo.Context) error {
	var request model.UserUpdateRequest
	if err := c.Bind(&request); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return err
	}

	if err := helpers.DoValidation(&request); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return nil
	}

	if err := clientController.service.UpdateClient(request); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return err
	}

	helpers.NewHandlerResponse("Successfully update client", nil).SuccessCreate(c)
	return nil
}

func (clientController *ClientController) DeleteClientController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return nil
	}

	if err := clientController.service.DeleteClient(id); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	helpers.NewHandlerResponse("Successfully delete client", nil).Success(c)
	return nil
}
