package var_user_service

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	ID           string             `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Description  string             `json:"description" bson:"description"`
	Code         int32              `json:"code" bson:"code"`
	Status       bool               `json:"status" bson:"status"`
	Organization *Organization      `json:"organization" bson:"organization"`
	Permissions  []*Permission      `json:"permissions" bson:"permissions"`
	CreatedAt    primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt    primitive.DateTime `json:"updated_at" bson:"updated_at"`
}
type GetRole struct {
	ID             string             `json:"id" bson:"_id"`
	OrganizationID string             `json:"organization_id" bson:"organization_id"`
	Name           string             `json:"name" bson:"name"`
	Description    string             `json:"description" bson:"description"`
	Code           int32              `json:"code" bson:"code"`
	Status         bool               `json:"status" bson:"status"`
	Permissions    []string           `json:"permissions" bson:"permissions"`
	CreatedAt      primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt      primitive.DateTime `json:"updated_at" bson:"updated_at"`
}
type Permission struct {
	ID          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	RuName      string `json:"ru_name" bson:"ru_name"`
	Description string `json:"description" bson:"description"`
	Label       string `json:"label" bson:"label"`
}

type GetRoleOrganization struct {
	OrganizationID string `bson:"organization_id"`
}

type CreateUpdateRole struct {
	ID             primitive.ObjectID   `bson:"_id"`
	OrganizationID primitive.ObjectID   `bson:"organization_id"`
	Name           string               `bson:"name"`
	Description    string               `bson:"description"`
	Code           int32                `bson:"code"`
	ExternalID     int32                `bson:"external_id"`
	Status         bool                 `bson:"status"`
	Permissions    []primitive.ObjectID `bson:"permissions"`
	CreatedAt      time.Time            `bson:"created_at"`
	UpdatedAt      time.Time            `bson:"updated_at"`
}

type CreateUpdateRoleSwag struct {
	Name           string   `json:"name" bson:"name" binding:"required"`
	OrganizationID string   `json:"organization_id" bson:"organization_id"`
	Description    string   `json:"description" bson:"description"`
	Code           int32    `json:"code" bson:"code"`
	Status         bool     `json:"status" bson:"status"`
	Permissions    []string `json:"permissions" bson:"permissions"`
}

type CreatePermission struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	RuName      string             `json:"ru_name" bson:"ru_name"`
	Description string             `json:"description" bson:"description"`
	Label       string             `json:"label" bson:"label"`
}

type CreateUpdatePermissionSwag struct {
	Name        string `json:"name" bson:"name" binding:"required"`
	RuName      string `json:"ru_name" bson:"ru_name" binding:"required"`
	Description string `json:"description" bson:"description"`
	Label       string `json:"label" bson:"label"`
}
