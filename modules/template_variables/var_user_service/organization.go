package var_user_service

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResponsibleUser struct {
	OrganizationId string `json:"organization_id" bson:"organization_id"`
	Soato          string `json:"soato" bson:"soato"`
}

type Organization struct {
	ID              string             `json:"id" bson:"_id"`
	Name            string             `json:"name" bson:"name"`
	FullName        string             `json:"full_name" bson:"full_name"`
	Description     string             `json:"description" bson:"description"`
	Status          bool               `json:"status" bson:"status"`
	ExternalID      int32              `json:"external_id" bson:"external_id"`
	Code            int32              `json:"code" bson:"code"`
	CreatedAt       primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt       primitive.DateTime `json:"updated_at" bson:"updated_at"`
	INN             string             `json:"inn" bson:"inn"`
	Type            string             `json:"type" bson:"type"`
	Soato           int32              `json:"soato" bson:"soato"`
	ResponsibleOrgs []*ResponsibleUser `json:"responsible_orgs" bson:"responsible_orgs"`
}

type CreateUpdateOrganization struct {
	ID              primitive.ObjectID `bson:"_id"`
	Name            string             `bson:"name"`
	FullName        string             `bson:"full_name"`
	Description     string             `bson:"description"`
	Status          bool               `bson:"status"`
	Code            int32              `bson:"code"`
	ExternalID      int32              `bson:"external_id"`
	CreatedAt       time.Time          `bson:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at"`
	INN             string             `json:"inn" bson:"inn"`
	Type            string             `json:"type" bson:"type"`
	ResponsibleOrgs []*ResponsibleUser `json:"responsible_orgs" bson:"responsible_orgs"`
}

type CreateUpdateOrganizationSwag struct {
	Id              string             `json:"id"`
	Name            string             `json:"name"`
	FullName        string             `json:"full_name"`
	Description     string             `json:"description"`
	Status          bool               `json:"status"`
	Code            string             `json:"code"`
	ExternalId      int32              `json:"external_id"`
	Type            string             `json:"type"`
	Inn             string             `json:"inn"`
	Soato           int32              `json:"soato"`
	ResponsibleOrgs []*ResponsibleUser `json:"responsible_orgs"`
}

type GetAllOrganizationsResponse struct {
	Count        uint32          `json:"count"`
	Organization []*Organization `json:"organizations"`
}

type SoliqSooguDetail struct {
	NameUzLatin string `json:"name_uz_latn" bson:"name_uz_latn"`
}

type Company struct {
	Name        string           `json:"name" bson:"name"`
	Description SoliqSooguDetail `json:"sooguDetail" bson:"sooguDetail"`
	Soato       int32            `json:"soato" bson:"soato"`
}

type SoliqResponse struct {
	Company Company `json:"company" bson:"company"`
}
