import {takeEvery, put, fork, call} from 'redux-saga/effects';
import * as ActionTypes from './InitAction.types';
import * as Actions from '../auth/AuthActions';
import {addAuthorization} from "../../api";
import {push} from "redux-first-history";
import {LOGIN_ROUTE_PATH} from "../../routes";


// Called when redux persist rehydrate the store
export function* onRehydrate() {
    yield takeEvery(
        ActionTypes.PERSIST_REHYDRATE,
        function* ({payload}: ActionTypes.IActionRehydrate) {
            console.log("On Rehydrate");
            const { accessToken, refreshToken } = payload.auth
            console.log("On Rehydrate: ", accessToken, refreshToken);
            if (!accessToken || !refreshToken) {
                yield put(push(LOGIN_ROUTE_PATH));
            } else {
                yield call(addAuthorization, accessToken);
                yield put(
                    Actions.refreshRequest({refreshToken: payload.auth.refreshToken})
                );
            }
        }
    );
}

export const initEffects = [fork(onRehydrate)];

export default initEffects;
