import * as ActionTypes from './WorkspaceActions.types';
import { GoogleRpcStatus, V1CreateOrganizationRequest, V1CreateWorkspaceRequest, V1OrganizationListResponse, V1OrganizationResponse, V1AuthenticatedUser, V1SignInRequest, V1WorkspaceResponse } from '../../api';

export const createOrganizationRequest = (payload: V1CreateOrganizationRequest): ActionTypes.WorkspaceActionType => ({
    type: ActionTypes.CREATE_ORGANIZATION_REQUEST,
    payload
});

export const createOrganizationSuccess = (payload: V1OrganizationResponse): ActionTypes.WorkspaceActionType => ({
    type: ActionTypes.CREATE_ORGANIZATION_SUCCESS,
    payload
});

export const createOrganizationFailure = (error: GoogleRpcStatus): ActionTypes.WorkspaceActionType => ({
    type: ActionTypes.CREATE_ORGANIZATION_FAILURE,
    error
});

export const listOrganizationsRequest = (): ActionTypes.WorkspaceActionType => ({
    type: ActionTypes.LIST_ORGANIZATIONS_REQUEST,
});

export const listOrganizationsSuccess = (payload: V1OrganizationListResponse): ActionTypes.WorkspaceActionType => ({
    type: ActionTypes.LIST_ORGANIZATIONS_SUCCESS,
    payload
});

export const listOrganizationsFailure = (error: GoogleRpcStatus): ActionTypes.WorkspaceActionType => ({
    type: ActionTypes.LIST_ORGANIZATIONS_FAILURE,
    error
});

export const listWorkspacesSuccess = (payload: V1OrganizationListResponse): ActionTypes.WorkspaceActionType => ({
    type: ActionTypes.LIST_WORKSPACES_SUCCESS,
    payload
});

export const listWorkspacesFailure = (error: GoogleRpcStatus): ActionTypes.WorkspaceActionType => ({
    type: ActionTypes.LIST_WORKSPACES_FAILURE,
    error
});


export const createWorkspaceRequest = (payload: V1CreateWorkspaceRequest): ActionTypes.WorkspaceActionType => ({
    type: ActionTypes.CREATE_WORKSPACE_REQUEST,
    payload
});

export const createWorkspaceSuccess = (payload: V1WorkspaceResponse): ActionTypes.WorkspaceActionType => ({
    type: ActionTypes.CREATE_WORKSPACE_SUCCESS,
    payload
});

export const createWorkspaceFailure = (error: GoogleRpcStatus): ActionTypes.WorkspaceActionType => ({
    type: ActionTypes.CREATE_WORKSPACE_FAILURE,
    error
});

export const addWorkspaceMembersRequest = (payload: V1SignInRequest): ActionTypes.WorkspaceActionType => ({
    type: ActionTypes.ADD_WORKSPACE_MEMBER_REQUEST,
    payload
});

export const addWorkspaceMembersSuccess = (payload: V1AuthenticatedUser): ActionTypes.WorkspaceActionType => ({
    type: ActionTypes.ADD_WORKSPACE_MEMBER_SUCCESS,
    payload
});

export const addWorkspaceMembersFailure = (error: GoogleRpcStatus): ActionTypes.WorkspaceActionType => ({
    type: ActionTypes.ADD_WORKSPACE_MEMBER_FAILURE,
    error
});

export const setCurrentWorkspace = (payload: { organizationId: string, workspaceId: string}): ActionTypes.WorkspaceActionType => ({
    type: ActionTypes.SET_CURRENT_WORKSPACE,
    payload
});

