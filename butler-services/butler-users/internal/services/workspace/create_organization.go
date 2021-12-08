package workspace

import (
	"context"
	butlerctx "github.com/butlerhq/butler/butler-core/context"
	"github.com/butlerhq/butler/butler-core/logger"
	"github.com/butlerhq/butler/butler-proto/gen/workspace"
	models2 "github.com/butlerhq/butler/butler-services/butler-users/internal/models"
	"github.com/butlerhq/butler/butler-services/butler-users/internal/services"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (svc *WorkspaceService) CreateOrganization(ctx context.Context, req *workspace.CreateOrganizationRequest) (*workspace.OrganizationResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace.CreateOrganization")
	defer span.Finish()

	// Retrieve userID from context
	userID, err := butlerctx.GetCtxTagUserID(ctx)
	if err != nil {
		return &workspace.OrganizationResponse{}, services.ErrMissingIndentity
	}

	// Default user members
	members := make([]models2.OrganizationMember, 1)
	members = append(members, models2.OrganizationMember{
		UserID: uuid.MustParse(userID),
		Role:   "owner",
	})

	org, err := svc.OrganizationRepo.CreateOne(&models2.Organization{
		Name:        req.Name,
		OwnerID:     uuid.MustParse(userID),
		Workspaces:  []models2.Workspace{},
		UserMembers: members,
	})
	if err != nil {
		logger.Error(ctx, "Unable to create organization", zap.Error(err))
		return &workspace.OrganizationResponse{}, services.ErrInternal
	}

	return &workspace.OrganizationResponse{
		Organization: org.ToPb(),
	}, nil
}