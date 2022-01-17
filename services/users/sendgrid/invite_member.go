package sendgrid

import (
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const INVITE_WORKSPACE_TEAM_MEMBER_TEMPLATE_ID = "d-d57ab194439f42f48161d6ee4f0c584a"

func NewInviteWorkspaceTeamMember(email, workspaceName, joinLink string) *mail.SGMailV3 {
	m := mail.NewV3Mail()
	e := mail.NewEmail(USER_NAME, FROM_ADDRESS)
	m.SetFrom(e)

	m.SetTemplateID(INVITE_WORKSPACE_TEAM_MEMBER_TEMPLATE_ID)

	p := mail.NewPersonalization()
	p.AddTos(mail.NewEmail("", email))
	p.SetDynamicTemplateData("workspace_name", workspaceName)
	p.SetDynamicTemplateData("join_link", joinLink)
	m.AddPersonalizations(p)
	return m
}
