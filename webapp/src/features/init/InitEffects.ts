import {takeEvery, put, fork, select, call} from 'redux-saga/effects';
import * as ActionTypes from './InitAction.types';
import * as Actions from '../auth/AuthActions';
import {addAuthorization} from "../../api";

const getAccessToken = (state) => state.auth.accessToken;
const getRefreshToken = (state) => state.auth.refreshToken;

// Called when redux persist rehydrate the store
export function* onRehydrate() {
    yield takeEvery(
        ActionTypes.PERSIST_REHYDRATE,
        function* ({payload}: ActionTypes.IActionRehydrate) {
            const accessToken = yield select(getAccessToken);
            yield call(addAuthorization, accessToken);
            const refreshToken = yield select(getRefreshToken);
            if (refreshToken) {
                yield put(
                    Actions.refreshRequest({refreshToken: payload.auth.refreshToken})
                );
            }
        }
    );
}

export const initEffects = [fork(onRehydrate)];

export default initEffects;
