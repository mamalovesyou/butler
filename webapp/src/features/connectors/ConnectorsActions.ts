import {
  GoogleRpcStatus,
  V1CatalogConnectorList,
  V1OAuthAuthorizationRequestConnectorCode,
  V1WorkspaceConnector,
  V1WorkspaceConnectorList
} from '../../api';
import * as ActionTypes from './ConnectorsActions.types';

export const listCatalogConnectorsRequest =
  (): ActionTypes.ConnectorsActionType => ({
    type: ActionTypes.LIST_CATALOG_CONNECTORS_REQUEST
  });

export const listCatalogConnectorsSuccess = (
  payload: V1CatalogConnectorList
): ActionTypes.ConnectorsActionType => ({
  type: ActionTypes.LIST_CATALOG_CONNECTORS_SUCCESS,
  payload
});

export const listCatalogConnectorsFailure = (
  error: GoogleRpcStatus
): ActionTypes.ConnectorsActionType => ({
  type: ActionTypes.LIST_CATALOG_CONNECTORS_FAILURE,
  error
});

export const listWorkspaceConnectorsRequest =
  (): ActionTypes.ConnectorsActionType => ({
    type: ActionTypes.LIST_WORKSPACE_CONNECTORS_REQUEST
  });

export const listWorkspaceConnectorsSuccess = (
  payload: V1WorkspaceConnectorList
): ActionTypes.ConnectorsActionType => ({
  type: ActionTypes.LIST_WORKSPACE_CONNECTORS_SUCCESS,
  payload
});

export const listWorkspaceConnectorsFailure = (
  error: GoogleRpcStatus
): ActionTypes.ConnectorsActionType => ({
  type: ActionTypes.LIST_WORKSPACE_CONNECTORS_FAILURE,
  error
});

export const connectOAuthConnectorRequest = (
  payload: V1OAuthAuthorizationRequestConnectorCode
): ActionTypes.ConnectorsActionType => ({
  type: ActionTypes.CONNECT_OAUTH_CONNECTOR_REQUEST,
  payload
});

export const connectOAuthConnectorSuccess = (
  payload: V1WorkspaceConnector
): ActionTypes.ConnectorsActionType => ({
  type: ActionTypes.CONNECT_OAUTH_CONNECTOR_SUCCESS,
  payload
});

export const connectOAuthConnectorFailure = (
  error: GoogleRpcStatus
): ActionTypes.ConnectorsActionType => ({
  type: ActionTypes.CONNECT_OAUTH_CONNECTOR_FAILURE,
  error
});