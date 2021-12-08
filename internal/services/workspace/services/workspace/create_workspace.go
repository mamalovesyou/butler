package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/butlerhq/butler/internal/services/gen/workspace"
	"github.com/butlerhq/butler/internal/services/workspace/models"
	"github.com/butlerhq/butler/internal/services/workspace/services"
	"github.com/opentracing/opentracing-go"
)

func (svc *WorkspaceService) CreateWorkspace(ctx context.Context, req *workspace.CreateWorkspaceRequest) (*workspace.WorkspaceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace.CreateWorkspace")
	defer span.Finish()

	// TODO: verify user is admin on ORGANIZATION_ID

	ws, err := svc.WorkspaceRepo.CreateOne(&models.Workspace{
		Name:           req.Workspace.Name,
		Description:    req.Workspace.Description,
		OrganizationID: uuid.MustParse(req.OrganizationId),
	});
	
	if err != nil {
		// TODO: Add logger
		return &workspace.WorkspaceResponse{}, services.ErrInternal
	}

	return &workspace.WorkspaceResponse{
		Workspace: ws.ToPb(),
	}, nil
}