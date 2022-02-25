import {takeEvery, put, fork, call, select} from 'redux-saga/effects';
import * as ActionTypes from './AuthAction.types';
import * as Actions from './AuthActions';
import * as WorkspaceActions from '../workspace/WorkspaceActions';
import {
    addAuthorization,
    GoogleRpcStatus,
    V1AuthenticatedUser
} from '../../api';
import {AxiosResponse} from 'axios';
import {persistor} from '../index';
import {push} from 'redux-first-history';
import {
    ANALYTICS_ROOT_PATH,
    DASHBOARD_ROOT_PATH,
    LOGIN_ROOT_PATH,
    ONBOARDING_ROOT_PATH
} from '../../routes';
import {Api} from '../configureEffects';
import {notificationsActions} from "../notifications";
import {listOrganizationsRequest} from "../workspace/WorkspaceActions";


// Called when a user try to login
export function* onLoginRequest() {
    yield takeEvery(
        ActionTypes.LOGIN_REQUEST,
        function* ({payload}: ActionTypes.ILoginRequest) {
            try {
                const response: AxiosResponse<V1AuthenticatedUser> =
                    yield Api.v1.usersServiceSignIn(payload);
                yield put(Actions.loginSuccess(response.data));
                yield call(addAuthorization, response.data.accessToken);
                // Load organizations
                yield put(WorkspaceActions.listOrganizationsRequest());
                yield put(push(ANALYTICS_ROOT_PATH));
            } catch (error) {
                const rpcError: GoogleRpcStatus = error?.response?.data || {
                    code: 0,
                    message: error.message
                };
                yield put(notificationsActions.createAlert({type: "error", message: rpcError.message}));
                yield put(Actions.loginFailure(rpcError));
            }
        }
    );
}

// Called when a user try to login
export function* onRefreshTokenRequest() {
    yield takeEvery(
        ActionTypes.REFRESH_TOKEN_REQUEST,
        function* ({payload}: ActionTypes.IRefreshRequest) {
            try {
                const response: AxiosResponse<V1AuthenticatedUser> =
                    yield Api.v1.usersServiceRefreshToken(payload);
                yield put(Actions.refreshSuccess(response.data));
                yield call(addAuthorization, response.data.accessToken);

            } catch (error) {
                const rpcError: GoogleRpcStatus = error?.response?.data || {
                    code: 0,
                    message: error.message
                };
                yield put(Actions.refreshFailure(rpcError));
            }
        }
    );
}

export function* onRefreshSuccess() {
    yield takeEvery(ActionTypes.REFRESH_TOKEN_SUCCESS, function* () {
        yield put(listOrganizationsRequest());
    });
}

// Redirect to login when refresh token request fails
export function* onRefreshTokenFailure() {
    yield takeEvery(ActionTypes.REFRESH_TOKEN_FAILURE, function* () {
        yield put(push(LOGIN_ROOT_PATH));
    });
}

// Called when a user try to signup
export function* onSignUpRequest() {
    yield takeEvery(
        ActionTypes.SIGNUP_REQUEST,
        function* ({payload}: ActionTypes.ISignupRequest) {
            try {
                const response: AxiosResponse<V1AuthenticatedUser> =
                    yield Api.v1.usersServiceSignUp(payload);
                yield put(Actions.signupSuccess(response.data));
                yield call(addAuthorization, response.data.accessToken);
                yield put(push(ONBOARDING_ROOT_PATH));
            } catch (error) {
                const rpcError: GoogleRpcStatus = error?.response?.data || {
                    code: 0,
                    message: error.message
                };
                yield put(notificationsActions.createAlert({type: "error", message: rpcError.message}));
                yield put(Actions.signupFailure(rpcError));
            }
        }
    );
}

export function* onSignupSuccess() {
    yield takeEvery(ActionTypes.SIGNUP_SUCCESS, function* () {
        yield put(listOrganizationsRequest());
    });
}

// Called when a user try to signup
export function* onSignUpWithInviteRequest() {
    yield takeEvery(
        ActionTypes.SIGNUP_WITH_INVITE_REQUEST,
        function* ({payload}: ActionTypes.ISignupWithInviteRequest) {
            try {
                const response: AxiosResponse<V1AuthenticatedUser> =
                    yield Api.v1.usersServiceSignUpWithInvite(payload);
                yield put(Actions.signupWithInviteSuccess(response.data));
                yield call(addAuthorization, response.data.accessToken);
                yield put(push(DASHBOARD_ROOT_PATH));
            } catch (error) {
                const rpcError: GoogleRpcStatus = error.response.data;
                yield put(Actions.signupWithInviteFailure(rpcError));
            }
        }
    );
}

// Called when a user try to logout
export function* onLogout() {
    yield takeEvery(ActionTypes.LOGOUT, function* () {
        yield call(persistor.purge);
    });
}

export const authEffects = [
    fork(onLoginRequest),
    fork(onSignUpRequest),
    fork(onSignupSuccess),
    fork(onSignUpWithInviteRequest),
    fork(onRefreshTokenRequest),
    fork(onRefreshSuccess),
    fork(onRefreshTokenFailure),
    fork(onLogout)
];

export default authEffects;
