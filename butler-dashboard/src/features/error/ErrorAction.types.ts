import { GoogleRpcStatus } from '../../api';

export const SET_ERROR = 'SET_ERROR';
export const CLEAR_ERROR = 'CLEAR_ERROR';

export interface ISetError {
  type: typeof SET_ERROR;
  error: GoogleRpcStatus;
}

export interface IClearError {
  type: typeof CLEAR_ERROR;
}

export type ErrorActionType = ISetError | IClearError;
