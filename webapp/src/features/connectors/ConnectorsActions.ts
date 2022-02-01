import {
  GoogleRpcStatus,
  V1CatalogConnectorList,
  V1ConnectWithCodeRequest, V1ListAccountsRequest, V1ProviderAccount, V1SelectAccountRequest,
  V1WorkspaceConnector,
  V1WorkspaceConnectorList, V1WorkspaceConnectorsRequest
} from '../../api';
import * as ActionTypes from './ConnectorsActions.types';
import {UPDATE_CONNECTOR_CONFIG_FAILURE} from "./ConnectorsActions.types";

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
  (payload: V1WorkspaceConnectorsRequest): ActionTypes.ConnectorsActionType => ({
    type: ActionTypes.LIST_WORKSPACE_CONNECTORS_REQUEST,
    payload
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
  payload: V1ConnectWithCodeRequest
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

export const listConnectorAccountsRequest = (
    payload: V1ListAccountsRequest
): ActionTypes.ConnectorsActionType => ({
  type: ActionTypes.LIST_CONNECTOR_ACCOUNTS_REQUEST,
  payload
});

export const listConnectorAccountsSuccess = (
    payload: V1ProviderAccount[]
): ActionTypes.ConnectorsActionType => ({
  type: ActionTypes.LIST_CONNECTOR_ACCOUNTS_SUCCESS,
  payload
});

export const listConnectorAccountsFailure = (
    error: string
): ActionTypes.ConnectorsActionType => ({
  type: ActionTypes.LIST_CONNECTOR_ACCOUNTS_FAILURE,
  error
});

export const updateConnectorConfigRequest = (
    payload: V1SelectAccountRequest
): ActionTypes.ConnectorsActionType => ({
  type: ActionTypes.UPDATE_CONNECTOR_CONFIG_REQUEST,
  payload
});

export const updateConnectorConfigSuccess = (
    payload: V1WorkspaceConnector
): ActionTypes.ConnectorsActionType => ({
  type: ActionTypes.UPDATE_CONNECTOR_CONFIG_SUCCESS,
  payload
});

export const updateConnectorConfigFailure = (
    error: GoogleRpcStatus
): ActionTypes.ConnectorsActionType => ({
  type: ActionTypes.UPDATE_CONNECTOR_CONFIG_FAILURE,
  error
});

export const setConfigureAccountDialogOpen = (
    open: boolean
): ActionTypes.ConnectorsActionType => ({
  type: ActionTypes.SET_CONFIGURE_ACCOUNT_DIALOG_OPEN,
  payload: open
});
