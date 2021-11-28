package workspace

import (
	"context"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/matthieuberger/butler/internal/logger"
	"github.com/matthieuberger/butler/internal/services/gen/auth"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/google/uuid"
	butlerctx "github.com/matthieuberger/butler/internal/context"
	"github.com/matthieuberger/butler/internal/services/gen/workspace"
	"github.com/matthieuberger/butler/internal/services/workspace/models"
	"github.com/matthieuberger/butler/internal/services/workspace/repositories"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// Service has router and db instances
type Service struct {
	OrganizationRepo *repositories.OrganizationRepo
	WorkspaceRepo    *repositories.WorkspaceRepo
	AuthClient       auth.AuthServiceClient
	workspace.UnimplementedWorkspaceServiceServer
}

// NewWorkspaceService initialize with predefined configuration
func NewWorkspaceService(config *ServiceConfig, db *gorm.DB) *Service {
	return &Service{
		OrganizationRepo: repositories.NewOrganizationRepo(db),
		WorkspaceRepo:    repositories.NewWorkspaceRepo(db),
		AuthClient:       connectAuth(config.AuthServiceAddr),
	}
}

func connectAuth(addr string) auth.AuthServiceClient {
	if conn, err := grpc.Dial(addr, []grpc.DialOption{
		grpc.WithInsecure(),
	}...); err != nil {
		logger.Error(context.Background(), "Unable to dial auth service", zap.Error(err))
		return nil
	} else {
		return auth.NewAuthServiceClient(conn)
	}
}

// RegisterGRPC Service to the specified grpc server
func (svc *Service) RegisterGRPC(server *grpc.Server) {
	workspace.RegisterWorkspaceServiceServer(server, svc)
}

func (svc *Service) ListOrganizations(ctx context.Context, req *emptypb.Empty) (*workspace.OrganizationListResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace.ListOrganizations")
	defer span.Finish()

	userID, err := butlerctx.GetCtxTagUserID(ctx)
	if err != nil {
		return &workspace.OrganizationListResponse{}, ErrMissingIndentity
	}
	logger.Debug(ctx, "Got userID ", zap.String("userId", userID))

	orgList, err := svc.OrganizationRepo.ListByUserID(userID)
	logger.Debug(ctx, "Got orglist of size ", zap.Int("orgList size", len(orgList)))
	if err != nil {
		logger.Error(ctx, "Failed to get organization list", zap.Error(err))
		return &workspace.OrganizationListResponse{}, ErrInternal
	}

	result := make([]*workspace.Organization, len(orgList))
	for i, org := range orgList {
		result[i] = org.ToPb()
	}
	logger.Debug(ctx, "About to send result", zap.Int("orgList size", len(result)))

	return &workspace.OrganizationListResponse{
		Organizations: result,
	}, nil
}

func (svc *Service) CreateOrganization(ctx context.Context, req *workspace.CreateOrganizationRequest) (*workspace.OrganizationResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace.CreateOrganization")
	defer span.Finish()

	// Retrieve userID from context
	userID, err := butlerctx.GetCtxTagUserID(ctx)
	if err != nil {
		return &workspace.OrganizationResponse{}, ErrMissingIndentity
	}

	// Default user members
	members := make([]models.OrganizationMember, 1)
	logger.Debug(ctx, "mambers array exists ")
	members = append(members, models.OrganizationMember{
		UserID: uuid.MustParse(userID),
		Role:   "owner",
	})

	logger.Debug(ctx, "Created members ")

	if org, err := svc.OrganizationRepo.CreateOne(&models.Organization{
		Name:        req.Name,
		OwnerID:     uuid.MustParse(userID),
		Workspaces:  []models.Workspace{},
		UserMembers: members,
	}); err != nil {
		logger.Error(ctx, "Unable to create organization", zap.Error(err))
		return &workspace.OrganizationResponse{}, ErrInternal
	} else {
		return &workspace.OrganizationResponse{
			Organization: org.ToPb(),
		}, nil
	}
}

func (svc *Service) CreateWorkspace(ctx context.Context, req *workspace.CreateWorkspaceRequest) (*workspace.WorkspaceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace.CreateWorkspace")
	defer span.Finish()

	// TODO: verify user is admin on ORGANIZATION_ID

	if ws, err := svc.WorkspaceRepo.CreateOne(&models.Workspace{
		Name:           req.Name,
		Description:    req.Description,
		OrganizationID: uuid.MustParse(req.OrganizationID),
	}); err != nil {
		return &workspace.WorkspaceResponse{}, ErrInternal
	} else {
		return &workspace.WorkspaceResponse{
			Workspace: ws.ToPb(),
		}, nil
	}
}

func (svc *Service) AddOrganizationMember(context.Context, *workspace.AddOrganizationMemberRequest) (*workspace.OrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrganizationMember not implemented")
}

func (svc *Service) AddWorkspaceMember(context.Context, *workspace.AddWorkspaceMemberRequest) (*workspace.OrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWorkspaceMember not implemented")
}

func (svc *Service) AuthFuncOverride(ctx context.Context, fullmethodName string) (context.Context, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace.AuthFuncOverride")
	defer span.Finish()

	logger.Debug(ctx, "Checking auth privilleges")

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	if resp, err := svc.AuthClient.IsValidAccessToken(ctx, &auth.IsValidAccessTokenRequest{AccessToken: token}); err != nil {
		logger.Error(ctx, "Failed to validate access token", zap.Error(err))
		return ctx, ErrInvalidGrant
	} else {
		butlerctx.SetCtxTagUserID(ctx, resp.UserID)
		//logger.Debug(ctx, "Got userID ", zap.String("userId", resp.UserID))
		return ctx, nil
	}
}
