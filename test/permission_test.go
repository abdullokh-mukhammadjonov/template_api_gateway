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

func createPermission(t *testing.T) string {
	var (
		createPermissionBody vrus.CreateUpdatePermissionSwag
		createResponse       = &vr.CreateResponse{}
	)
	_ = faker.FakeData(&createPermissionBody)
	createPermissionBody.Description = faker.Sentence()
	createPermissionBody.Name = faker.Word()
	createPermissionBody.RuName = createPermissionBody.Name
	resp, err := PerformRequest(http.MethodPost, "/v1/permission", createPermissionBody, createResponse)
	assert.NoError(t, err)
	assert.Equal(t, resp.Code, 201)

	return createResponse.ID
}

func TestCreatePermission(t *testing.T) {
	for i := 0; i < 2; i++ {
		_ = createPermission(t)
	}
}

func TestPermissionSet(t *testing.T) {
	var (
		createPermissionBody vrus.CreateUpdatePermissionSwag
		createResponse       = &vr.CreateResponse{}
		id                   string
	)
	_ = faker.FakeData(&createPermissionBody)
	resp, err := PerformRequest(http.MethodPost, "/v1/permission", createPermissionBody, createResponse)
	assert.NoError(t, err)
	assert.Equal(t, resp.Code, 201)
	id = createResponse.ID

	var testPermission = []struct {
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
			nameOfTest: "create permission with bad request",
			routerURL:  "/v1/permission",
			methodType: http.MethodPost,
			code:       201,
			response:   &vr.FailureResponse{},
			want:       &vr.FailureResponse{},
			equal:      true,
			request: vrus.CreateUpdatePermissionSwag{
				RuName:      faker.Name(),
				Description: faker.Word(),
			},
		},
		{
			nameOfTest: "get permission",
			routerURL:  "/v1/permission/" + id,
			methodType: http.MethodGet,
			code:       200,
			equal:      true,
			want: &vrus.Permission{
				ID:          id,
				Name:        createPermissionBody.Name,
				RuName:      createPermissionBody.RuName,
				Description: createPermissionBody.Description,
			},
			response: &vrus.Permission{},
			request: vr.GetRequest{
				ID: id,
			},
		},
		{
			nameOfTest: "get permission",
			routerURL:  "/v1/permission/" + primitive.NewObjectID().Hex(),
			methodType: http.MethodGet,
			code:       404,
			equal:      false,
			want:       &vr.FailureResponse{},
			response:   &vr.FailureResponse{},
		},
		{
			nameOfTest: "update permission",
			routerURL:  "/v1/permission/" + id,
			methodType: http.MethodPut,
			code:       200,
			equal:      true,
			want:       &vr.EmptyResponse{},
			response:   &vr.EmptyResponse{},
			request: vrus.CreateUpdatePermissionSwag{
				Name:        faker.Word(),
				RuName:      faker.Word(),
				Description: faker.Word(),
			},
		},
		{
			nameOfTest: "create permission with bad request",
			routerURL:  "/v1/permission",
			methodType: http.MethodPost,
			code:       201,
			response:   &vr.FailureResponse{},
			want:       &vr.FailureResponse{},
			equal:      true,
			request: vrus.CreateUpdatePermissionSwag{
				Name:        faker.Word(),
				RuName:      faker.Name(),
				Description: faker.Word(),
			},
		},
		{
			nameOfTest: "get permission",
			routerURL:  "/v1/permission/" + id,
			methodType: http.MethodGet,
			code:       200,
			equal:      false,
			want: &vrus.Permission{
				ID:          id,
				Name:        createPermissionBody.Name,
				RuName:      createPermissionBody.RuName,
				Description: createPermissionBody.Description,
			},
			request: vr.GetRequest{
				ID: id,
			},
		},
		{
			nameOfTest: "get all permission",
			routerURL:  "/v1/permission",
			methodType: http.MethodGet,
			code:       200,
			equal:      true,
		},
	}

	for i := range testPermission {
		test := testPermission[i]
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
	fmt.Println(id)
}
