package var_user_service

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Staff struct {
	ID                 string             `json:"id" bson:"_id"`
	UniqueName         string             `json:"unique_name" bson:"unique_name"`
	MiddleName         string             `json:"middle_name" bson:"middle_name"`
	FirstName          string             `json:"first_name" bson:"first_name"`
	LastName           string             `json:"last_name" bson:"last_name"`
	PhoneNumber        string             `json:"phone_number" bson:"phone_number"`
	UserType           string             `json:"user_type" bson:"user_type"`
	Login              string             `json:"login" bson:"login"`
	LastLogin          string             `json:"last_login" bson:"last_login"`
	Address            string             `json:"address" bson:"address"`
	Password           string             `json:"password" bson:"password"`
	PassportNumber     string             `json:"passport_number" bson:"passport_number"`
	PassportIssuePlace string             `json:"passport_issue_place" bson:"passport_issue_place"`
	Email              string             `json:"email" bson:"email"`
	ExtraInfo          string             `json:"extra_info" bson:"extra_info"`
	Pinfl              string             `json:"pinfl" bson:"pinfl"`
	Soato              string             `json:"soato" bson:"soato"`
	ExternalID         int32              `json:"external_id" bson:"external_id"`
	Policy             int64              `json:"policy" bson:"policy"`
	Inn                int64              `json:"inn" bson:"inn"`
	OrganizationID     *Organization      `json:"organization" bson:"organization"`
	City               *City              `json:"city" bson:"city"`
	Region             *Region            `json:"region" bson:"region"`
	Role               *GetRole           `json:"role" bson:"role"`
	PassportIssueDate  primitive.DateTime `json:"passport_issue_date" bson:"passport_issue_date"`
	CreatedAt          primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt          primitive.DateTime `json:"updated_at" bson:"updated_at"`
	Verified           bool               `json:"verified" bson:"verified"`
	Status             bool               `json:"status" bson:"status"`
	Confirmed          bool               `json:"confirmed" bson:"confirmed"`
	OneIdLogin         string             `json:"one_id_login" bson:"one_id_login"`
	OrganizationInn    string             `json:"organization_inn" bson:"organization_inn"`
	LastLoginType      string             `json:"last_login_type" bson:"last_login_type"`
}

type CreateUpdateStaff struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id"`
	RoleID             primitive.ObjectID `json:"role_id" bson:"role_id"`
	OrganizationID     primitive.ObjectID `json:"organization_id" bson:"organization_id"`
	UniqueName         string             `json:"unique_name" bson:"unique_name"`
	MiddleName         string             `json:"middle_name" bson:"middle_name"`
	FirstName          string             `json:"first_name" bson:"first_name"`
	LastName           string             `json:"last_name" bson:"last_name"`
	PhoneNumber        string             `json:"phone_number" bson:"phone_number"`
	UserType           string             `json:"user_type" bson:"user_type"`
	Login              string             `json:"login" bson:"login"`
	LastLogin          string             `json:"last_login" bson:"last_login"`
	ExtraInfo          string             `json:"extra_info" bson:"extra_info"`
	Address            string             `json:"address" bson:"address"`
	Password           string             `json:"password" bson:"password"`
	PassportNumber     string             `json:"passport_number" bson:"passport_number"`
	PassportIssuePlace string             `json:"passport_issue_place" bson:"passport_issue_place"`
	Email              string             `json:"email" bson:"email"`
	Pinfl              string             `json:"pinfl" bson:"pinfl"`
	Soato              string             `json:"soato" bson:"soato"`
	ExternalID         int32              `json:"external_id" bson:"external_id"`
	Policy             int64              `json:"policy" bson:"policy"`
	Inn                int64              `json:"inn" bson:"inn"`
	City               *City              `json:"city" bson:"city"`
	Region             *Region            `json:"region" bson:"region"`
	PassportIssueDate  time.Time          `json:"passport_issue_date" bson:"passport_issue_date"`
	CreatedAt          time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at" bson:"updated_at"`
	Verified           bool               `json:"verified" bson:"verified"`
	Confirmed          bool               `json:"confirmed" bson:"confirmed"`
	Status             bool               `json:"status" bson:"status"`
	OneIdLogin         string             `json:"one_id_login" bson:"one_id_login"`
	OrganizationInn    string             `json:"organization_inn" bson:"organization_inn"`
}

type CreateUpdateStaffSwag struct {
	MiddleName         string  `json:"middle_name"`
	FirstName          string  `json:"first_name"`
	LastName           string  `json:"last_name"`
	PhoneNumber        string  `json:"phone_number"`
	Login              string  `json:"login"`
	Address            string  `json:"address"`
	Password           string  `json:"password"`
	PassportNumber     string  `json:"passport_number"`
	PassportIssuePlace string  `json:"passport_issue_place"`
	Email              string  `json:"email"`
	ExtraInfo          string  `json:"extra_info"`
	RoleID             string  `json:"role_id"`
	OrganizationID     string  `json:"organization_id"`
	PassportIssueDate  string  `json:"passport_issue_date"`
	Pinfl              string  `json:"pinfl"`
	Inn                int64   `json:"inn"`
	City               *City   `json:"city"`
	Region             *Region `json:"region"`
	Status             bool    `json:"status"`
	OneIdLogin         string  `json:"one_id_login" bson:"one_id_login"`
	OrganizationInn    string  `json:"organization_inn" bson:"organization_inn"`
}

type City struct {
	ID     string `json:"id" bson:"id"`
	Name   string `json:"name" bson:"name"`
	RuName string `json:"ru_name" bson:"ru_name"`
	Soato  uint32 `json:"soato" bson:"soato"`
	Code   uint32 `json:"code" bson:"code"`
}
type Region struct {
	ID     string `json:"id" bson:"id"`
	Name   string `json:"name" bson:"name"`
	RuName string `json:"ru_name" bson:"ru_name"`
	Soato  uint32 `json:"soato" bson:"soato"`
	Code   uint32 `json:"code" bson:"code"`
}

