package connector

//
//func (uc *ConnectorUsecase) GetConnectorSecret(ctx context.Context, req *octopus.GetConnectorSecretRequest) (*octopus.ConnectorSecretPair, error) {
//	span, ctx := opentracing.StartSpanFromContext(ctx, "connector_ucase.GetConnectorSecret")
//	defer span.Finish()
//
//	if connector, err := uc.ConnectorRepo.FindByWorkspaceAndProvider(
//		uuid.MustParse(req.WorkspaceId),
//		req.Provider,
//	); err != nil {
//		logger.Error(ctx, "Unable to find a workspace connector", zap.Error(err), zap.Any("request", req))
//		return &octopus.ConnectorSecretPair{}, nil
//	} else {
//		return connector.ToConnectorSecretPairPb(), nil
//	}
//}
