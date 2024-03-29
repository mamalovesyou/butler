import {
    GoogleRpcStatus,
    V1CreateOrganizationRequest,
    V1OrganizationListResponse,
    V1OrganizationResponse,
    V1AuthenticatedUser,
    V1WorkspaceResponse,
    V1CreateWorkspaceRequest,
    V1GetOrganizationRequest,
    V1Organization, V1Workspace, V1BatchInviteMemberRequest,
} from '../../api';

export const LIST_ORGANIZATIONS_REQUEST = 'LIST_ORGANIZATIONS_REQUEST';
export const LIST_ORGANIZATIONS_SUCCESS = 'LIST_ORGANIZATIONS_SUCCESS';
export const LIST_ORGANIZATIONS_FAILURE = 'LIST_ORGANIZATIONS_FAILURE';

export const CREATE_ORGANIZATION_REQUEST = 'CREATE_ORGANIZATION_REQUEST';
export const CREATE_ORGANIZATION_SUCCESS = 'CREATE_ORGANIZATION_SUCCESS';
export const CREATE_ORGANIZATION_FAILURE = 'CREATE_ORGANIZATION_FAILURE';

export const GET_ORGANIZATION_REQUEST = 'GET_ORGANIZATION_REQUEST';
export const GET_ORGANIZATION_SUCCESS = 'GET_ORGANIZATION_SUCCESS';
export const GET_ORGANIZATION_FAILURE = 'GET_ORGANIZATION_FAILURE';

export const LIST_WORKSPACES_REQUEST = 'LIST_WORKSPACES_REQUEST';
export const LIST_WORKSPACES_SUCCESS = 'LIST_WORKSPACES_SUCCESS';
export const LIST_WORKSPACES_FAILURE = 'LIST_WORKSPACES_FAILURE';

export const CREATE_WORKSPACE_REQUEST = 'CREATE_WORKSPACE_REQUEST';
export const CREATE_WORKSPACE_SUCCESS = 'CREATE_WORKSPACE_SUCCESS';
export const CREATE_WORKSPACE_FAILURE = 'CREATE_WORKSPACE_FAILURE';

export const ADD_WORKSPACE_MEMBER_REQUEST = 'ADD_WORKSPACE_MEMBER_REQUEST';
export const ADD_WORKSPACE_MEMBER_SUCCESS = 'ADD_WORKSPACE_MEMBER_SUCCESS';
export const ADD_WORKSPACE_MEMBER_FAILURE = 'ADD_WORKSPACE_MEMBER_FAILURE';

export const SET_CURRENT_ORGANIZATION = 'SET_CURRENT_ORGANIZATION';
export const SET_CURRENT_WORKSPACE = 'SET_CURRENT_WORKSPACE';

// CreateOrganization
export interface ICreateOrganizationRequest {
    type: typeof CREATE_ORGANIZATION_REQUEST;
    payload: V1CreateOrganizationRequest;
}

export interface ICreateOrganizationSuccess {
    type: typeof CREATE_ORGANIZATION_SUCCESS;
    payload: V1OrganizationResponse;
}

export interface ICreateOrganizationFailure {
    type: typeof CREATE_ORGANIZATION_FAILURE;
    error: GoogleRpcStatus;
}

// GetOrganization
export interface IGetOrganizationRequest {
    type: typeof GET_ORGANIZATION_REQUEST;
    payload: V1GetOrganizationRequest;
}

export interface IGetOrganizationSuccess {
    type: typeof GET_ORGANIZATION_SUCCESS;
    payload: V1Organization;
}

export interface IGetOrganizationFailure {
    type: typeof GET_ORGANIZATION_FAILURE;
    error: GoogleRpcStatus;
}

// CreateWorkspaceDialog
export interface IListOrganizationsRequest {
    type: typeof LIST_ORGANIZATIONS_REQUEST;
}

export interface IListOrganizationsSuccess {
    type: typeof LIST_ORGANIZATIONS_SUCCESS;
    payload: V1OrganizationListResponse;
}

export interface IListOrganizationsFailure {
    type: typeof LIST_ORGANIZATIONS_FAILURE;
    error: GoogleRpcStatus;
}

// CreateWorkspaceDialog
export interface IListWorkspacesRequest {
    type: typeof LIST_WORKSPACES_REQUEST;
}

export interface IListWorkspacesSuccess {
    type: typeof LIST_WORKSPACES_SUCCESS;
    payload: V1OrganizationListResponse;
}

export interface IListWorkspacesFailure {
    type: typeof LIST_WORKSPACES_FAILURE;
    error: GoogleRpcStatus;
}

// CreateWorkspaceDialog
export interface ICreateWorkspaceRequest {
    type: typeof CREATE_WORKSPACE_REQUEST;
    payload: V1CreateWorkspaceRequest;
}

export interface ICreateWorkspaceSuccess {
    type: typeof CREATE_WORKSPACE_SUCCESS;
    payload: V1WorkspaceResponse;
}

export interface ICreateWorkspaceFailure {
    type: typeof CREATE_WORKSPACE_FAILURE;
    error: GoogleRpcStatus;
}

// AddWorkspaceMembers
export interface IAddWorkspaceMembersRequest {
    type: typeof ADD_WORKSPACE_MEMBER_REQUEST;
    payload: V1BatchInviteMemberRequest;
}

export interface IAddWorkspaceMembersSuccess {
    type: typeof ADD_WORKSPACE_MEMBER_SUCCESS;
    payload: V1AuthenticatedUser;
}

export interface IAddWorkspaceMembersFailure {
    type: typeof ADD_WORKSPACE_MEMBER_FAILURE;
    error: GoogleRpcStatus;
}


export interface ISetCurrentOrganization {
    type: typeof SET_CURRENT_ORGANIZATION;
    payload: V1Organization;
}

export interface ISetCurrentWorkspace {
    type: typeof SET_CURRENT_WORKSPACE;
    payload: V1Workspace;
}

export type WorkspaceActionType =
    | IListOrganizationsRequest
    | IListOrganizationsSuccess
    | IListOrganizationsFailure
    | ICreateOrganizationRequest
    | ICreateOrganizationSuccess
    | ICreateOrganizationFailure
    | IGetOrganizationRequest
    | IGetOrganizationSuccess
    | IGetOrganizationFailure
    | ICreateWorkspaceRequest
    | ICreateWorkspaceSuccess
    | ICreateWorkspaceFailure
    | IListWorkspacesRequest
    | IListWorkspacesSuccess
    | IListWorkspacesFailure
    | IAddWorkspaceMembersRequest
    | IAddWorkspaceMembersSuccess
    | IAddWorkspaceMembersFailure
    | ISetCurrentWorkspace
    | ISetCurrentOrganization;
