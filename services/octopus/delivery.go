package octopus

import (
	"google.golang.org/grpc"
)

type UseCase interface {
	RegisterUseCaseEndpoints(server *grpc.Server)
}

// TODO Remove deprecated code
//func (svc *OctopusService) ConnectWithCode(ctx context.Context, req *octopus.ConnectWithCodeRequest) (*octopus.Connector, error) {
//	span, ctx := opentracing.StartSpanFromContext(ctx, "octopus.ConnectWithCode")
//	defer span.Finish()
//
//	connector, err := svc.ConnectorUsecase.ConnectWithCode(ctx, req.WorkspaceId, req.Provider, req.Code)
//	if err != nil {
//		return &octopus.Connector{}, err
//	}
//
//	return connector.ToPb(), nil
//}

//func (svc *OctopusService) SelectAccount(ctx context.Context, req *octopus.SelectAccountRequest) (*octopus.Connector, error) {
//	span, ctx := opentracing.StartSpanFromContext(ctx, "octopus.SelectAccount")
//	defer span.Finish()
//
//	return svc.ConnectorUsecase.SelectProviderAccount(ctx, req)
//}
//
//func (svc *OctopusService) GetConnectorSecret(ctx context.Context, req *octopus.GetConnectorSecretRequest) (*octopus.ConnectorSecretPair, error) {
//	span, ctx := opentracing.StartSpanFromContext(ctx, "octopus.GetConnectorSecret")
//	defer span.Finish()
//
//	return svc.ConnectorUsecase.GetConnectorSecret(ctx, req)
//}
