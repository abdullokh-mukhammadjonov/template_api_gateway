package services

import (
	"fmt"

	"github.com/abdullokh-mukhammadjonov/template_api_gateway/config"
	cs "github.com/abdullokh-mukhammadjonov/template_api_gateway/genproto/content_service"
	us "github.com/abdullokh-mukhammadjonov/template_api_gateway/genproto/user_service"
	"google.golang.org/grpc"
)

type ServiceManager interface {
	HandbookService() cs.HandbooksServiceClient
	StaffService() us.StaffServiceClient
	RoleService() us.RoleServiceClient
	OrganizationService() us.OrganizationServiceClient
}

type grpcClients struct {
	handbookService     cs.HandbooksServiceClient
	staffService        us.StaffServiceClient
	roleService         us.RoleServiceClient
	organizationService us.OrganizationServiceClient
}

// content_service grpclient methods. to satisfy interface methods
// please if you add another function to content_service on proto add
// below not to other section.
// follow the order.

/*                          CONTENT SERVICE                         */
// handbook service
func (g grpcClients) HandbookService() cs.HandbooksServiceClient {
	return g.handbookService
}

/*                          USER SERVICE                         */
// staff service
func (g grpcClients) StaffService() us.StaffServiceClient {
	return g.staffService
}

// role service
func (g grpcClients) RoleService() us.RoleServiceClient {
	return g.roleService
}

// organization service
func (g grpcClients) OrganizationService() us.OrganizationServiceClient {
	return g.organizationService
}

func NewGrpcClients(cfg config.Config) (ServiceManager, error) {
	connContentService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(10*1024*1024)))
	if err != nil {
		return nil, err
	}

	connUserService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(10*1024*1024)))
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		// ** discussion_logic_service
		// ** follow the rule, do not mess around

		/*                          CONTENT SERVICE                         */
		handbookService: cs.NewHandbooksServiceClient(connContentService),

		/*                           USER SERVICE                           */
		staffService:        us.NewStaffServiceClient(connUserService),
		roleService:         us.NewRoleServiceClient(connUserService),
		organizationService: us.NewOrganizationServiceClient(connUserService),
	}, nil
}
