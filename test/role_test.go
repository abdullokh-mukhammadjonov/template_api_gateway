package v1_test

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	vr "github.com/abdullokh-mukhammadjonov/template_api_gateway/modules/template_variables"
	vrus "github.com/abdullokh-mukhammadjonov/template_api_gateway/modules/template_variables/var_user_service"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createRole(t *testing.T) (string, string) {
	var (
		createRoleBody vrus.CreateUpdateRoleSwag
		createResponse = &vr.CreateResponse{}
	)
	_ = faker.FakeData(&createRoleBody)
	createRoleBody.Permissions = []string{createPermission(t), createPermission(t), createPermission(t)}
	createRoleBody.Status = true
	createRoleBody.OrganizationID = createOrganization(t)
	resp, err := PerformRequest(http.MethodPost, "/v1/role", createRoleBody, createResponse)
	assert.NoError(t, err)
	assert.Equal(t, resp.Code, 201)

	return createResponse.ID, createRoleBody.OrganizationID
}

func TestCreateRole(t *testing.T) {
	_, _ = createRole(t)
}

func TestRoleSet(t *testing.T) {
	var (
		createRoleBody vrus.CreateUpdateRoleSwag
		createResponse = &vr.CreateResponse{}
		roleID         string
	)
	_ = faker.FakeData(&createRoleBody)
	createRoleBody.Permissions = []string{createPermission(t)}
	resp, err := PerformRequest(http.MethodPost, "/v1/role", createRoleBody, createResponse)
	assert.NoError(t, err)
	assert.Equal(t, resp.Code, 201)
	roleID = createResponse.ID

	var testRole = []struct {
		nameOfTest string
		routerURL  string
		methodType string
		code       int
		equal      bool
		request    interface{}
		response   interface{}
		want       interface{}
	}{
		{
			nameOfTest: "create role with bad request",
			routerURL:  "/v1/role",
			methodType: http.MethodPost,
			code:       400,
			response:   &vr.FailureResponse{},
			want:       &vr.FailureResponse{},
			equal:      true,
			request: vrus.CreateUpdateRoleSwag{
				Description: faker.Word(),
			},
		},
		{
			nameOfTest: "get role",
			routerURL:  "/v1/role/" + roleID,
			methodType: http.MethodGet,
			code:       200,
			equal:      true,
		},
		{
			nameOfTest: "update role",
			routerURL:  "/v1/role/" + roleID,
			methodType: http.MethodPut,
			code:       200,
			equal:      true,
			want:       &vr.EmptyResponse{},
			response:   &vr.EmptyResponse{},
			request: vrus.CreateUpdateRoleSwag{
				Name:        faker.Word(),
				Description: faker.Word(),
				Status:      true,
				Permissions: []string{createPermission(t)},
			},
		},
		{
			nameOfTest: "get role",
			routerURL:  "/v1/role/" + primitive.NewObjectID().Hex(),
			methodType: http.MethodGet,
			code:       404,
			equal:      false,
			want:       &vr.FailureResponse{},
			response:   &vr.FailureResponse{},
		},
		{
			nameOfTest: "get all role",
			routerURL:  "/v1/role",
			methodType: http.MethodGet,
			code:       200,
			equal:      true,
		},
	}

	for i := range testRole {
		test := testRole[i]
		t.Run(test.nameOfTest, func(t *testing.T) {
			resp, err := PerformRequest(test.methodType, test.routerURL, test.request, test.response)
			assert.NoError(t, err)
			assert.NotEmpty(t, resp)
			assert.Equal(t, test.code, resp.Code)
			fmt.Println(test.want)
			fmt.Println(test.response)
			fmt.Println(reflect.DeepEqual(test.want, test.response))
			if test.equal {
				assert.Equal(t, test.want, test.response)
			} else {
				assert.NotEqual(t, test.want, test.response)
			}
		})
	}

	fmt.Println(roleID)
}
