import { takeEvery, put, fork, select } from 'redux-saga/effects';
import * as ActionTypes from './InitAction.types';
import * as Actions from '../auth/AuthActions';

const getRefreshToken = state => state.auth.refreshToken;

// Called when redux persist rehydrate the store
export function* onRehydrate() {
    yield takeEvery(ActionTypes.PERSIST_REHYDRATE, function* ({ payload }:  ActionTypes.IActionRehydrate) {
        const refreshToken = yield select(getRefreshToken);
        if (refreshToken) {
            yield put(Actions.refreshRequest({refreshToken: payload.auth.refreshToken}));
        }
    });
}

export const initEffects = [
    fork(onRehydrate)
];

export default initEffects