import {
    GoogleRpcStatus,
    V1DataSourceList,
} from '../../api';

export const LIST_AVAILABLE_SOURCES_REQUEST =
    'LIST_AVAILABLE_SOURCES_REQUEST';
export const LIST_AVAILABLE_SOURCES_SUCCESS =
    'LIST_AVAILABLE_SOURCES_SUCCESS';
export const LIST_AVAILABLE_SOURCES_FAILURE =
    'LIST_AVAILABLE_SOURCES_FAILURE';

// List Available Sources
export interface IListAvailableSourcesRequest {
    type: typeof LIST_AVAILABLE_SOURCES_REQUEST;
}

export interface IListAvailableSourcesSuccess {
    type: typeof LIST_AVAILABLE_SOURCES_SUCCESS;
    payload: V1DataSourceList;
}

export interface IListAvailableSourcesFailure {
    type: typeof LIST_AVAILABLE_SOURCES_FAILURE;
    error: GoogleRpcStatus;
}


export type DataSourceActionType =
    | IListAvailableSourcesRequest
    | IListAvailableSourcesSuccess
    | IListAvailableSourcesFailure;
