import {ArrayToObject} from '../../utils/array';
import {V1CatalogConnector, V1ProviderAccount, V1WorkspaceConnector} from '../../api';
import * as ActionType from './ConnectorsActions.types';

export type ConnectorStateType = {
    error: string;
    loading: boolean; // Used when loading catalog or connectors
    catalog: V1CatalogConnector[];
    connectors: Record<string, V1WorkspaceConnector>;
    configure: {
        open: boolean;
        connected: boolean;
        provider: string;
        connectorId: string;
    };
    accounts: V1ProviderAccount[];
};

const initialConnectorState: ConnectorStateType = {
    error: "",
    loading: false,
    configure: {
        open: false,
        connected: false,
        provider: "",
        connectorId: "",
    },
    catalog: [],
    connectors: {},
    accounts: [],
};

const connectorsReducer = (
    state: ConnectorStateType = initialConnectorState,
    action: ActionType.ConnectorsActionType
): ConnectorStateType => {
    switch (action.type) {
        case ActionType.LIST_CATALOG_CONNECTORS_REQUEST:
        case ActionType.LIST_WORKSPACE_CONNECTORS_REQUEST:
        case ActionType.LIST_CONNECTOR_ACCOUNTS_REQUEST:
        case ActionType.UPDATE_CONNECTOR_CONFIG_REQUEST:
            return {...state, loading: true};

        case ActionType.CONNECT_OAUTH_CONNECTOR_REQUEST:
            return {
                ...state,
                loading: true,
                configure: {
                    ...initialConnectorState.configure,
                    connected: false,
                    open: true,
                    provider: action.payload.provider,
                }
            };

        case ActionType.LIST_CATALOG_CONNECTORS_SUCCESS:
            return {
                ...state,
                catalog: action.payload.connectors
            };

        case ActionType.LIST_WORKSPACE_CONNECTORS_SUCCESS:
            return {
                ...state,
                connectors: ArrayToObject(action.payload.connectors, 'id')
            };

        case ActionType.CONNECT_OAUTH_CONNECTOR_SUCCESS:
            return {
                ...state,
                connectors: {...state.connectors, [action.payload.id]: action.payload},
                configure: {
                    ...state.configure,
                    connected: true,
                    connectorId: action.payload.id
                }
            };

        case ActionType.UPDATE_CONNECTOR_CONFIG_SUCCESS:
            return {
                ...state,
                connectors: {...state.connectors, [action.payload.id]: action.payload},
                configure: initialConnectorState.configure
            };

        case ActionType.LIST_CONNECTOR_ACCOUNTS_SUCCESS:
            return {...state, loading: false, accounts: action.payload};

        case ActionType.LIST_WORKSPACE_CONNECTORS_FAILURE:
            return {...state, loading: false, error: action.error.message, connectors: {}};

        case ActionType.LIST_CATALOG_CONNECTORS_FAILURE:
            return {...state, loading: false, error: action.error.message, catalog: []};

        case ActionType.CONNECT_OAUTH_CONNECTOR_FAILURE:
            return {
                ...state, loading: false, error: action.error.message, accounts: [],
                configure: { ...state.configure, connected: false, connectorId: "" }
            };

        case ActionType.UPDATE_CONNECTOR_CONFIG_FAILURE:
            return {...state, loading: false, error: action.error.message };

        case ActionType.LIST_CONNECTOR_ACCOUNTS_FAILURE:
            return {...state, loading: false, error: action.error, accounts: []};

        case ActionType.SET_CONFIGURE_ACCOUNT_DIALOG_OPEN:
            return {...state, configure: {...state.configure, open: action.payload}}
        default:
            return state;
    }
};

export default connectorsReducer;
