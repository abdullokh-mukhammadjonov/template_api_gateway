package v1

import (
	"context"
	"net/http"
	"strconv"

	"github.com/abdullokh-mukhammadjonov/template_api_gateway/genproto/user_service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Router /v1/role [post]
// @Summary Create role
// @Description API for creating role
// @Tags role
// @Accept json
// @Produce json
// @Param role body var_user_service.CreateUpdateRoleSwag  true "role"
// @Success 201 {object} template_variables.CreateResponse
// @Failure 400 {object} template_variables.FailureResponse
// @Failure 404 {object} template_variables.FailureResponse
// @Failure 500 {object} template_variables.FailureResponse
// @Failure 503 {object} template_variables.FailureResponse
func (h *HandlerV1) CreateRole(c *gin.Context) {
	var (
		role user_service.CreateUpdateRole
	)
	err := c.ShouldBindJSON(&role)

	if HandleHTTPError(c, http.StatusBadRequest, "DiscussionLogicService.Action.Create.BindingAction", err) {
		return
	}
	id := primitive.NewObjectID().Hex()
	role.Id = id
	resp, err := h.services.RoleService().CreateRole(
		context.Background(),
		&role,
	)

	if HandleRPCError(c, "error while creating role", err) {
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Router /v1/role/{role_id} [get]
// @Summary Get role
// @Description API for getting role
// @Tags role
// @Accept json
// @Produce json
// @Param role_id path string  true "role_id"
// @Success 200 {object} user_service.Role
// @Failure 400 {object} template_variables.FailureResponse
// @Failure 404 {object} template_variables.FailureResponse
// @Failure 500 {object} template_variables.FailureResponse
// @Failure 503 {object} template_variables.FailureResponse
func (h *HandlerV1) GetRole(c *gin.Context) {
	var (
		id     = c.Param("role_id")
		_, err = primitive.ObjectIDFromHex(id)
	)
	if HandleHTTPError(c, http.StatusBadRequest, "error while parsing id", err) {
		return
	}

	role, err := h.services.RoleService().GetRole(context.Background(), &user_service.IdRequest{
		Id: id,
	})

	if HandleRPCError(c, "error while getting role ", err) {
		return
	}

	c.JSON(http.StatusOK, role)
}

// @Router /v1/role [get]
// @Summary Getting All roles
// @Description API for getting all rolees
// @Tags role
// @Accept json
// @Produce json
// @Param name query string  false "name"
// @Param organization_id query string  false "organization_id"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} user_service.GetAllRolesResponse
// @Failure 400 {object} template_variables.FailureResponse
// @Failure 404 {object} template_variables.FailureResponse
// @Failure 500 {object} template_variables.FailureResponse
// @Failure 503 {object} template_variables.FailureResponse
func (h *HandlerV1) GetAllRoles(c *gin.Context) {
	var response *user_service.GetAllRolesResponse

	name := c.Query("name")
	organizationId := c.Query("organization_id")

	page, err := ParseQueryParam(c, h.log, "page", "1")
	if err != nil {
		return
	}

	limit, err := ParseQueryParam(c, h.log, "limit", "0")
	if err != nil {
		return
	}

	roles, err := h.services.RoleService().GetAllRoles(
		context.Background(),
		&user_service.GetAllRolesRequest{
			Name:           name,
			OrganizationId: organizationId,
			Page:           uint32(page),
			Limit:          uint32(limit),
		})

	if HandleRPCError(c, "error while getting all roles", err) {
		return
	}

	err = ProtoToStructNumeric(&response, roles)
	if HandleHTTPError(c, http.StatusInternalServerError, "error while parsing roles response", err) {
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Router /v1/role/code [get]
// @Summary Getting All roles by code
// @Description API for getting all rolees by code
// @Tags role
// @Accept json
// @Produce json
// @Param code query int32  false "code"
// @Param organization_id query string  false "organization_id"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} user_service.GetAllRolesResponse
// @Failure 400 {object} template_variables.FailureResponse
// @Failure 404 {object} template_variables.FailureResponse
// @Failure 500 {object} template_variables.FailureResponse
// @Failure 503 {object} template_variables.FailureResponse
func (h *HandlerV1) GetAllRolesByCode(c *gin.Context) {
	var (
		code           = c.Query("code")
		organizationId = c.Query("organization_id")
		response       *user_service.GetAllRolesResponse
	)
	codeInt, err := strconv.Atoi(code)
	if err != nil {
		return
	}

	page, err := ParseQueryParam(c, h.log, "page", "1")
	if err != nil {
		return
	}

	limit, err := ParseQueryParam(c, h.log, "limit", "20")
	if err != nil {
		return
	}

	roles, err := h.services.RoleService().GetAllRolesByCode(
		context.Background(),
		&user_service.GetAllRolesByCodeRequest{
			Code:           uint32(codeInt),
			OrganizationId: organizationId,
			Page:           uint32(page),
			Limit:          uint32(limit),
		})

	if HandleRPCError(c, "error while getting all roles", err) {
		return
	}

	err = ProtoToStructNumeric(&response, roles)
	if HandleHTTPError(c, http.StatusInternalServerError, "error while parsing roles response", err) {
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Router /v1/role/{role_id} [put]
// @Summary Update role
// @Description API for updating role
// @Tags role
// @Accept json
// @Produce json
// @Param role_id path string  true "role_id"
// @Param role body var_user_service.CreateUpdateRoleSwag true "role"
// @Success 200 {object} template_variables.CreateResponse
// @Failure 400 {object} template_variables.FailureResponse
// @Failure 404 {object} template_variables.FailureResponse
// @Failure 500 {object} template_variables.FailureResponse
// @Failure 503 {object} template_variables.FailureResponse
func (h *HandlerV1) UpdateRole(c *gin.Context) {
	var (
		role user_service.CreateUpdateRole
	)
	id := c.Param("role_id")
	_, err := primitive.ObjectIDFromHex(id)
	if HandleHTTPError(c, http.StatusBadRequest, "error while parsing id", err) {
		return
	}
	err = c.ShouldBindJSON(&role)
	if HandleHTTPError(c, http.StatusBadRequest, "error while updating roles", err) {
		return
	}
	role.Id = id

	resp, err := h.services.RoleService().UpdateRole(
		context.Background(),
		&role)

	if HandleRPCError(c, "error while updating role", err) {
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /v1/role/{role_id} [delete]
// @Summary Delete Role
// @Description API for deleting role
// @Tags role
// @Accept json
// @Produce json
// @Param role_id path string  true "role_id"
// @Success 200 {object} template_variables.SuccessResponse
// @Failure 400 {object} template_variables.FailureResponse
// @Failure 404 {object} template_variables.FailureResponse
// @Failure 500 {object} template_variables.FailureResponse
// @Failure 503 {object} template_variables.FailureResponse
func (h *HandlerV1) DeleteRole(c *gin.Context) {
	id := c.Param("role_id")

	_, err := primitive.ObjectIDFromHex(id)
	if HandleHTTPError(c, http.StatusBadRequest, "Error while parsing object id", err) {
		return
	}

	resp, err := h.services.RoleService().DeleteRole(
		context.Background(),
		&user_service.IdRequest{
			Id: id,
		},
	)

	if HandleRPCError(c, "error while deleting role", err) {
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /v1/permission [post]
// @Summary Create permission
// @Description API for creating permission
// @Tags permission
// @Accept json
// @Produce json
// @Param permission body var_user_service.CreateUpdatePermissionSwag  true "permission"
// @Success 201 {object} template_variables.CreateResponse
// @Failure 400 {object} template_variables.FailureResponse
// @Failure 404 {object} template_variables.FailureResponse
// @Failure 500 {object} template_variables.FailureResponse
// @Failure 503 {object} template_variables.FailureResponse
func (h *HandlerV1) CreatePermission(c *gin.Context) {
	var (
		permission user_service.Permission
	)
	err := c.ShouldBindJSON(&permission)

	if HandleHTTPError(c, http.StatusBadRequest, "DiscussionLogicService.Action.Create.BindingAction", err) {
		return
	}

	id := primitive.NewObjectID().Hex()

	permission.Id = id
	resp, err := h.services.RoleService().CreatePermission(
		context.Background(),
		&permission,
	)

	if HandleRPCError(c, "error while creating permission", err) {
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Router /v1/permission/{permission_id} [get]
// @Summary Get Permission
// @Tags permission
// @Accept json
// @Produce json
// @Param permission_id path string true "permission_id"
// @Success 200 {object} user_service.Permission
// @Failure 400 {object} template_variables.FailureResponse
// @Failure 404 {object} template_variables.FailureResponse
// @Failure 500 {object} template_variables.FailureResponse
// @Failure 503 {object} template_variables.FailureResponse
func (h *HandlerV1) GetPermission(c *gin.Context) {
	var response *user_service.Permission
	id := c.Param("permission_id")

	_, err := primitive.ObjectIDFromHex(id)
	if HandleHTTPError(c, http.StatusBadRequest, "error while parsing id", err) {
		return
	}

	permission, err := h.services.RoleService().GetPermission(
		context.Background(),
		&user_service.IdRequest{
			Id: id,
		},
	)

	if HandleRPCError(c, "error while getting permission", err) {
		return
	}

	err = ProtoToStruct(&response, permission)
	if HandleHTTPError(c, http.StatusInternalServerError, "error while parsing permission response", err) {
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Router /v1/permission [get]
// @Summary Getting All permissions
// @Description API for getting all rolees
// @Tags permission
// @Accept json
// @Produce json
// @Param name query string  false "name"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} user_service.GetAllPermissionsResponse
// @Failure 400 {object} template_variables.FailureResponse
// @Failure 404 {object} template_variables.FailureResponse
// @Failure 500 {object} template_variables.FailureResponse
// @Failure 503 {object} template_variables.FailureResponse
func (h *HandlerV1) GetAllPermissions(c *gin.Context) {
	var response *user_service.GetAllPermissionsResponse
	name := c.Query("name")

	page, err := ParseQueryParam(c, h.log, "page", "1")
	if err != nil {
		return
	}

	limit, err := ParseQueryParam(c, h.log, "limit", "20")
	if err != nil {
		return
	}

	permissions, err := h.services.RoleService().GetAllPermissions(
		context.Background(),
		&user_service.GetAllPermissionsRequest{
			Name:  name,
			Page:  uint32(page),
			Limit: uint32(limit),
		})

	if HandleRPCError(c, "error while getting all roles", err) {
		return
	}

	err = ProtoToStruct(&response, permissions)
	if HandleHTTPError(c, http.StatusInternalServerError, "error while parsing permissions response", err) {
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Router /v1/permission/{permission_id} [put]
// @Summary Update permission
// @Description API for updating permission
// @Tags permission
// @Accept json
// @Produce json
// @Param permission_id path string  true "permission_id"
// @Param permission body var_user_service.CreateUpdatePermissionSwag true "permission"
// @Success 200 {object} template_variables.CreateResponse
// @Failure 400 {object} template_variables.FailureResponse
// @Failure 404 {object} template_variables.FailureResponse
// @Failure 500 {object} template_variables.FailureResponse
// @Failure 503 {object} template_variables.FailureResponse
func (h *HandlerV1) UpdatePermission(c *gin.Context) {
	var (
		permission user_service.Permission
	)

	id := c.Param("permission_id")

	_, err := primitive.ObjectIDFromHex(id)
	if HandleHTTPError(c, http.StatusBadRequest, "Error while parsing object id", err) {
		return
	}

	err = c.ShouldBindJSON(&permission)

	if HandleHTTPError(c, http.StatusBadRequest, "error while updating roles", err) {
		return
	}
	permission.Id = id

	resp, err := h.services.RoleService().UpdatePermission(
		context.Background(),
		&permission,
	)

	if HandleRPCError(c, "error while updating permission", err) {
		return
	}

	c.JSON(http.StatusOK, resp)
}
