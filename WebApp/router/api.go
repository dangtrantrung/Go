package router

import (
	"github-trend-BE/handler"
	"github-trend-BE/middleware"

	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
	adminMiddleware "github-trend-BE/middleware"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	//users APIs
	//group
	g:=api.Echo.Group("/user")//middleware.AddTrailingSlash())
	g.POST("/sign-in", api.UserHandler.HandleSignIn, adminMiddleware.IsAdmin())
	g.POST("/sign-up", api.UserHandler.HandleSignUp)
	g.POST("/profile", api.UserHandler.Profile, middleware.JWTMiddleware())
	api.Echo.GET("/user/profile", api.UserHandler.Profile, middleware.JWTMiddleware())
	api.Echo.PUT("/user/profile/update", api.UserHandler.UpdateProfile, middleware.JWTMiddleware())

}