type Discrit struct {
	ID     string `json:"id" bson:"id"`
	Name   string `json:"name" bson:"name"`
	RuName string `json:"ru_name" bson:"ru_name"`
	Soato  uint64 `json:"soato" bson:"soato"`
	Code   uint32 `json:"code" bson:"code"`
}

type CreateUpdateStaffImport struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id"`
	RoleID             primitive.ObjectID `json:"role_id" bson:"role_id"`
	OrganizationID     primitive.ObjectID `json:"organization_id" bson:"organization_id"`
	UniqueName         string             `json:"unique_name" bson:"unique_name"`
	MiddleName         string             `json:"middle_name" bson:"middle_name"`
	FirstName          string             `json:"first_name" bson:"first_name"`
	LastName           string             `json:"last_name" bson:"last_name"`
	PhoneNumber        string             `json:"phone_number" bson:"phone_number"`
	UserType           string             `json:"user_type" bson:"user_type"`
	Login              string             `json:"login" bson:"login"`
	LastLogin          string             `json:"last_login" bson:"last_login"`
	ExtraInfo          string             `json:"extra_info" bson:"extra_info"`
	Address            string             `json:"address" bson:"address"`
	Password           string             `json:"password" bson:"password"`
	PassportNumber     string             `json:"passport_number" bson:"passport_number"`
	PassportIssuePlace string             `json:"passport_issue_place" bson:"passport_issue_place"`
	Email              string             `json:"email" bson:"email"`
	ExternalID         int32              `json:"external_id" bson:"external_id"`
	Policy             int64              `json:"policy" bson:"policy"`
	Inn                int64              `json:"inn" bson:"inn"`
	Pinfl              string             `json:"pinfl" bson:"pinfl"`
	City               *City              `json:"city" bson:"city"`
	Region             *Region            `json:"region" bson:"region"`
	PassportIssueDate  string             `json:"passport_issue_date" bson:"passport_issue_date"`
	CreatedAt          time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at" bson:"updated_at"`
}

type StaffImport struct {
	ID                 string `json:"id" bson:"_id"`
	RoleID             string `json:"role_id" bson:"role_id"`
	OrganizationID     string `json:"organization_id" bson:"organization_id"`
	UniqueName         string `json:"unique_name" bson:"unique_name"`
	MiddleName         string `json:"middle_name" bson:"middle_name"`
	FirstName          string `json:"first_name" bson:"first_name"`
	LastName           string `json:"last_name" bson:"last_name"`
	PhoneNumber        string `json:"phone_number" bson:"phone_number"`
	UserType           string `json:"user_type" bson:"user_type"`
	Login              string `json:"login" bson:"login"`
	LastLogin          string `json:"last_login" bson:"last_login"`
	ExtraInfo          string `json:"extra_info" bson:"extra_info"`
	Address            string `json:"address" bson:"address"`
	Password           string `json:"password" bson:"password"`
	PassportNumber     string `json:"passport_number" bson:"passport_number"`
	PassportIssuePlace string `json:"passport_issue_place" bson:"passport_issue_place"`
	Email              string `json:"email" bson:"email"`
	Pinfl              string `json:"pinfl" bson:"pinfl"`
	ExternalID         int32  `json:"external_id" bson:"external_id"`
	Policy             int64  `json:"policy" bson:"policy"`
	Inn                int64  `json:"inn" bson:"inn"`
	OneIdLogin         string `json:"one_id_login" bson:"one_id_login"`
}

type UpdatePassword struct {
	NewPassword string `json:"new_password" bson:"new_password"`
	OldPassword string `json:"old_password" bson:"old_password"`
}

type GetAllStaffsRequestSwag struct {
	PhoneNumber    string `json:"phone_number"`
	OrganizationID string `json:"organization_id"`
	CityID         string `json:"city_id"`
	SearchString   string `json:"search_string"`
	RoleID         string `json:"role_id"`
	Soato          string `json:"soato"`
	Status         bool   `json:"status"`
	Confirmed      bool   `json:"confirmed"`
	Page           uint32 `json:"page"`
	Limit          uint32 `json:"limit"`
	OneIdLogin     string `json:"one_id_login" bson:"one_id_login"`
}

type EntityActionWithStatus struct {
	OrganizationName string    `json:"organization_name" bson:"organization_name"`
	Conclusion       string    `json:"conclusion" bson:"conclusion"`
	ConclTime        time.Time `json:"concl_time" bson:"concl_time"`
}

type MereAsistReqSwag struct {
	OneIdLogin string   `json:"one_id_login" bson:"one_id_login"`
	CityId     *City    `json:"city" bson:"city"`
	Region     *Region  `json:"region" bson:"region"`
	Discrit    *Discrit `json:"discrit" bson:"discrit"`
	Soato      string   `json:"soato" bson:"soato"`
}

type MereAsistResSwag struct {
	Id         string   `json:"id" bson:"_id"`
	OneIdLogin string   `json:"one_id_login" bson:"one_id_login"`
	City       *City    `json:"city" bson:"city"`
	Region     *Region  `json:"region" bson:"region"`
	Discrit    *Discrit `json:"discrit" bson:"discrit"`
}

type GetMereAsistReqSwag struct {
	DiscritSoato int64  `json:"discrit_soato" bson:"discrit_soato"`
	OneIdLogin   string `json:"one_id_login" bson:"one_id_login"`
	RegionId     string `json:"region_id" bson:"region_id"`
}
