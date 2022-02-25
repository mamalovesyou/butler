import {takeEvery, put, fork, call, select} from 'redux-saga/effects';
import * as ActionTypes from './InitAction.types';
import * as Actions from '../auth/AuthActions';
import {addAuthorization} from "../../api";
import {RootState} from "../index";
import {LOGOUT_ROOT_PATH} from "../../routes";


// Called when redux persist rehydrate the store
export function* onRehydrate() {
    yield takeEvery(
        ActionTypes.PERSIST_REHYDRATE,
        function* ({payload, key}: ActionTypes.IActionRehydrate) {
            const getPath = (state: RootState) => state.router.location;
            const { pathname } = yield select(getPath);
            if (pathname !== LOGOUT_ROOT_PATH && key === "root") {
                // @ts-ignore
                const { accessToken, refreshToken } = payload.auth;
                if (refreshToken) {
                    yield call(addAuthorization, accessToken);
                    yield put(
                        Actions.refreshRequest({refreshToken})
                    );
                }
            }
        }
    );
}

export const initEffects = [fork(onRehydrate)];

export default initEffects;
