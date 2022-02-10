import {takeEvery, put, fork} from 'redux-saga/effects';
import {AxiosResponse} from 'axios';
import * as ActionTypes from './ConnectorsActions.types';
import * as Actions from './ConnectorsActions';
import {
    GoogleRpcStatus, V1ConnectorList,
    V1WorkspaceConnector,
    V1WorkspaceConnectorList
} from '../../api';
import {Api} from '../configureEffects';


export function* onListWorkspaceConnectorsRequest() {
    yield takeEvery(ActionTypes.LIST_WORKSPACE_CONNECTORS_REQUEST, function* ({payload}: ActionTypes.IListWorkspaceConnectorsRequest) {
        try {
            const response: AxiosResponse<V1ConnectorList> = yield Api.v1.connectorsServiceListConnectors(payload);
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
//
// export function* onUpdateConnectorConfigRequest() {
//     yield takeEvery(
//         ActionTypes.UPDATE_CONNECTOR_CONFIG_REQUEST,
//         function* ({payload}: ActionTypes.IUpdateConnectorConfigRequest) {
//             try {
//                 const response: AxiosResponse<V1WorkspaceConnector> = yield Api.v1.octopusServiceSelectAccount(payload);
//                 yield put(Actions.updateConnectorConfigSuccess(response.data));
//             } catch (error) {
//                 console.log(error);
//                 const rpcError: GoogleRpcStatus = error?.response?.data;
//                 yield put(Actions.updateConnectorConfigFailure(rpcError));
//             }
//         }
//     );
// }


export const connectorEffects = [
    fork(onListWorkspaceConnectorsRequest),
    fork(onConnectOAuthConnectorRequest),
];

export default connectorEffects;
