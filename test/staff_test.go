package v1_test

import (
	"net/http"
	"testing"
	"time"

	vr "github.com/abdullokh-mukhammadjonov/template_api_gateway/modules/template_variables"
	vrus "github.com/abdullokh-mukhammadjonov/template_api_gateway/modules/template_variables/var_user_service"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createStaff(t *testing.T) string {
	var (
		createStaffBody vrus.CreateUpdateStaffSwag
		createResponse         = &vr.CreateResponse{}
		layout          string = "2021-01-02"
	)
	_ = faker.FakeData(&createStaffBody)
	createStaffBody.Password = "1234567890"
	createStaffBody.Login = faker.Username()
	createStaffBody.FirstName = faker.FirstName()
	createStaffBody.LastName = faker.LastName()
	createStaffBody.MiddleName = faker.TitleMale()
	createStaffBody.PhoneNumber = faker.Phonenumber()
	createStaffBody.Email = faker.Email()
	createStaffBody.RoleID, createStaffBody.OrganizationID = createRole(t)
	createStaffBody.PassportIssueDate = time.Now().Format(layout)
	createStaffBody.City.ID = primitive.NewObjectID().Hex()
	createStaffBody.Region.ID = primitive.NewObjectID().Hex()

	resp, err := PerformRequest(http.MethodPost, "/v1/staff", createStaffBody, createResponse)
	assert.NoError(t, err)
	assert.Equal(t, resp.Code, 201)

	return createResponse.ID
}

func TestCreateStaff(t *testing.T) {
	_ = createStaff(t)
}

func TestStaffSet(t *testing.T) {
	var (
		createStaffBody vrus.CreateUpdateStaffSwag
		createResponse         = &vr.CreateResponse{}
		layout          string = "2006-01-02"
		id              string
	)
	_ = faker.FakeData(&createStaffBody)
	createStaffBody.RoleID, createStaffBody.OrganizationID = createRole(t)
	createStaffBody.PassportIssueDate = time.Now().Format(layout)
	createStaffBody.City.ID = primitive.NewObjectID().Hex()
	createStaffBody.Region.ID = primitive.NewObjectID().Hex()
	createStaffBody.Password = faker.Password()
	createStaffBody.Login = faker.Username()
	createStaffBody.FirstName = faker.FirstName()
	createStaffBody.LastName = faker.LastName()
	createStaffBody.MiddleName = faker.TitleMale()
	createStaffBody.PhoneNumber = faker.Phonenumber()
	createStaffBody.Email = faker.Email()
	createStaffBody.City.ID = primitive.NewObjectID().Hex()
	createStaffBody.Region.ID = primitive.NewObjectID().Hex()

	updateStaffBody := createStaffBody
	updateStaffBody.PhoneNumber = faker.Phonenumber()
	updateStaffBody.FirstName = faker.Name()

	resp, err := PerformRequest(http.MethodPost, "/v1/staff", createStaffBody, createResponse)
	assert.NoError(t, err)
	assert.Equal(t, resp.Code, 201)
	id = createResponse.ID

	var testStaff = []struct {
		nameOfTest  string
		routerURL   string
		methodType  string
		code        int
		equal       bool
		request     interface{}
		response    interface{}
		want        interface{}
		getResponse *vrus.Staff
	}{
		{
			nameOfTest: "create staff with bad request",
			routerURL:  "/v1/staff",
			methodType: http.MethodPost,
			code:       400,
			response:   &vr.FailureResponse{},
			want:       &vr.FailureResponse{},
			equal:      false,
			request:    vrus.CreateUpdateStaffSwag{},
		},
		{
			nameOfTest: "get staff",
			routerURL:  "/v1/staff/" + id,
			methodType: http.MethodGet,
			code:       200,
			equal:      true,
			want:       id,
			response:   &vrus.Staff{},
			request: vr.GetRequest{
				ID: id,
			},
		},
		{
			nameOfTest: "update staff",
			routerURL:  "/v1/staff/" + id,
			methodType: http.MethodPut,
			code:       200,
			equal:      true,
			want:       &vr.EmptyResponse{},
			response:   &vr.EmptyResponse{},
			request:    updateStaffBody,
		},
		{
			nameOfTest: "get staff with error",
			routerURL:  "/v1/staff/" + primitive.NewObjectID().Hex(),
			methodType: http.MethodGet,
			code:       404,
			equal:      false,
			want:       &vr.FailureResponse{},
			response:   &vr.FailureResponse{},
		},
		{
			nameOfTest: "get all staff",
			routerURL:  "/v1/staff",
			methodType: http.MethodGet,
			code:       200,
			equal:      true,
		},
	}

	for i := range testStaff {
		test := testStaff[i]
		t.Run(test.nameOfTest, func(t *testing.T) {
			resp, err := PerformRequest(test.methodType, test.routerURL, test.request, test.response)
			if test.nameOfTest == "get staff" {
				test.getResponse = test.response.(*vrus.Staff)
			}
			assert.NoError(t, err)
			assert.NotEmpty(t, resp)
			assert.Equal(t, test.code, resp.Code)
			if test.equal {
				if test.nameOfTest == "get staff" {
					assert.Equal(t, test.want, test.getResponse.ID)
				} else {
					assert.Equal(t, test.want, test.response)
				}
			} else {
				assert.NotEqual(t, test.want, test.response)
			}
		})
	}
}
