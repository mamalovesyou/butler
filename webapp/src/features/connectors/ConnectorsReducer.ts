import {V1Connector} from '../../api';
import * as ActionType from './ConnectorsActions.types';

export type ConnectorStateType = {
    error: string;
    loading: boolean; // Used when loading catalog or data-sources
    connectors: V1Connector[];
};

const initialConnectorState: ConnectorStateType = {
    error: "",
    loading: false,
    connectors: [],
};

const connectorsReducer = (
    state: ConnectorStateType = initialConnectorState,
    action: ActionType.ConnectorsActionType
): ConnectorStateType => {
    switch (action.type) {
        case ActionType.LIST_WORKSPACE_CONNECTORS_REQUEST:
            return {...state, loading: true};


        case ActionType.LIST_WORKSPACE_CONNECTORS_SUCCESS:
            return {
                ...state,
                loading: false,
                connectors: action.payload.connectors
            };

        case ActionType.CONNECT_OAUTH_CONNECTOR_SUCCESS:
            return {
                ...state,
                connectors: {...state.connectors, [action.payload.id]: action.payload},
            };

        case ActionType.LIST_WORKSPACE_CONNECTORS_FAILURE:
            return {...state, loading: false, error: action.error.message};


        case ActionType.CONNECT_OAUTH_CONNECTOR_FAILURE:
            return {
                ...state, loading: false, error: action.error.message
            };

        default:
            return state;
    }
};

export default connectorsReducer;
