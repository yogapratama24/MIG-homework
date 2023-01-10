package controller

import (
	"fmt"
	"homework_mitramas/helpers"
	"homework_mitramas/model"
	"homework_mitramas/service"
	"strconv"

	"github.com/labstack/echo"
)

type MemberController struct {
	service service.MemberService
}

func NewMemberController(service service.MemberService) *MemberController {
	return &MemberController{service}
}

func (memberController *MemberController) GetMemberController(c echo.Context) error {
	auth, _ := helpers.ParseJWT(c)
	data, err := memberController.service.GetMember(auth.Id)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return err
	}

	helpers.NewHandlerResponse("Successfully get members", data).Success(c)
	return nil
}

func (memberController *MemberController) CreateMemberController(c echo.Context) error {
	var request model.MemberCreateRequest
	auth, _ := helpers.ParseJWT(c)
	if err := c.Bind(&request); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return err
	}
	fmt.Println("=====USER ID======", auth.Id)
	request.UserId = auth.Id

	if err := helpers.DoValidation(&request); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return nil
	}

	if err := memberController.service.CreateMember(request); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return err
	}

	helpers.NewHandlerResponse("Successfully create member", nil).SuccessCreate(c)
	return nil
}

func (memberController *MemberController) UpdateMemberController(c echo.Context) error {
	var request model.MemberUpdateRequest
	if err := c.Bind(&request); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return err
	}

	if err := helpers.DoValidation(&request); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return nil
	}

	if err := memberController.service.UpdateMember(request); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return err
	}

	helpers.NewHandlerResponse("Successfully update member", nil).SuccessCreate(c)
	return nil
}

func (memberController *MemberController) DeleteMemberController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return nil
	}

	if err := memberController.service.DeleteMember(id); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	helpers.NewHandlerResponse("Successfully delete member", nil).Success(c)
	return nil
}
