import {
  V1SignInRequest,
  V1AuthenticatedUser,
  V1RefreshRequest,
  V1SignUpRequest,
  GoogleRpcStatus
} from '../../api';

export const LOGIN_REQUEST = 'LOGIN_REQUEST';
export const LOGIN_SUCCESS = 'LOGIN_SUCCESS';
export const LOGIN_FAILURE = 'LOGIN_FAILURE';

export const SIGNUP_REQUEST = 'SIGNUP_REQUEST';
export const SIGNUP_SUCCESS = 'SIGNUP_SUCCESS';
export const SIGNUP_FAILURE = 'SIGNUP_FAILURE';

export const REFRESH_TOKEN_REQUEST = 'REFRESH_TOKEN_REQUEST';
export const REFRESH_TOKEN_SUCCESS = 'REFRESH_TOKEN_SUCCESS';
export const REFRESH_TOKEN_FAILURE = 'REFRESH_TOKEN_FAILURE';

export const LOGOUT = 'LOGOUT';

// LOGIN
export interface ILoginRequest {
  type: typeof LOGIN_REQUEST;
  payload: V1SignInRequest;
}

export interface ILoginSuccess {
  type: typeof LOGIN_SUCCESS;
  payload: V1AuthenticatedUser;
}

export interface ILoginFailure {
  type: typeof LOGIN_FAILURE;
  error: GoogleRpcStatus;
}

// SIGNUP
export interface ISignupRequest {
  type: typeof SIGNUP_REQUEST;
  payload: V1SignUpRequest;
}

export interface ISignupSuccess {
  type: typeof SIGNUP_SUCCESS;
  payload: V1AuthenticatedUser;
}

export interface ISignupFailure {
  type: typeof SIGNUP_FAILURE;
  error: GoogleRpcStatus;
}

// REFRESH TOKEN
export interface IRefreshRequest {
  type: typeof REFRESH_TOKEN_REQUEST;
  payload: V1RefreshRequest;
}

export interface IRefreshSuccess {
  type: typeof REFRESH_TOKEN_SUCCESS;
  payload: V1AuthenticatedUser;
}

export interface IRefreshFailure {
  type: typeof REFRESH_TOKEN_FAILURE;
  error: GoogleRpcStatus;
}

export interface ILogout {
  type: typeof LOGOUT;
}

export type AuthActionType =
  | ILoginRequest
  | ILoginSuccess
  | ILoginFailure
  | ISignupRequest
  | ISignupSuccess
  | ISignupFailure
  | IRefreshRequest
  | IRefreshSuccess
  | IRefreshFailure
  | ILogout;
