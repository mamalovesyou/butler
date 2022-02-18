import {takeEvery, put, fork } from 'redux-saga/effects';
import {AxiosResponse} from 'axios';
import * as ActionTypes from './DataSourceActions.types';
import * as Actions from './DataSourceActions';
import {GoogleRpcStatus, V1DataSourceList} from '../../api';
import {Api} from '../configureEffects';

export function* onListAvailableSourcesRequest() {
    yield takeEvery(ActionTypes.LIST_AVAILABLE_SOURCES_REQUEST, function* () {
        try {
            const response: AxiosResponse<V1DataSourceList> = yield Api.v1.dataSourcesServiceListAvailableSources();
            yield put(Actions.listAvailableSourcesSuccess(response.data));
        } catch (error) {
            const rpcError: GoogleRpcStatus = error?.response?.data;
            yield put(Actions.listAvailableSourcesFailure(rpcError));
        }
    });
}


export const dataSourceEffects = [
    fork(onListAvailableSourcesRequest),
];

export default dataSourceEffects;
