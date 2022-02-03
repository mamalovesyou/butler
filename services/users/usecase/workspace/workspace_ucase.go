package workspace

import (
	"fmt"

	"github.com/butlerhq/butler/services/users/repositories"
	"github.com/butlerhq/butler/services/users/sendgrid"
	"gorm.io/gorm"
)

type WorkspaceUsecase struct {
	OrganizationRepo *repositories.OrganizationRepo
	WorkspaceRepo    *repositories.WorkspaceRepo
	InvitationRepo   *repositories.InvitationRepo
	EmailClient      *sendgrid.EmailClient
	DB               *gorm.DB
	WebAppBaseURL    string
	AirbyteServerURL string
}

func NewWorkspaceUsecase(db *gorm.DB, sendgridAPIKey, webAppBaseURL, airbyteServerUrl string) *WorkspaceUsecase {
	return &WorkspaceUsecase{
		OrganizationRepo: repositories.NewOrganizationRepo(db),
		WorkspaceRepo:    repositories.NewWorkspaceRepo(db),
		InvitationRepo:   repositories.NewInvitationRepo(db),
		EmailClient:      sendgrid.NewEmailClient(sendgridAPIKey),
		DB:               db,
		WebAppBaseURL:    webAppBaseURL,
		AirbyteServerURL: airbyteServerUrl,
	}
}

func (uc *WorkspaceUsecase) GetJoinLink(invitationID string, token string) string {
	return fmt.Sprintf("%s/register?invitationId=%s&token=%s", uc.WebAppBaseURL, invitationID, token)
}
