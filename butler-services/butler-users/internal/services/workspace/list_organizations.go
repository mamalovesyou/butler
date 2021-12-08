package workspace

import (
	"context"
	butlerctx "github.com/butlerhq/butler/butler-core/context"
	"github.com/butlerhq/butler/butler-core/logger"
	"github.com/butlerhq/butler/butler-proto/gen/workspace"
	"github.com/butlerhq/butler/butler-services/butler-users/internal/services"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (svc *WorkspaceService) ListOrganizations(ctx context.Context, req *emptypb.Empty) (*workspace.OrganizationListResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace.ListOrganizations")
	defer span.Finish()

	userID, err := butlerctx.GetCtxTagUserID(ctx)
	if err != nil {
		return &workspace.OrganizationListResponse{}, services.ErrMissingIndentity
	}

	orgList, err := svc.OrganizationRepo.ListByUserID(userID)
	if err != nil {
		logger.Error(ctx, "Failed to get organization list", zap.Error(err))
		return &workspace.OrganizationListResponse{}, services.ErrInternal
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
