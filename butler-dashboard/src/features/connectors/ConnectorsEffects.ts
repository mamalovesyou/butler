import { takeEvery, put, fork, select } from 'redux-saga/effects';
import { AxiosResponse } from 'axios';
import * as ActionTypes from './ConnectorsActions.types';
import * as Actions from './ConnectorsActions';
import {
  GoogleRpcStatus,
  V1CatalogConnectorList,
  V1WorkspaceConnector,
  V1WorkspaceConnectorList
} from '../../api';
import { Api } from '../configureEffects';
import { useWorkspace } from '../../hooks/use-workspace';

export function* onListCatalogConnectorsRequest() {
  yield takeEvery(ActionTypes.LIST_CATALOG_CONNECTORS_REQUEST, function* () {
    try {
      const { workspace } = yield select();
      const response: AxiosResponse<V1CatalogConnectorList> = yield Api.v1.connectorsServiceListCatalogConnectors(workspace.workspaceId);
      yield put(Actions.listCatalogConnectorsSuccess(response.data));
    } catch (error) {
      console.log(error);
      const rpcError: GoogleRpcStatus = error?.response?.data;
      yield put(Actions.listCatalogConnectorsFailure(rpcError));
    }
  });
}
export function* onListWorkspaceConnectorsRequest() {
  yield takeEvery(ActionTypes.LIST_WORKSPACE_CONNECTORS_REQUEST, function* () {
    try {
      const { workspace } = yield select();
      const response: AxiosResponse<V1WorkspaceConnectorList> = yield Api.v1.connectorsServiceListWorkspaceConnectors(workspace.workspaceId);
      yield put(Actions.listWorkspaceConnectorsSuccess(response.data));
    } catch (error) {
      console.log(error);
      const rpcError: GoogleRpcStatus = error?.response?.data;
      yield put(Actions.listWorkspaceConnectorsFailure(rpcError));
    }
  });
}

export function* onConnectOAuthConnectorRequest() {
  yield takeEvery(
    ActionTypes.CONNECT_OAUTH_CONNECTOR_REQUEST,
    function* ({ payload }: ActionTypes.IConnectOAuthConnectorRequest) {
      try {
        const { workspace } = yield select();
        const response: AxiosResponse<V1WorkspaceConnector> = yield Api.v1.connectorsServiceGetOauthConnectorAuthorization(
          workspace.workspaceId,
            payload
          );
        yield put(Actions.connectOAuthConnectorSuccess(response.data));
      } catch (error) {
        console.log(error);
      const rpcError: GoogleRpcStatus = error?.response?.data;
        yield put(Actions.connectOAuthConnectorFailure(rpcError));
      }
    }
  );
}

export const connectorEffects = [
  fork(onListCatalogConnectorsRequest),
  fork(onListWorkspaceConnectorsRequest),
  fork(onConnectOAuthConnectorRequest)
];

export default connectorEffects;
