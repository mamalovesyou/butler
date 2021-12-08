package workspace

import (
	"context"
	"github.com/butlerhq/butler/internal/services/gen/workspace"
	"github.com/opentracing/opentracing-go"
)

func (svc *WorkspaceService) InviteOrganizationMember(ctx context.Context, req *workspace.IniviteOrganizationMemberRequest) (*workspace.Invitation, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace.InviteOrganizationMember")
	defer span.Finish()

	return &workspace.Invitation{}, nil
}