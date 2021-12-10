package workspace

import (
	"context"
	butlerctx "github.com/butlerhq/butler/butler-core/context"
	"github.com/butlerhq/butler/butler-core/logger"
	"github.com/butlerhq/butler/proto/gen/workspace"
	"github.com/butlerhq/butler/services/users/models"
	"github.com/butlerhq/butler/services/users/services"
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
	members := make([]models.OrganizationMember, 1)
	members = append(members, models.OrganizationMember{
		UserID: uuid.MustParse(userID),
		Role:   "owner",
	})

	org, err := svc.OrganizationRepo.CreateOne(&models.Organization{
		Name:        req.Name,
		OwnerID:     uuid.MustParse(userID),
		Workspaces:  []models.Workspace{},
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
