package var_user_service

// Login struct
type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginInfo struct {
	ID             string `json:"id" bson:"_id"`
	UserType       string `json:"user_type" bson:"user_type"`
	RoleID         string `json:"role_id" bson:"role_id"`
	OrganizationID string `json:"organization_id" bson:"organization_id"`
	FirstName      string `json:"first_name" bson:"first_name"`
	LastName       string `json:"last_name" bson:"last_name"`
	FullName       string `json:"full_name" bson:"full_name"`
	Password       string `json:"password" bson:"password"`
	Login          string `json:"login" bson:"login"`
	Soato          string `json:"soato" bson:"soato"`
	Role           *Role  `json:"role" bson:"role"`
	Verified       bool   `json:"verified" bson:"verified"`
	Confirmed      bool   `json:"confirmed" bson:"confirmed"`
}

type LoginExistsRequest struct {
	Login string `json:"login"`
}
type LoginExistsResponse struct {
	Exist bool `json:"exist"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Verified     bool   `json:"verified"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}
