package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/abdullokh-mukhammadjonov/template_api_gateway/modules/template_variables"
	"github.com/abdullokh-mukhammadjonov/template_api_gateway/modules/template_variables/var_user_service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	us "github.com/abdullokh-mukhammadjonov/template_api_gateway/genproto/user_service"
)

// @Security ApiKeyAuth
// @Router /v1/organization [post]
// @Summary Create Organization
// @Description API for creating organization
// @Tags organization
// @Accept json
// @Produce json
// @Param organization body var_user_service.CreateUpdateOrganizationSwag true "organization"
// @Success 201 {object} template_variables.CreateResponse
// @Failure 400 {object} template_variables.FailureResponse
// @Failure 404 {object} template_variables.FailureResponse
// @Failure 500 {object} template_variables.FailureResponse
// @Failure 503 {object} template_variables.FailureResponse
func (h *HandlerV1) CreateOrganization(c *gin.Context) {
	var (
		req          var_user_service.CreateUpdateOrganizationSwag
		organization us.CreateUpdateOrganization
		resp         *us.IdResponse
	)

	if err := c.ShouldBindJSON(&req); HandleHTTPError(c, http.StatusBadRequest, "APIgateway.CreateOrganization.BindingAction", err) {
		return
	}

	organization.Name = req.Name
	organization.FullName = req.FullName
	organization.Description = req.Description
	organization.Status = req.Status
	organization.ExternalId = req.ExternalId
	organization.Type = req.Type
	organization.Inn = req.Inn
	organization.Soato = req.Soato
	for _, org := range req.ResponsibleOrgs {
		organization.ResponsibleOrgs = append(organization.ResponsibleOrgs, &us.ResponsibleUser{
			OrganizationId: org.OrganizationId,
			Soato:          org.Soato,
		})
	}

	id := primitive.NewObjectID().Hex()
	organization.Id = id
	userInfo, err := h.UserInfo(c, true)
	if HandleRPCError(c, "CreateOrganization > userInfo: error while getting user info ", err) {
		log.Error(err.Error())
		return
	}

	// flag to know when create a new role
	projectOrganizationCreated := false
	var role = us.CreateUpdateRole{}

	if organization.Type == "PROJECT" {
		organizationsInDB, _ := h.services.OrganizationService().Get(
			context.Background(),
			&us.GetOrganizationRequest{
				Inn: organization.Inn,
			})

		if organizationsInDB != nil {
			exists := false
			for _, org := range organizationsInDB.ResponsibleOrgs {
				if org.OrganizationId == userInfo.OrganizationID && org.Soato == userInfo.Soato {
					exists = true
				}
			}

			if !exists {
				organizationsInDB.ResponsibleOrgs = append(organizationsInDB.ResponsibleOrgs, &us.ResponsibleUser{
					OrganizationId: userInfo.OrganizationID,
					Soato:          userInfo.Soato,
				})
			}

			_, err := h.services.OrganizationService().Update(context.Background(), &us.CreateUpdateOrganization{
				Id:              organizationsInDB.Id,
				Name:            organizationsInDB.Name,
				FullName:        organizationsInDB.FullName,
				Description:     organizationsInDB.Description,
				Status:          organization.Status,
				Code:            organizationsInDB.Code,
				ExternalId:      organizationsInDB.ExternalId,
				Inn:             organizationsInDB.Inn,
				Type:            organizationsInDB.Type,
				ResponsibleOrgs: organizationsInDB.ResponsibleOrgs,
			})
			if HandleRPCError(c, "CreateOrganization, error on update organization ", err) {
				return
			}

			// return updated organization id
			resp = &us.IdResponse{Id: organizationsInDB.Id}
		} else {
			log.Info("Organization type is PROJECT, creating ...")
			details, err := h.SoliqQomGetOrganizationDetails(organization.Inn)

			if HandleRPCError(c, "error while getting organization from DSQ", err) {
				log.Error(err.Error())
				return
			}
			log.Info("- - - - - - New OzGAShKLITI Organization name:" + details.Company.Name)
			organization.Name = details.Company.Name
			organization.FullName = details.Company.Name
			organization.Soato = details.Company.Soato
			organization.Description = details.Company.Description.NameUzLatin
			organization.ResponsibleOrgs = append(organization.ResponsibleOrgs, &us.ResponsibleUser{
				OrganizationId: userInfo.OrganizationID,
				Soato:          userInfo.Soato,
			})

			// a new role should be created accordingly
			projectOrganizationCreated = true

			role = us.CreateUpdateRole{
				Id:             primitive.NewObjectID().Hex(),
				OrganizationId: id,
				Name:           "Tuman xodimi",
				Description:    details.Company.Name + " (loyihalash tashkiloti) tuman xodimi",
				ExternalId:     0,
				Status:         true,
				Permissions:    template_variables.ProjectOrganizationDefaultPermissions,
				Code:           3,
			}

			resp, err = h.services.OrganizationService().Create(
				context.Background(),
				&organization,
			)

			if HandleRPCError(c, "error while creating organization", err) {
				return
			}
		}
	} else {
		resp, err = h.services.OrganizationService().Create(
			context.Background(),
			&organization,
		)

		if HandleRPCError(c, "error while creating organization", err) {
			return
		}
	}

	if projectOrganizationCreated {
		_, err = h.services.RoleService().CreateRole(
			context.Background(),
			&role,
		)

		if HandleRPCError(c, "error while creating role", err) {
			return
		}
	}

	c.JSON(http.StatusCreated, resp)
}

