import {ArrayToObject} from '../../utils/array';
import {V1DataSource} from '../../api';
import * as ActionType from './DataSourceActions.types';

export type DataSourceStateType = {
    error: string;
    loading: boolean; // Used when loading catalog or data-sources
    sources: V1DataSource[];
};

const initialDataSourceState: DataSourceStateType = {
    error: "",
    loading: false,
    sources: [],
};

const dataSourceReducer = (
    state: DataSourceStateType = initialDataSourceState,
    action: ActionType.DataSourceActionType
): DataSourceStateType => {
    switch (action.type) {
        case ActionType.LIST_AVAILABLE_SOURCES_REQUEST:
            return {...state, loading: true};
        case ActionType.LIST_AVAILABLE_SOURCES_SUCCESS:
            return {
                ...state,
                loading: false,
                sources: action.payload.sources
            };
        case ActionType.LIST_AVAILABLE_SOURCES_FAILURE:
            return {...state, loading: false, sources: [], error: action.error.message};

        default:
            return state;
    }
};

export default dataSourceReducer;
