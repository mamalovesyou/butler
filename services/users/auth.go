package users

import (
	"context"
	butlerctx "github.com/butlerhq/butler/internal/context"
	"github.com/butlerhq/butler/services/users/errors"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/opentracing/opentracing-go"
)

func (svc *UsersService) AuthFuncOverride(ctx context.Context, fullmethodName string) (context.Context, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "users.AuthFuncOverride")
	defer span.Finish()

	// Skip authentication when user sign in or register
	switch fullmethodName {
	case "SignIn", "SignUp":
		return ctx, nil
	}

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return ctx, errors.ErrMissingAccessToken
	}

	user, err := svc.PermissionUseCase.IsAuthenticated(ctx, token)
	if err != nil {
		return ctx, err
	}

	butlerctx.SetCtxTagUserID(ctx, user.ID.String())
	return ctx, nil
}
