import {
  GoogleRpcStatus,
  V1ConnectWithCodeRequest,
  V1ConnectorList,
  V1ListConnectorsRequest, V1Connector
} from '../../api';
import * as ActionTypes from './ConnectorsActions.types';

export const listWorkspaceConnectorsRequest =
  (payload: V1ListConnectorsRequest): ActionTypes.ConnectorsActionType => ({
    type: ActionTypes.LIST_WORKSPACE_CONNECTORS_REQUEST,
    payload
  });

export const listWorkspaceConnectorsSuccess = (
  payload: V1ConnectorList
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
  payload: V1Connector
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

export const getConnectorRequest =
(payload: V1ListConnectorsRequest): ActionTypes.ConnectorsActionType => ({
  type: ActionTypes.LIST_WORKSPACE_CONNECTORS_REQUEST,
  payload
});

export const getConnectorSuccess = (
    payload: V1ConnectorList
): ActionTypes.ConnectorsActionType => ({
  type: ActionTypes.LIST_WORKSPACE_CONNECTORS_SUCCESS,
  payload
});

export const getConnectorFailure = (
    error: GoogleRpcStatus
): ActionTypes.ConnectorsActionType => ({
  type: ActionTypes.LIST_WORKSPACE_CONNECTORS_FAILURE,
  error
});