// @Router /v1/organization/{organization_id} [get]
// @Summary Get Organization
// @Tags organization
// @Accept json
// @Produce json
// @Param organization_id path string true "organization_id"
// @Param inn query string false "inn"
// @Success 200 {object} user_service.Organization
// @Failure 400 {object} template_variables.FailureResponse
// @Failure 404 {object} template_variables.FailureResponse
// @Failure 500 {object} template_variables.FailureResponse
// @Failure 503 {object} template_variables.FailureResponse
func (h *HandlerV1) GetOrganization(c *gin.Context) {
	var response *us.Organization
	id := c.Param("organization_id")
	org_inn := c.Query("inn")

	_, err := primitive.ObjectIDFromHex(id)
	if HandleHTTPError(c, http.StatusBadRequest, "error while parsing id", err) {
		return
	}

	organization, err := h.services.OrganizationService().Get(
		context.Background(),
		&us.GetOrganizationRequest{
			Id:  id,
			Inn: org_inn,
		})

	if HandleRPCError(c, "error while getting organization", err) {
		return
	}

	if err = ProtoToStructNumeric(&response, organization); HandleHTTPError(c, http.StatusInternalServerError, "error while parsing organization response", err) {
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Security ApiKeyAuth
// @Router /v1/organization [get]
// @Summary Getting All Organizations
// @Description API for getting all organizations
// @Tags organization
// @Accept json
// @Produce json
// @Param type query string false "type"
// @Param search query string false "search"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} user_service.GetAllOrganizationsResponse
// @Failure 400 {object} template_variables.FailureResponse
// @Failure 404 {object} template_variables.FailureResponse
// @Failure 500 {object} template_variables.FailureResponse
// @Failure 503 {object} template_variables.FailureResponse
func (h *HandlerV1) GetAllOrganizations(c *gin.Context) {
	var (
		response          *us.GetAllOrganizationsResponse
		organization_user *us.ResponsibleUser
	)
	search := c.Query("search")
	org_type := c.Query("type")

	page, err := ParseQueryParam(c, h.log, "page", "1")
	if err != nil {
		return
	}

	limit, err := ParseQueryParam(c, h.log, "limit", "10")
	if err != nil {
		return
	}

	if org_type == "PROJECT" {
		userInfo, err := h.UserInfo(c, true)
		if HandleRPCError(c, "CreateOrganization > userInfo: error while getting user info ", err) {
			log.Error(err.Error())
			return
		}
		if len(userInfo.Soato) == 7 {
			if err := MarshalUnmarshal(userInfo, &organization_user); err != nil {
				return
			}
		}
	}

	organizations, err := h.services.OrganizationService().GetAll(
		context.Background(),
		&us.GetAllOrganizationsRequest{
			Type:  org_type,
			Name:  search,
			Page:  uint32(page),
			Limit: uint32(limit),
			User:  organization_user,
		})
	if HandleRPCError(c, "Erro while getting all organizations", err) {
		return
	}

	if org_type == "PROJECT" {
		ozGashlitiFirst := []*us.Organization{}

		index := 0
		found := false
		for i, org := range organizations.Organizations {
			if org.Id == template_variables.UzGASHLITI {
				index = i
				found = true
				break
			}
		}

		if !found {
			organization_uzgashkliti, err := h.services.OrganizationService().Get(
				context.Background(),
				&us.GetOrganizationRequest{
					Id: template_variables.UzGASHLITI,
				})
			if HandleRPCError(c, "Erro while getting organization: Ozgashkliti", err) {
				fmt.Println(err)
				return
			}
			// Placing Ozgashkliti at index 0
			ozGashlitiFirst = append(ozGashlitiFirst, organization_uzgashkliti)
			organizations = &us.GetAllOrganizationsResponse{
				Organizations: organizations.Organizations,
				Count:         organizations.Count + 1,
			}
		} else {
			// Placing Ozgashkliti at index 0
			ozGashlitiFirst = append(ozGashlitiFirst, organizations.Organizations[index])
		}

		// append other organizations after
		for i, org := range organizations.Organizations {
			if org.Id != template_variables.UzGASHLITI {
				ozGashlitiFirst = append(ozGashlitiFirst, organizations.Organizations[i])
			}
		}

		organizations.Organizations = ozGashlitiFirst
	}

	if err = ProtoToStructNumeric(&response, organizations); HandleHTTPError(c, http.StatusInternalServerError, "error while parsing organizations response", err) {
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Router /v1/organization/{organization_id} [put]
// @Summary Update Organization
// @Description API for updating organization
// @Tags organization
// @Accept json
// @Produce json
// @Param organization_id path string  true "organization_id"
// @Param organization body var_user_service.CreateUpdateOrganizationSwag true "organization"
// @Success 200 {object} template_variables.CreateResponse
// @Failure 400 {object} template_variables.FailureResponse
// @Failure 404 {object} template_variables.FailureResponse
// @Failure 500 {object} template_variables.FailureResponse
// @Failure 503 {object} template_variables.FailureResponse
func (h *HandlerV1) UpdateOrganization(c *gin.Context) {
	var (
		organization us.CreateUpdateOrganization
	)
	organizationID := c.Param("organization_id")

	_, err := primitive.ObjectIDFromHex(organizationID)
	if HandleHTTPError(c, http.StatusBadRequest, "Error while parsing object id", err) {
		return
	}

	err = c.ShouldBindJSON(&organization)
	if HandleHTTPError(c, http.StatusBadRequest, "error while binding model to json", err) {
		return
	}

	organization.Id = organizationID

	resp, err := h.services.OrganizationService().Update(
		context.Background(),
		&organization)

	if HandleRPCError(c, "error while updating organization", err) {
		return
	}

	c.JSON(http.StatusOK, resp)
}

func SoliqQomLogin() (string, error) {
	type LoginRequestBody struct {
		Username string `json:"username" bson:"username"`
		Password string `json:"password" bson:"password"`
	}
	var body = LoginRequestBody{
		Username: template_variables.SoliqUsername,
		Password: template_variables.SoliqPassword,
	}
	log.Info("Get token start")
	byteBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("*2 SoliqQomLogin > Error on marshal ", err)
		return "", err
	}

	resp, err := http.Post(template_variables.SoliqLoginURL, "application/json", bytes.NewBuffer(byteBody))

	if err != nil {
		fmt.Println("*3 SoliqQomLogin > Error on making post request to ", template_variables.SoliqLoginURL, err)
		return "", err
	}

	// parsing response body
	tokenResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("*4 SoliqQomLogin > Error on reading response body:", err)
		return "", err
	}

	// decoding response body into string
	//err = json.Unmarshal(tokenResp, &token)
	token := string(tokenResp)

	return token, nil
}

func (h *HandlerV1) SoliqQomGetOrganizationDetails(inn string) (var_user_service.SoliqResponse, error) {
	var response var_user_service.SoliqResponse

	token, err := SoliqQomLogin()
	if err != nil {
		fmt.Println(err)
		return var_user_service.SoliqResponse{}, err
	}
	// bearer token
	var bearer = "Bearer " + token
	req, err := http.NewRequest("GET", template_variables.SoliqURL+inn, nil)
	if err != nil {
		fmt.Println(err)
		return var_user_service.SoliqResponse{}, err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	api_resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error on GET response to reestr-admin-api\n[ERROR] -", err)
		return var_user_service.SoliqResponse{}, err
	}
	defer api_resp.Body.Close()

	if api_resp.StatusCode == 404 {
		err := errors.New("No organization found with INN " + inn + " at " + template_variables.DavReestrUrl)
		return var_user_service.SoliqResponse{}, err
	}

	bodyBytes, err := ioutil.ReadAll(api_resp.Body)
	if err != nil {
		fmt.Println("SoliqQomGetOrganizationDetails > Error on reading response body:", err)
		return var_user_service.SoliqResponse{}, err
	}

	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		fmt.Println("SoliqQomGetOrganizationDetails > Error unmarshaling data from response.")
		return var_user_service.SoliqResponse{}, err
	}

	return response, nil
}

// @Security ApiKeyAuth
// @Router /v1/organizations/dashboard [get]
// @Summary Getting All Organizations
// @Description API for getting all organizations
// @Tags organization
// @Accept json
// @Produce json
// @Param type query string false "type"
// @Param search query string false "search"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} user_service.GetAllOrganizationResponse
// @Failure 400 {object} template_variables.FailureResponse
// @Failure 404 {object} template_variables.FailureResponse
// @Failure 500 {object} template_variables.FailureResponse
// @Failure 503 {object} template_variables.FailureResponse
func (h *HandlerV1) GetAllOrganizationsForDashboard(c *gin.Context) {
	var (
		response          *us.GetAllOrganizationResponse
		organization_user *us.ResponsibleUser
	)
	search := c.Query("search")
	org_type := c.Query("type")

	page, err := ParseQueryParam(c, h.log, "page", "1")
	if err != nil {
		return
	}

	limit, err := ParseQueryParam(c, h.log, "limit", "10")
	if err != nil {
		return
	}

	if org_type == "PROJECT" {
		userInfo, err := h.UserInfo(c, true)
		if HandleRPCError(c, "CreateOrganization > userInfo: error while getting user info ", err) {
			log.Error(err.Error())
			return
		}
		if len(userInfo.Soato) == 7 {
			if err := MarshalUnmarshal(userInfo, &organization_user); err != nil {
				return
			}
		}
	}

	organizationsCount, err := h.services.OrganizationService().GetAll(
		context.Background(),
		&us.GetAllOrganizationsRequest{
			Type:  org_type,
			Name:  search,
			Page:  uint32(page),
			Limit: uint32(limit),
			User:  organization_user,
		})

	if HandleRPCError(c, "Erro while getting all organizations", err) {
		return
	}
	if organizationsCount.Count != 0 {
		limit = int(organizationsCount.Count)
	}
	organizations, err := h.services.OrganizationService().GetAllForDashboard(
		context.Background(),
		&us.GetAllOrganizationsRequest{
			Type:  org_type,
			Name:  search,
			Page:  uint32(page),
			Limit: uint32(limit),
			User:  organization_user,
		})
	if HandleRPCError(c, "Erro while getting all organizations", err) {
		return
	}

	ozGashlitiFirst := []*us.Organization{}

	index := -1
	for i, org := range organizations.Project {
		if org.Id == template_variables.UzGASHLITI {
			index = i
			break
		}
	}

	if index > -1 {
		organization_uzgashkliti, err := h.services.OrganizationService().Get(
			context.Background(),
			&us.GetOrganizationRequest{
				Id: template_variables.UzGASHLITI,
			})
		if HandleRPCError(c, "Erro while getting organization: Ozgashkliti", err) {
			fmt.Println(err)
			return
		}
		// Placing Ozgashkliti at index 0
		ozGashlitiFirst = append(ozGashlitiFirst, organization_uzgashkliti)
		organizations = &us.GetAllOrganizationResponse{
			Project: organizations.Project,
			Simple:  organizations.Simple,
			Count:   organizations.Count + 1,
		}
	} else {
		// Placing Ozgashkliti at index 0
		ozGashlitiFirst = append(ozGashlitiFirst, organizations.Project[index])
	}

	// append other organizations after
	for i, org := range organizations.Project {
		if org.Id != template_variables.UzGASHLITI {
			ozGashlitiFirst = append(ozGashlitiFirst, organizations.Project[i])
		}
	}

	organizations.Project = ozGashlitiFirst

	if err = ProtoToStructNumeric(&response, organizations); HandleHTTPError(c, http.StatusInternalServerError, "error while parsing organizations response", err) {
		return
	}

	c.JSON(http.StatusOK, response)
}
