import {
    GoogleRpcStatus,
    V1ConnectWithCodeRequest,
    V1WorkspaceConnector,
    V1WorkspaceConnectorList, V1WorkspaceConnectorsRequest
} from '../../api';

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

export const GET_CONNECTOR_REQUEST =
    'GET_CONNECTOR_REQUEST';
export const GET_CONNECTOR_SUCCESS =
    'GET_CONNECTOR_SUCCESS';
export const GET_CONNECTOR_FAILURE =
    'GET_CONNECTOR_FAILURE';


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
    | IListWorkspaceConnectorsRequest
    | IListWorkspaceConnectorsSuccess
    | IListWorkspaceConnectorsFailure
    | IConnectOAuthConnectorRequest
    | IConnectOAuthConnectorSuccess
    | IConnectOAuthConnectorFailure
