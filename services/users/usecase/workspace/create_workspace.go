package workspace

import (
	"context"

	"gorm.io/gorm"

	"github.com/butlerhq/airbyte-client-go/airbyte"
	"github.com/butlerhq/butler/internal/errors"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/users/models"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// CreateWorkspace create a butler workspace and create an associated airbyte workspace
func (svc *WorkspaceUsecase) CreateWorkspace(ctx context.Context, organizationID uuid.UUID, name, description string) (*models.Workspace, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace_ucase.CreateWorkspaceDialog")
	defer span.Finish()

	// TODO: verify user is admin on ORGANIZATION_ID or have permission to create workspace

	// Create airbyte client
	client, err := airbyte.NewClientWithResponses(svc.AirbyteServerURL)
	if err != nil {
		logger.Error(ctx, "Unable to create airbyte client", zap.Error(err))
		return nil, errors.ErrInternal
	}

	butlerWorkspace := models.Workspace{
		Name:           name,
		Description:    description,
		OrganizationID: organizationID,
	}

	err = svc.DB.Transaction(func(tx *gorm.DB) error {
		body := airbyte.CreateWorkspaceJSONRequestBody{
			AnonymousDataCollection: airbyte.PtrBool(true),
			DisplaySetupWizard:      airbyte.PtrBool(false),
			Name:                    name,
			News:                    airbyte.PtrBool(false),
			SecurityUpdates:         airbyte.PtrBool(false),
		}
		wsResponse, err := client.CreateWorkspaceWithResponse(ctx, body)
		if err != nil {
			logger.Error(ctx, "Unable to create airbyte workspace", zap.Error(err))
			return err
		}
		butlerWorkspace.AirbyteWorkspaceID = wsResponse.JSON200.WorkspaceId

		// Create workspace
		if err = tx.Create(&butlerWorkspace).Error; err != nil {
			logger.Error(ctx, "Unable to create butler workspace", zap.Error(err))
			return err
		}

		// Create s3 destination
		destination, ok := svc.Catalog.GetByName("S3")
		if !ok {
			logger.Error(ctx, "Unable to get S3 destination from catalog")
			return errors.ErrInternal
		}

		s3ConnCfg := destination.BuildConfig(butlerWorkspace.ID.String())
		airbyteClient, err := airbyte.NewClientWithResponses(svc.AirbyteServerURL)
		if err != nil {
			logger.Error(ctx, "Unable to create airbyte client", zap.Error(err))
			return errors.ErrInternal
		}
		createS3Body := airbyte.CreateDestinationJSONRequestBody{
			ConnectionConfiguration: s3ConnCfg,
			DestinationDefinitionId: destination.AirbyteDefinitionID(),
			Name:                    "Butler S3 Airbyte Data",
			WorkspaceId:             butlerWorkspace.AirbyteWorkspaceID,
		}
		resp, err := airbyteClient.CreateDestinationWithResponse(ctx, createS3Body)
		if err != nil || resp.JSON200 == nil {
			logger.Error(ctx, "Unable to create airbyte s3 destination", zap.Error(err), zap.Any("airbyteErr", resp.JSON422))
			return errors.ErrInternal
		}

		// Update workspace
		butlerWorkspace.AirbyteDestinationID = resp.JSON200.DestinationId
		if err = tx.Updates(&butlerWorkspace).Error; err != nil {
			logger.Error(ctx, "Unable to update butler workspace", zap.Error(err))
			return err
		}
		return nil
	})

	if err != nil {
		return nil, errors.ErrInternal
	}

	return &butlerWorkspace, nil
}
