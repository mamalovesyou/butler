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
		if err = tx.Create(&butlerWorkspace).Error; err != nil {
			logger.Error(ctx, "Unable to create butler workspace", zap.Error(err))
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return nil, errors.ErrInternal
	}

	return &butlerWorkspace, nil
}
