import { ArrayToObject } from '../../utils/array';
import { GoogleRpcStatus, V1CatalogConnector, V1WorkspaceConnector } from '../../api';
import * as ActionType from './ConnectorsActions.types';

export type ConnectorStateType = {
  error: GoogleRpcStatus;
  catalog: Record<string, V1CatalogConnector>
  connectors: Record<string, V1WorkspaceConnector>
};

const initialConnectorState: ConnectorStateType = {
  error: null,
  catalog: {},
  connectors: {},
};

const connectorsReducer = (
  state: ConnectorStateType = initialConnectorState,
  action: ActionType.ConnectorsActionType
) => {
  switch (action.type) {
    case ActionType.LIST_CATALOG_CONNECTORS_SUCCESS:
      return {
        ...state,
        catalog: ArrayToObject(action.payload.connectors, 'id')
      };

    case ActionType.LIST_WORKSPACE_CONNECTORS_SUCCESS:
      return {
        ...state,
        connectors: ArrayToObject(action.payload.connectors, 'id')
      };

    case ActionType.CONNECT_OAUTH_CONNECTOR_SUCCESS:
      return {
        ...state,
        connectors: { ...state.connectors, [action.payload.id]: action.payload }
      };

    case ActionType.LIST_WORKSPACE_CONNECTORS_FAILURE:
    case ActionType.LIST_CATALOG_CONNECTORS_FAILURE:
    case ActionType.CONNECT_OAUTH_CONNECTOR_FAILURE:
      return { ...state, error: action.error };

    default:
      return state;
  }
};

export default connectorsReducer;
