package connector

//func (uc *ConnectorUsecase) SelectProviderAccount(ctx context.Context, req *octopus.SelectAccountRequest) (*octopus.Connector, error) {
//	span, ctx := opentracing.StartSpanFromContext(ctx, "octopus.SelectProviderAccount")
//	defer span.Finish()
//
//	fmt.Printf("Request payload: %v\n", req)
//	accountConfig := models.ConnectorConfig{
//		ConnectorID: uuid.MustParse(req.ConnectorId),
//		AccountID:   req.AccountId,
//		AccountName: req.AccountName,
//		IsTest:      req.IsTestAccount,
//	}
//
//	wsConnector, err := uc.ConnectorRepo.UpsertConnectorConfig(accountConfig)
//	if err != nil {
//		logger.Error(ctx, "Unable to find a workspace connector", zap.Error(err), zap.Any("request", req))
//		return &octopus.Connector{}, err
//	}
//
//	return wsConnector.ToPb(), err
//}
