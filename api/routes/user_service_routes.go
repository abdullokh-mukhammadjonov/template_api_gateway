package routes

import (
	v1 "github.com/abdullokh-mukhammadjonov/template_api_gateway/api/handler/v1"
	"github.com/gin-gonic/gin"
)

func AddOrganizationRoutes(routerGroup *gin.RouterGroup, handlerV1 *v1.HandlerV1) {
	//Organization endpoints
	routerGroup.POST("/organization", handlerV1.CreateOrganization)
	routerGroup.GET("/organization/:organization_id", handlerV1.GetOrganization)
	routerGroup.GET("/organization", handlerV1.GetAllOrganizations)
	routerGroup.PUT("/organization/:organization_id", handlerV1.UpdateOrganization)
	routerGroup.GET("/organizations/dashboard", handlerV1.GetAllOrganizationsForDashboard)

	// Auth
	routerGroup.POST("/login", handlerV1.Login)
	routerGroup.POST("/login-exists", handlerV1.LoginExist)
	routerGroup.POST("/login-refresh", handlerV1.LoginRefresh)
	routerGroup.POST("/update-password/:user_id", handlerV1.UpdatePassword)
	routerGroup.POST("/update-password", handlerV1.UpdatePasswordFromToken)

	// Role endpoints
	routerGroup.POST("/role", handlerV1.CreateRole)
	routerGroup.GET("/role/:role_id", handlerV1.GetRole)
	routerGroup.GET("/role", handlerV1.GetAllRoles)
	routerGroup.PUT("/role/:role_id", handlerV1.UpdateRole)
	routerGroup.DELETE("/role/:role_id", handlerV1.DeleteRole)

	// Permission endpoints
	routerGroup.POST("/permission", handlerV1.CreatePermission)
	routerGroup.GET("/permission/:permission_id", handlerV1.GetPermission)
	routerGroup.GET("/permission", handlerV1.GetAllPermissions)
	routerGroup.PUT("/permission/:permission_id", handlerV1.UpdatePermission)
}
