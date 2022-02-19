import {
  GoogleRpcStatus, V1DataSourceList
} from '../../api';
import * as ActionTypes from './DataSourceActions.types';

export const listAvailableSourcesRequest =
  (): ActionTypes.DataSourceActionType => ({
    type: ActionTypes.LIST_AVAILABLE_SOURCES_REQUEST
  });

export const listAvailableSourcesSuccess = (
  payload: V1DataSourceList
): ActionTypes.DataSourceActionType => ({
  type: ActionTypes.LIST_AVAILABLE_SOURCES_SUCCESS,
  payload
});

export const listAvailableSourcesFailure = (
  error: GoogleRpcStatus
): ActionTypes.DataSourceActionType => ({
  type: ActionTypes.LIST_AVAILABLE_SOURCES_FAILURE,
  error
});
