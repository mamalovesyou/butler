package context

import (
	"context"
	"fmt"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
)

var (
	contextKeyIdentity = contextKey("identity")
	contextKeyUserID   = contextKey("userID")
	ctxTagUserID       = "user.userid"
)

type contextKey string

func (c contextKey) String() string {
	return fmt.Sprintf("butler:%s", string(c))
}

func WithIdentity(ctx context.Context, idt *Identity) context.Context {
	return context.WithValue(ctx, contextKeyIdentity, idt)
}

func IdentityFromContext(ctx context.Context) (*Identity, error) {
	if identity, ok := ctx.Value(contextKeyIdentity).(*Identity); !ok {
		return nil, ErrIdentityNotFound
	} else {
		return identity, nil
	}
}

func SetCtxTagUserID(ctx context.Context, userID string) grpc_ctxtags.Tags {
	return grpc_ctxtags.Extract(ctx).Set(ctxTagUserID, userID)
}

func GetCtxTagUserID(ctx context.Context) (string, error) {
	values := grpc_ctxtags.Extract(ctx).Values()
	if userID, ok := values[ctxTagUserID]; ok {
		return userID.(string), nil
	} else {
		return "", ErrUserIDNotFound
	}
}

func WithUserID(ctx context.Context, userID string) context.Context {
	//SetCtxTagUserID(ctx, userID)
	return context.WithValue(ctx, contextKeyUserID, userID)
}

func UserIDFromContext(ctx context.Context) (string, error) {
	if userID, ok := ctx.Value(contextKeyUserID).(string); !ok {
		return "", ErrUserIDNotFound
	} else {
		return userID, nil
	}
}
