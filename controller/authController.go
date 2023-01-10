package controller

import (
	"fmt"
	"homework_mitramas/helpers"
	"homework_mitramas/model"
	"homework_mitramas/service"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	service service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{service}
}

func (authController *AuthController) LoginController(c echo.Context) error {
	var (
		request model.UserLogin
	)
	if err := c.Bind(&request); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return nil
	}

	if err := helpers.DoValidation(&request); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return nil
	}

	userData, err := authController.service.Login(request)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	errHash := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(request.Password))
	if errHash != nil {
		fmt.Printf("Password Incorrect with err: %s\n", errHash)
		helpers.NewHandlerResponse("Password Incorrect", nil).BadRequest(c)
		return nil
	}

	tokenString, err := helpers.GenerateJWT(userData)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	id := strconv.Itoa(userData.Id)
	roleId := strconv.Itoa(userData.RoleId)

	cookieToken := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24 * 60),
		Path:     "/",
		SameSite: 2,
		HttpOnly: true,
	}

	cookieUserId := &http.Cookie{
		Name:     "user_id",
		Value:    id,
		Expires:  time.Now().Add(time.Hour * 24 * 60),
		Path:     "/",
		SameSite: 2,
		HttpOnly: true,
	}

	cookieRole := &http.Cookie{
		Name:     "role",
		Value:    roleId,
		Expires:  time.Now().Add(time.Hour * 24 * 60),
		Path:     "/",
		SameSite: 2,
		HttpOnly: true,
	}
	c.SetCookie(cookieToken)
	c.SetCookie(cookieUserId)
	c.SetCookie(cookieRole)
	helpers.NewHandlerResponse("Successfully Login", nil).Success(c)
	return nil
}
