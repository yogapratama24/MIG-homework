package route

import (
	"homework_mitramas/config"
	"homework_mitramas/controller"
	"homework_mitramas/middlewares"
	"homework_mitramas/repository"
	"homework_mitramas/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	db := config.Connect()
	e := echo.New()
	e.HideBanner = true
	middlewares.LoggerMiddleware(e)
	middlewares.RecoverMiddleware(e)
	middlewares.CorsMiddleware(e)
	// Logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	v1 := e.Group("/api/v1")

	// AUTH
	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	// CLIENT
	clientRepository := repository.NewClientRepository(db)
	clientService := service.NewClientService(clientRepository)
	clientController := controller.NewClientController(clientService)

	// MEMBER
	memberRepository := repository.NewMemberRepository(db)
	memberService := service.NewMemberService(memberRepository)
	memberController := controller.NewMemberController(memberService)

	auth := v1.Group("/login")
	auth.POST("", authController.LoginController)

	midd := v1.Group("")
	middlewares.SetJwtMiddlewares(midd)

	client := midd.Group("/client")
	client.POST("", clientController.CreateClientController)
	client.GET("", clientController.GetClientController)
	client.PUT("", clientController.UpdateClientController)
	client.DELETE("/:id", clientController.DeleteClientController)

	member := midd.Group("/member")
	member.POST("", memberController.CreateMemberController)
	member.GET("", memberController.GetMemberController)
	member.PUT("", memberController.UpdateMemberController)
	member.DELETE("/:id", memberController.DeleteMemberController)
	return e
}
