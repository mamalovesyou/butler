import {
  GoogleRpcStatus,
  V1CatalogConnectorList,
  V1ConnectWithCodeRequest,
  V1WorkspaceConnector,
  V1WorkspaceConnectorList, V1WorkspaceConnectorsRequest
} from '../../api';

export const LIST_CATALOG_CONNECTORS_REQUEST =
  'LIST_CATALOG_CONNECTORS_REQUEST';
export const LIST_CATALOG_CONNECTORS_SUCCESS =
  'LIST_CATALOG_CONNECTORS_SUCCESS';
export const LIST_CATALOG_CONNECTORS_FAILURE =
  'LIST_CATALOG_CONNECTORS_FAILURE';

export const LIST_WORKSPACE_CONNECTORS_REQUEST =
  'LIST_WORKSPACE_CONNECTORS_REQUEST';
export const LIST_WORKSPACE_CONNECTORS_SUCCESS =
  'LIST_WORKSPACE_CONNECTORS_SUCCESS';
export const LIST_WORKSPACE_CONNECTORS_FAILURE =
  'LIST_WORKSPACE_CONNECTORS_FAILURE';

export const CONNECT_OAUTH_CONNECTOR_REQUEST =
  'CONNECT_OAUTH_CONNECTOR_REQUEST';
export const CONNECT_OAUTH_CONNECTOR_SUCCESS =
  'CONNECT_OAUTH_CONNECTOR_SUCCESS';
export const CONNECT_OAUTH_CONNECTOR_FAILURE =
  'CONNECT_OAUTH_CONNECTOR_FAILURE';

// List Catalog Connectors
export interface IListCatalogConnectorsRequest {
  type: typeof LIST_CATALOG_CONNECTORS_REQUEST;
}

export interface IListCatalogConnectorsSuccess {
  type: typeof LIST_CATALOG_CONNECTORS_SUCCESS;
  payload: V1CatalogConnectorList;
}

export interface IListCatalogConnectorsFailure {
  type: typeof LIST_CATALOG_CONNECTORS_FAILURE;
  error: GoogleRpcStatus;
}

// List Workspace Connectors
export interface IListWorkspaceConnectorsRequest {
  type: typeof LIST_WORKSPACE_CONNECTORS_REQUEST;
  payload: V1WorkspaceConnectorsRequest
}

export interface IListWorkspaceConnectorsSuccess {
  type: typeof LIST_WORKSPACE_CONNECTORS_SUCCESS;
  payload: V1WorkspaceConnectorList;
}

export interface IListWorkspaceConnectorsFailure {
  type: typeof LIST_WORKSPACE_CONNECTORS_FAILURE;
  error: GoogleRpcStatus;
}

// List Workspace Connectors
export interface IConnectOAuthConnectorRequest {
  type: typeof CONNECT_OAUTH_CONNECTOR_REQUEST;
  payload: V1ConnectWithCodeRequest;
}

export interface IConnectOAuthConnectorSuccess {
  type: typeof CONNECT_OAUTH_CONNECTOR_SUCCESS;
  payload: V1WorkspaceConnector;
}

export interface IConnectOAuthConnectorFailure {
  type: typeof CONNECT_OAUTH_CONNECTOR_FAILURE;
  error: GoogleRpcStatus;
}

export type ConnectorsActionType =
  | IListCatalogConnectorsRequest
  | IListCatalogConnectorsSuccess
  | IListCatalogConnectorsFailure
  | IListWorkspaceConnectorsRequest
  | IListWorkspaceConnectorsSuccess
  | IListWorkspaceConnectorsFailure
  | IConnectOAuthConnectorRequest
  | IConnectOAuthConnectorSuccess
  | IConnectOAuthConnectorFailure;
