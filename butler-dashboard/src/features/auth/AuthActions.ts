import * as ActionTypes from './AuthAction.types';
import { GoogleRpcStatus, V1AuthenticatedUser, V1RefreshRequest, V1SignInRequest, V1SignUpRequest } from '../../api';

export const loginRequest = (payload: V1SignInRequest): ActionTypes.AuthActionType => ({
    type: ActionTypes.LOGIN_REQUEST,
    payload
});

export const loginSuccess = (payload: V1AuthenticatedUser): ActionTypes.AuthActionType => ({
    type: ActionTypes.LOGIN_SUCCESS,
    payload
});

export const loginFailure = (error: GoogleRpcStatus): ActionTypes.AuthActionType => ({
    type: ActionTypes.LOGIN_FAILURE,
    error
});


export const signupRequest = (payload: V1SignUpRequest): ActionTypes.AuthActionType => ({
    type: ActionTypes.SIGNUP_REQUEST,
    payload 
});

export const signupSuccess = (payload: V1AuthenticatedUser): ActionTypes.AuthActionType => ({
    type: ActionTypes.SIGNUP_SUCCESS,
    payload
});

export const signupFailure = (error: GoogleRpcStatus): ActionTypes.AuthActionType => ({
    type: ActionTypes.SIGNUP_FAILURE,
    error
});

export const refreshRequest = (payload: V1RefreshRequest): ActionTypes.AuthActionType => ({
    type: ActionTypes.REFRESH_TOKEN_REQUEST,
    payload 
});

export const refreshSuccess = (payload: V1AuthenticatedUser): ActionTypes.AuthActionType => ({
    type: ActionTypes.REFRESH_TOKEN_SUCCESS,
    payload 
});

export const refreshFailure = (error: GoogleRpcStatus): ActionTypes.AuthActionType => ({
    type: ActionTypes.REFRESH_TOKEN_FAILURE,
    error 
});

export const logout = (): ActionTypes.AuthActionType => ({
    type: ActionTypes.LOGOUT,
});

