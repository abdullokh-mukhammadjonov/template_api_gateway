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

func createOrganization(t *testing.T) string {
	var (
		createOrganizationBody vrus.CreateUpdateOrganizationSwag
		createResponse         = &vr.CreateResponse{}
	)
	_ = faker.FakeData(&createOrganizationBody)
	createOrganizationBody.Status = true
	createOrganizationBody.Description = faker.Sentence()
	createOrganizationBody.Name = faker.Word()
	createOrganizationBody.FullName = createOrganizationBody.Name
	resp, err := PerformRequest(http.MethodPost, "/v1/organization", createOrganizationBody, createResponse)
	assert.NoError(t, err)
	assert.Equal(t, resp.Code, 201)

	return createResponse.ID
}

func TestCreateOrganization(t *testing.T) {
	_ = createOrganization(t)
}

func TestOrganizationSet(t *testing.T) {
	var (
		createOrganizationBody vrus.CreateUpdateOrganizationSwag
		createResponse         = &vr.CreateResponse{}
		id                     string
	)
	_ = faker.FakeData(&createOrganizationBody)
	resp, err := PerformRequest(http.MethodPost, "/v1/organization", createOrganizationBody, createResponse)
	assert.NoError(t, err)
	assert.Equal(t, resp.Code, 201)
	id = createResponse.ID

	var testOrganization = []struct {
		nameOfTest string
		routerURL  string
		methodType string
		code       int
		equal      bool
		request    interface{}
		response   interface{}
		want       interface{}
	}{
		// {
		// 	nameOfTest: "create organization with bad request",
		// 	routerURL:  "/v1/organization",
		// 	methodType: http.MethodPost,
		// 	code:       400,
		// 	response:   &vr.FailureResponse{},
		// 	want:       &vr.FailureResponse{},
		// 	equal:      false,
		// 	request: vrus.CreateUpdateOrganizationSwag{
		// 		Name:       "",
		// 		Code:       int32(faker.RandomUnixTime()),
		// 		ExternalID: int32(faker.RandomUnixTime()),
		// 	},
		// },
		{
			nameOfTest: "get organization",
			routerURL:  "/v1/organization/" + primitive.NewObjectID().Hex(),
			methodType: http.MethodGet,
			code:       404,
			equal:      false,
			want:       &vr.FailureResponse{},
			response:   &vr.FailureResponse{},
		},
		{
			nameOfTest: "update organization",
			routerURL:  "/v1/organization/" + id,
			methodType: http.MethodPut,
			code:       200,
			equal:      true,
			want:       &vr.EmptyResponse{},
			response:   &vr.EmptyResponse{},
			request: vrus.CreateUpdateOrganizationSwag{
				Name: faker.Word(),
			},
		},
		{
			nameOfTest: "get organization",
			routerURL:  "/v1/organization/" + id,
			methodType: http.MethodGet,
			code:       200,
			equal:      false,
			want: &vrus.Organization{
				ID:   id,
				Name: createOrganizationBody.Name,
			},
			response: &vrus.Organization{},
			request: vr.GetRequest{
				ID: id,
			},
		},
		{
			nameOfTest: "get all organization",
			routerURL:  "/v1/organization",
			methodType: http.MethodGet,
			code:       200,
			equal:      true,
		},
	}

	for i := range testOrganization {
		test := testOrganization[i]
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
