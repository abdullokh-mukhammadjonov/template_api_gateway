// Note: please follow rules

package api

import (
	_ "github.com/abdullokh-mukhammadjonov/template_api_gateway/api/docs"
	v1 "github.com/abdullokh-mukhammadjonov/template_api_gateway/api/handler/v1"
	"github.com/abdullokh-mukhammadjonov/template_api_gateway/api/routes"
	"github.com/abdullokh-mukhammadjonov/template_api_gateway/config"
	"github.com/abdullokh-mukhammadjonov/template_api_gateway/pkg/logger"
	"github.com/abdullokh-mukhammadjonov/template_api_gateway/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Log         logger.Logger
	Cfg         config.Config
	Services    services.ServiceManager
	RedisClient services.RedisManager
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "*")

	router.Use(cors.New(corsConfig))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:         opt.Log,
		Cfg:         opt.Cfg,
		Services:    opt.Services,
		RedisClient: opt.RedisClient,
	})
	routesV1 := router.Group("/v1")

	// adding api_gateway endpoints
	routes.AddAPIGatewayRoutes(routesV1, handlerV1)

	// adding user_service endpoints
	routes.AddOrganizationRoutes(routesV1, handlerV1)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router

}
