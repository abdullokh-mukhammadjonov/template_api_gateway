package routes

import (
	"net/http"

	v1 "github.com/abdullokh-mukhammadjonov/template_api_gateway/api/handler/v1"
	"github.com/gin-gonic/gin"
)

func AddAPIGatewayRoutes(routerGroup *gin.RouterGroup, handlerV1 *v1.HandlerV1) {
	routerGroup.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    "template api_gateway",
		})
	})
	// FileUpload
	routerGroup.POST("/file-upload", handlerV1.FileUpload)
	routerGroup.POST("/image-upload", handlerV1.ImageUpload)
}
