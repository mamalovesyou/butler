import {ArrayToObject} from '../../utils/array';
import {GoogleRpcStatus, V1Organization, V1Workspace} from '../../api';
import * as ActionType from './WorkspaceActions.types';
import * as OnboardingActionType from '../onboarding/OnboardingAction.types';

export type WorkspaceStateType = {
    loading: boolean;
    organizations: V1Organization[];
    selected: {
        workspace: V1Workspace;
        organization : V1Organization;
    };
};

const initialWorkspaceState: WorkspaceStateType = {
    loading: true,
    organizations: [],
    selected: {
        workspace: null,
        organization: null
    },
};

const workspaceReducer = (
    state: WorkspaceStateType = initialWorkspaceState,
    action: ActionType.WorkspaceActionType | OnboardingActionType.OnboardingActionType
): WorkspaceStateType => {
    switch (action.type) {

        case ActionType.LIST_ORGANIZATIONS_SUCCESS:
            return {
                ...state,
                loading: false,
                organizations: action.payload.organizations,
            };

        case ActionType.LIST_ORGANIZATIONS_FAILURE:
            return {
                ...state,
                loading: false
            };



        case ActionType.ADD_WORKSPACE_MEMBER_SUCCESS:
            return {
                ...state,
                ...action.payload
            };

        case ActionType.SET_CURRENT_WORKSPACE:
            return {
                ...state,
                selected: {
                    ...state.selected,
                    workspace: action.payload,
                },
            };

        case ActionType.SET_CURRENT_ORGANIZATION:
            return {
                ...state,
                selected: {
                    ...state.selected,
                    organization: action.payload,
                },
            };
        default:
            return state;
    }
};

export default workspaceReducer;
