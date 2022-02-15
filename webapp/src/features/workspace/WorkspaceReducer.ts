import {ArrayToObject} from '../../utils/array';
import {GoogleRpcStatus, V1Organization, V1Workspace} from '../../api';
import * as ActionType from './WorkspaceActions.types';
import * as OnboardingActionType from '../onboarding/OnboardingAction.types';

export type WorkspaceStateType = {
    attempts: number;
    loading: boolean;
    organizationId: string;
    workspaceId: string;
    organizations: V1Organization[];
    organization: V1Organization;
};

const initialWorkspaceState: WorkspaceStateType = {
    attempts: 0,
    loading: false,
    organizationId: null,
    workspaceId: null,
    organizations: [],
    organization: null
};

const authReducer = (
    state: WorkspaceStateType = initialWorkspaceState,
    action: ActionType.WorkspaceActionType | OnboardingActionType.OnboardingActionType
): WorkspaceStateType => {
    switch (action.type) {
        case ActionType.LIST_ORGANIZATIONS_REQUEST:
        case ActionType.GET_ORGANIZATION_REQUEST:
            return { ...state, loading: true, attempts: state.attempts + 1 };

        case ActionType.LIST_ORGANIZATIONS_SUCCESS:
            return {
                ...state,
                loading: false,
                organizationId: action.payload.organizations[0]?.id || null,
                workspaceId: action.payload.organizations[0]?.workspaces[0]?.id || null,
                organizations: action.payload.organizations,
                organization: state.organization ?? action.payload.organizations[0],
            };

        case ActionType.GET_ORGANIZATION_SUCCESS:
            return {
                ...state,
                loading: false,
                organizationId: action.payload.id,
                workspaceId: action.payload.workspaces[0]?.id || null,
                organization: action.payload
            };

        case ActionType.LIST_ORGANIZATIONS_FAILURE:
            return {
                ...state,
                loading: false
            };

        case ActionType.CREATE_ORGANIZATION_SUCCESS:
            return {
                ...state,
                organizationId: state.organizationId ?? action.payload.organization.id,
                workspaceId: action.payload.organization.workspaces[0]?.id || null,
                organizations: {
                    ...state.organizations,
                    [action.payload.organization.id]: action.payload.organization
                },
                organization: action.payload.organization
            };

        case OnboardingActionType.COMPLETE_ONBOARDING_SUCCESS:
            const updatedOrg = [...state.organizations]
            const index = updatedOrg.findIndex((org: V1Organization) => org.id === action.payload.organization.id);
            updatedOrg[index] = action.payload.organization;
            return { ...state, organizations: updatedOrg };

        case ActionType.CREATE_WORKSPACE_SUCCESS:
            return {
                ...state,
                workspaceId: state.workspaceId ?? action.payload.workspace.id,
                organizations: {
                    ...state.organizations,
                    [action.payload.workspace.organizationId]: {
                        ...state.organizations[action.payload.workspace.organizationId],
                        workspaces: [
                            ...state.organizations[action.payload.workspace.organizationId]
                                .workspaces,
                            action.payload.workspace
                        ]
                    }
                }
            };

        case ActionType.ADD_WORKSPACE_MEMBER_SUCCESS:
            return {
                ...state,
                ...action.payload
            };

        case ActionType.SET_CURRENT_WORKSPACE:
            return {
                ...state,
                organizationId: action.payload.organizationId,
                workspaceId: action.payload.workspaceId
            };

        // case ActionType.LIST_WORKSPACES_FAILURE:
        // case ActionType.CREATE_ORGANIZATION_FAILURE:
        // case ActionType.CREATE_WORKSPACE_FAILURE:
        // case ActionType.ADD_WORKSPACE_MEMBER_FAILURE:
        //     return {...state};

        default:
            return state;
    }
};

export default authReducer;
