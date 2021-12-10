package workspace

import (
	"context"
	"github.com/butlerhq/butler/proto/gen/workspace"
	"github.com/opentracing/opentracing-go"
)

func (svc *WorkspaceService) InviteWorkspaceMember(ctx context.Context, req *workspace.InviteWorkspaceMemberRequest) (*workspace.Invitation, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace.IniviteWorkspaceMember")
	defer span.Finish()

	return &workspace.Invitation{}, nil
}
