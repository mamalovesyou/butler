import { GoogleRpcStatus } from '../../api';
import * as ActionTypes from './ErrorAction.types';

export const setError = (error: GoogleRpcStatus) => ({
  type: ActionTypes.SET_ERROR,
  error: error
});

export const clearError = () => ({
  type: ActionTypes.CLEAR_ERROR
});
