package handler

import (
	"pro_pay/auth/middleware"
	"pro_pay/config"
	"pro_pay/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (handler *Handler) InitRoutes() (route *gin.Engine) {
	cfg := config.Config()
	route = gin.New()
	gin.SetMode(gin.ReleaseMode)
	if cfg.Server.Environment == "development" {
		gin.SetMode(gin.DebugMode)
		route.Use(gin.Logger())
		route.Use(gin.Recovery())
	}
	//405 error
	route.HandleMethodNotAllowed = true
	middleware.GinMiddleware(route)
	//swagger settings
	docs.SwaggerInfo.Title = cfg.Server.AppName
	docs.SwaggerInfo.Version = cfg.Server.AppVersion
	// docs.SwaggerInfo.Host = cfg.Server.AppURL
	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler),
		func(ctx *gin.Context) {
			docs.SwaggerInfo.Host = ctx.Request.Host
			if ctx.Request.TLS != nil {
				docs.SwaggerInfo.Schemes = []string{"https"}
			}
		})
	//static files
	// route.Static("/public", "./public/")
	//custom routes
	handler.Routers(route)
	return
}
