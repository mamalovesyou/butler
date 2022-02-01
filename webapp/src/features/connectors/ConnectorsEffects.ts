import {takeEvery, put, fork, select} from 'redux-saga/effects';
import {AxiosResponse} from 'axios';
import * as ActionTypes from './ConnectorsActions.types';
import * as Actions from './ConnectorsActions';
import {
    GoogleRpcStatus,
    V1CatalogConnectorList, V1ListAccountsResponse,
    V1WorkspaceConnector,
    V1WorkspaceConnectorList
} from '../../api';
import {Api} from '../configureEffects';

export function* onListCatalogConnectorsRequest() {
    yield takeEvery(ActionTypes.LIST_CATALOG_CONNECTORS_REQUEST, function* () {
        try {
            const response: AxiosResponse<V1CatalogConnectorList> = yield Api.v1.octopusServiceGetCatalogConnectors();
            yield put(Actions.listCatalogConnectorsSuccess(response.data));
        } catch (error) {
            console.log("err:", error);
            const rpcError: GoogleRpcStatus = error?.response?.data;
            yield put(Actions.listCatalogConnectorsFailure(rpcError));
        }
    });
}

export function* onListWorkspaceConnectorsRequest() {
    yield takeEvery(ActionTypes.LIST_WORKSPACE_CONNECTORS_REQUEST, function* ({payload}: ActionTypes.IListWorkspaceConnectorsRequest) {
        try {
            const response: AxiosResponse<V1WorkspaceConnectorList> = yield Api.v1.octopusServiceListWorkspaceConnectors(payload);
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
        function* ({payload}: ActionTypes.IConnectOAuthConnectorRequest) {
            try {
                const response: AxiosResponse<V1WorkspaceConnector> = yield Api.v1.octopusServiceConnectWithCode(payload);
                yield put(Actions.connectOAuthConnectorSuccess(response.data));
            } catch (error) {
                console.log(error);
                const rpcError: GoogleRpcStatus = error?.response?.data;
                yield put(Actions.connectOAuthConnectorFailure(rpcError));
            }
        }
    );
}

export function* onUpdateConnectorConfigRequest() {
    yield takeEvery(
        ActionTypes.UPDATE_CONNECTOR_CONFIG_REQUEST,
        function* ({payload}: ActionTypes.IUpdateConnectorConfigRequest) {
            try {
                const response: AxiosResponse<V1WorkspaceConnector> = yield Api.v1.octopusServiceSelectAccount(payload);
                yield put(Actions.updateConnectorConfigSuccess(response.data));
            } catch (error) {
                console.log(error);
                const rpcError: GoogleRpcStatus = error?.response?.data;
                yield put(Actions.updateConnectorConfigFailure(rpcError));
            }
        }
    );
}

export function* onListConnectorAccountsRequest() {
    yield takeEvery(
        ActionTypes.LIST_CONNECTOR_ACCOUNTS_REQUEST,
        function* ({payload}: ActionTypes.IListWorkspaceConnectorsRequest) {
            try {
                const response: AxiosResponse<V1ListAccountsResponse> = yield Api.v1.octopusServiceListAccounts(payload);
                yield put(Actions.listConnectorAccountsSuccess(response.data.accounts));
            } catch (error) {
                yield put(Actions.listConnectorAccountsFailure(String(error)));
            }
        }
    );
}

export const connectorEffects = [
    fork(onListCatalogConnectorsRequest),
    fork(onListWorkspaceConnectorsRequest),
    fork(onConnectOAuthConnectorRequest),
    fork(onUpdateConnectorConfigRequest),
    fork(onListConnectorAccountsRequest)
];

export default connectorEffects;
