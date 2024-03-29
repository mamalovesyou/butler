package users

import (
	"context"

	"github.com/butlerhq/butler/api/services/users/v1"
	butlerctx "github.com/butlerhq/butler/internal/context"
	"github.com/butlerhq/butler/internal/errors"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

// SignIn users a user with an email/passsword combination
func (svc *UsersService) SignIn(ctx context.Context, req *users.SignInRequest) (*users.AuthenticatedUser, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.SignIn")
	defer span.Finish()

	authUser, err := svc.UserUsecase.SignIn(ctx, req.Email, req.Password)
	if err != nil {
		return &users.AuthenticatedUser{}, err
	}

	return authUser.ToPb(), nil
}

// SignUp creates a new user
func (svc *UsersService) SignUp(ctx context.Context, req *users.SignUpRequest) (*users.AuthenticatedUser, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.SignUp")
	defer span.Finish()

	authUser, err := svc.UserUsecase.SignUp(ctx, req.FirstName, req.LastName, req.Email, req.Password)
	if err != nil {
		return &users.AuthenticatedUser{}, err
	}

	return authUser.ToPb(), nil
}

// SignUpWithInvite creates a new user
func (svc *UsersService) SignUpWithInvite(ctx context.Context, req *users.SignUpWithInvitationRequest) (*users.AuthenticatedUser, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.SignUpWithInvite")
	defer span.Finish()

	authUser, err := svc.UserUsecase.SignUpWithInvitation(ctx, req.FirstName, req.LastName, req.Password, req.InvitationId, req.Token)
	if err != nil {
		return &users.AuthenticatedUser{}, err
	}

	return authUser.ToPb(), nil
}

func (svc *UsersService) RefreshToken(ctx context.Context, req *users.RefreshRequest) (*users.AuthenticatedUser, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.RefreshToken")
	defer span.Finish()

	authUser, err := svc.UserUsecase.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return &users.AuthenticatedUser{}, err
	}

	return authUser.ToPb(), nil
}

func (svc *UsersService) CreateOrganization(ctx context.Context, req *users.CreateOrganizationRequest) (*users.OrganizationResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.CreateOrganization")
	defer span.Finish()

	org, err := svc.WorkspaceUsecase.CreateOrganization(ctx, req.Name)
	if err != nil {
		return &users.OrganizationResponse{}, err
	}

	return &users.OrganizationResponse{
		Organization: org.ToPb(),
	}, nil
}

func (svc *UsersService) GetOrganization(ctx context.Context, req *users.GetOrganizationRequest) (*users.OrganizationResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.GetOrganization")
	defer span.Finish()

	org, err := svc.WorkspaceUsecase.GetOrganization(ctx, req.OrganizationId)
	if err != nil {
		return &users.OrganizationResponse{}, err
	}

	return &users.OrganizationResponse{
		Organization: org.ToPb(),
	}, nil
}

func (svc *UsersService) CreateWorkspace(ctx context.Context, req *users.CreateWorkspaceRequest) (*users.WorkspaceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.CreateWorkspaceDialog")
	defer span.Finish()

	organizationID := uuid.MustParse(req.OrganizationId)
	ws, err := svc.WorkspaceUsecase.CreateWorkspace(ctx, organizationID, req.Workspace.Name, req.Workspace.Description)
	if err != nil {
		return &users.WorkspaceResponse{}, err
	}

	return &users.WorkspaceResponse{
		Workspace: ws.ToPb(),
	}, nil
}

func (svc *UsersService) GetWorkspace(ctx context.Context, req *users.GetWorkspaceRequest) (*users.WorkspaceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.CetWorkspace")
	defer span.Finish()

	ws, err := svc.WorkspaceUsecase.GetWorkspace(ctx, req.WorkspaceId)
	if err != nil {
		return &users.WorkspaceResponse{}, err
	}

	return &users.WorkspaceResponse{
		Workspace: ws.ToPb(),
	}, nil
}

func (svc *UsersService) GetInvitation(ctx context.Context, req *users.GetInvitationRequest) (*users.Invitation, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.GetInvitation")
	defer span.Finish()

	invitation, err := svc.WorkspaceUsecase.GetInvitation(ctx, req.InvitationId, req.Token)
	if err != nil {
		return &users.Invitation{}, err
	}

	return invitation.ToPb(), nil
}

func (svc *UsersService) CompleteOnboarding(ctx context.Context, req *users.CompleteOnboardingRequest) (*users.OrganizationResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.CompleteOnboarding")
	defer span.Finish()

	org, err := svc.WorkspaceUsecase.CompleteOnboarding(ctx, req.OrganizationId)
	if err != nil {
		return &users.OrganizationResponse{}, err
	}

	return &users.OrganizationResponse{
		Organization: org.ToPb(),
	}, nil
}

func (svc *UsersService) ListOrganizations(ctx context.Context, req *emptypb.Empty) (*users.OrganizationListResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.ListOrganizations")
	defer span.Finish()

	userID, err := butlerctx.GetCtxTagUserID(ctx)
	if err != nil {
		logger.Error(ctx, "Missing userID in context", zap.Error(err))
		return nil, errors.ErrInternal
	}

	orgList, err := svc.WorkspaceUsecase.ListOrganizations(ctx, userID)
	if err != nil {
		return &users.OrganizationListResponse{}, err
	}

	result := make([]*users.Organization, len(orgList))
	for i, org := range orgList {
		result[i] = org.ToPb()
	}

	return &users.OrganizationListResponse{
		Organizations: result,
	}, nil
}

func (svc *UsersService) SendBatchInvitations(ctx context.Context, req *users.BatchInviteMemberRequest) (*users.InvitationListResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.BatchInviteMember")
	defer span.Finish()

	invites, err := svc.WorkspaceUsecase.BatchInviteTeamMembers(ctx, req.OrganizationId, req.WorkspaceId, req.Emails)
	if err != nil {
		return &users.InvitationListResponse{}, err
	}

	result := make([]*users.Invitation, len(invites))
	for i, invitation := range invites {
		result[i] = invitation.ToPb()
	}

	return &users.InvitationListResponse{
		Invitations: result,
	}, nil
}
