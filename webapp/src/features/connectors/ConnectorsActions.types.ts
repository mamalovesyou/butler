import {
    GoogleRpcStatus,
    V1CatalogConnectorList,
    V1ConnectWithCodeRequest, V1ListAccountsRequest, V1ProviderAccount, V1SelectAccountRequest,
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

export const LIST_CONNECTOR_ACCOUNTS_REQUEST =
    'LIST_CONNECTOR_ACCOUNTS_REQUEST';
export const LIST_CONNECTOR_ACCOUNTS_SUCCESS =
    'LIST_CONNECTOR_ACCOUNTS_SUCCESS';
export const LIST_CONNECTOR_ACCOUNTS_FAILURE =
    'LIST_CONNECTOR_ACCOUNTS_FAILURE';

export const UPDATE_CONNECTOR_CONFIG_REQUEST =
    'UPDATE_CONNECTOR_CONFIG_REQUEST';
export const UPDATE_CONNECTOR_CONFIG_SUCCESS =
    'UPDATE_CONNECTOR_CONFIG_SUCCESS';
export const UPDATE_CONNECTOR_CONFIG_FAILURE =
    'UPDATE_CONNECTOR_CONFIG_FAILURE';

export const SET_CONFIGURE_ACCOUNT_DIALOG_OPEN =
    'SET_CONFIGURE_ACCOUNT_DIALOG_OPEN';

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

// Update connector config
export interface IUpdateConnectorConfigRequest {
    type: typeof UPDATE_CONNECTOR_CONFIG_REQUEST;
    payload: V1SelectAccountRequest;
}

export interface IUpdateConnectorConfigSuccess {
    type: typeof UPDATE_CONNECTOR_CONFIG_SUCCESS;
    payload: V1WorkspaceConnector;
}

export interface IUpdateConnectorConfigFailure {
    type: typeof UPDATE_CONNECTOR_CONFIG_FAILURE;
    error: GoogleRpcStatus;
}

export interface IListConnectorAccountsRequest {
    type: typeof LIST_CONNECTOR_ACCOUNTS_REQUEST;
    payload: V1ListAccountsRequest;
}

export interface IListConnectorAccountsSuccess {
    type: typeof LIST_CONNECTOR_ACCOUNTS_SUCCESS;
    payload: V1ProviderAccount[];
}

export interface IListConnectorAccountsFailure {
    type: typeof LIST_CONNECTOR_ACCOUNTS_FAILURE;
    error: string;
}

export interface ISetConfigureDialogOpen {
    type: typeof SET_CONFIGURE_ACCOUNT_DIALOG_OPEN;
    payload: boolean;
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
    | IConnectOAuthConnectorFailure
    | IListConnectorAccountsRequest
    | IListConnectorAccountsSuccess
    | IListConnectorAccountsFailure
    | IUpdateConnectorConfigRequest
    | IUpdateConnectorConfigSuccess
    | IUpdateConnectorConfigFailure
    | ISetConfigureDialogOpen;
