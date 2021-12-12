package workspace

import (
	"context"
	butlerctx "github.com/butlerhq/butler/internal/context"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/users/models"
	"github.com/butlerhq/butler/services/users/services"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (svc *WorkspaceUsecase) CreateOrganization(ctx context.Context, name string) (*models.Organization, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace_ucase.CreateOrganization")
	defer span.Finish()

	// Retrieve userID from context
	userID, err := butlerctx.GetCtxTagUserID(ctx)
	if err != nil {
		return nil, services.ErrMissingIndentity
	}

	// Default user members
	members := make([]models.OrganizationMember, 1)
	members = append(members, models.OrganizationMember{
		UserID: uuid.MustParse(userID),
		Role:   "owner",
	})

	org, err := svc.OrganizationRepo.CreateOne(&models.Organization{
		Name:        name,
		OwnerID:     uuid.MustParse(userID),
		Workspaces:  []models.Workspace{},
		UserMembers: members,
	})

	if err != nil {
		logger.Error(ctx, "Unable to create organization", zap.Error(err))
		return nil, err
	}

	return org, nil
}
