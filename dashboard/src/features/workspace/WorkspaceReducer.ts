
import { ArrayToObject } from '../../utils/array';
import { GoogleRpcStatus, V1Organization, V1Workspace } from '../../api';
import * as ActionType from './WorkspaceActions.types';

export type WorkspaceStateType = {
    error: GoogleRpcStatus;

    organizationId: string;
    workspaceId: string;

    organizations: Record<string, V1Organization>;

    selectedOrganization: V1Organization;
    selectedWorkspace: V1Workspace;
}

const initialWorkspaceState: WorkspaceStateType = {
    error: null,
    organizationId: null,
    workspaceId: null,
    organizations: {},
    selectedOrganization: null,
    selectedWorkspace: null,
}

const authReducer = (state: WorkspaceStateType = initialWorkspaceState, action: ActionType.WorkspaceActionType) => {

    switch (action.type) {

        case ActionType.LIST_ORGANIZATIONS_SUCCESS:
            return {
                ...state,
                organizationId: action.payload.organizations[0]?.ID || null,
                workspaceId: action.payload.organizations[0]?.workspaces[0]?.ID || null,
                organizations: ArrayToObject(action.payload.organizations, 'ID')
            }

        case ActionType.CREATE_ORGANIZATION_SUCCESS:
            return { ...state,
                organizationId: state.organizationId ?? action.payload.organization.ID,
                organizations: { 
                    ...state.organizations,
                    [action.payload.organization.ID]: action.payload.organization
                  }
                };

        case ActionType.CREATE_WORKSPACE_SUCCESS:
            return { ...state,
                workspaceId: state.workspaceId ?? action.payload.workspace.ID,
                organizations: { 
                    ...state.organizations,
                    [action.payload.workspace.organizationID]: {
                        ...state.organizations[action.payload.workspace.organizationID],
                        workspaces: [
                            ...state.organizations[action.payload.workspace.organizationID].workspaces,
                            action.payload.workspace
                        ]
                    }
                  }
                };

        case ActionType.ADD_WORKSPACE_MEMBER_SUCCESS:
            return {
                ...state,
                ...action.payload
            }
        
        case ActionType.SET_CURRENT_WORKSPACE:
            return {
                ...state,
                organizationId: action.payload.organizationId,
                workspaceId: action.payload.workspaceId
            }


        case ActionType.LIST_WORKSPACES_FAILURE:
        case ActionType.CREATE_ORGANIZATION_FAILURE:
        case ActionType.CREATE_WORKSPACE_FAILURE:
        case ActionType.ADD_WORKSPACE_MEMBER_FAILURE:
            return { ...state, error: action.error };

        default:
            return state
    }
}

export default authReducer;