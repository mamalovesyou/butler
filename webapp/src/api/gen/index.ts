/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface GoogleProtobufAny {
  "@type"?: string;
}

export interface GoogleRpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: GoogleProtobufAny[];
}

/**
 * The response message containing the User.
 */
export interface V1AuthenticatedUser {
  user?: V1User;
  accessToken?: string;
  refreshToken?: string;
}

export interface V1CatalogConnector {
  name?: string;
  iconSvg?: string;
  authType?: V1CatalogConnectorAuthType;
  authUrl?: string;
}

export enum V1CatalogConnectorAuthType {
  OAUTH2 = "OAUTH2",
  API_KEY = "API_KEY",
}

export interface V1CatalogConnectorList {
  connectors?: V1CatalogConnector[];
}

export interface V1ConnectWithCodeRequest {
  workspaceId?: string;
  provider?: string;
  code?: string;
}

export interface V1CreateOrganizationRequest {
  name?: string;
  userRole?: string;
}

export interface V1CreateWorkspaceRequest {
  organizationId?: string;
  workspace?: V1CreateWorkspaceRequestWorkspaceInfo;
}

export interface V1CreateWorkspaceRequestWorkspaceInfo {
  name?: string;
  description?: string;
}

export interface V1Invitation {
  id?: string;
  firstName?: string;
  lastName?: string;
  email?: string;
  role?: string;

  /** @format date-time */
  expiresAt?: string;

  /** @format date-time */
  createdAt?: string;
}

export interface V1InviteInfos {
  firstName?: string;
  lastName?: string;
  email?: string;
  role?: string;
}

export interface V1InviteOrganizationMemberRequest {
  invitation?: V1InviteInfos;
  organizationId?: string;
}

export interface V1InviteWorkspaceMemberRequest {
  invitation?: V1InviteInfos;
  workspaceId?: string;
}

export interface V1Organization {
  id?: string;
  name?: string;
  ownerId?: string;
  workspaces?: V1Workspace[];
  members?: V1UserMember[];
  invitations?: V1Invitation[];

  /** @format date-time */
  createdAt?: string;

  /** @format date-time */
  updatedAt?: string;
}

export interface V1OrganizationListResponse {
  organizations?: V1Organization[];
}

export interface V1OrganizationResponse {
  organization?: V1Organization;
}

export interface V1RefreshRequest {
  refreshToken?: string;
}

export interface V1SignInRequest {
  email?: string;
  password?: string;
}

export interface V1SignOutRequest {
  accessToken?: string;
}

export interface V1SignUpRequest {
  lastName?: string;
  firstName?: string;
  email?: string;
  password?: string;
  companyName?: string;
  companyRole?: string;
}

export interface V1SignUpWithInvitationRequest {
  invitationId?: string;
  infos?: V1SignUpWithInvitationRequestSignupInfo;
}

export interface V1SignUpWithInvitationRequestSignupInfo {
  organizationId?: string;
  workspaceId?: string;
  token?: string;
  password?: string;
}

export interface V1User {
  ID?: string;
  email?: string;
  firstName?: string;
  lastName?: string;

  /** @format date-time */
  createdAt?: string;

  /** @format date-time */
  updatedAt?: string;
}

export interface V1UserMember {
  userId?: string;
  firstName?: string;
  lastName?: string;
  role?: string;

  /** @format date-time */
  updatedAt?: string;

  /** @format date-time */
  createdAt?: string;
}

export interface V1Workspace {
  id?: string;
  name?: string;
  description?: string;
  organizationId?: string;
  members?: V1UserMember[];
  invitations?: V1Invitation[];

  /** @format date-time */
  createdAt?: string;

  /** @format date-time */
  updatedAt?: string;
}

export interface V1WorkspaceConnector {
  id?: string;
  workspaceId?: string;
  name?: string;
  status?: string;

  /** @format date-time */
  expiresIn?: string;

  /** @format date-time */
  createdAt?: string;

  /** @format date-time */
  updatedAt?: string;
}

export interface V1WorkspaceConnectorList {
  connectors?: V1WorkspaceConnector[];
}

export interface V1WorkspaceConnectorsRequest {
  workspaceId?: string;
}

export interface V1WorkspaceResponse {
  workspace?: V1Workspace;
}

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.instance.defaults.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createFormData(input: Record<string, unknown>): FormData {
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      formData.append(
        key,
        property instanceof Blob
          ? property
          : typeof property === "object" && property !== null
          ? JSON.stringify(property)
          : `${property}`,
      );
      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = (format && this.format) || void 0;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      requestParams.headers.common = { Accept: "*/*" };
      requestParams.headers.post = {};
      requestParams.headers.put = {};

      body = this.createFormData(body as Record<string, unknown>);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
    });
  };
}

/**
 * @title google/api/http.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  v1 = {
    /**
     * No description
     *
     * @tags OctopusService
     * @name OctopusServiceListWorkspaceConnectors
     * @request POST:/v1/connector
     */
    octopusServiceListWorkspaceConnectors: (body: V1WorkspaceConnectorsRequest, params: RequestParams = {}) =>
      this.request<V1WorkspaceConnectorList, GoogleRpcStatus>({
        path: `/v1/connector`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags OctopusService
     * @name OctopusServiceGetCatalogConnectors
     * @request GET:/v1/connector/catalog
     */
    octopusServiceGetCatalogConnectors: (params: RequestParams = {}) =>
      this.request<V1CatalogConnectorList, GoogleRpcStatus>({
        path: `/v1/connector/catalog`,
        method: "GET",
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags OctopusService
     * @name OctopusServiceConnectWithCode
     * @request POST:/v1/connector/connect/oauth
     */
    octopusServiceConnectWithCode: (body: V1ConnectWithCodeRequest, params: RequestParams = {}) =>
      this.request<V1WorkspaceConnector, GoogleRpcStatus>({
        path: `/v1/connector/connect/oauth`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags UsersService
     * @name UsersServiceListOrganizations
     * @summary Organization
     * @request GET:/v1/organizations
     */
    usersServiceListOrganizations: (params: RequestParams = {}) =>
      this.request<V1OrganizationListResponse, GoogleRpcStatus>({
        path: `/v1/organizations`,
        method: "GET",
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags UsersService
     * @name UsersServiceCreateOrganization
     * @request POST:/v1/organizations
     */
    usersServiceCreateOrganization: (body: V1CreateOrganizationRequest, params: RequestParams = {}) =>
      this.request<V1OrganizationResponse, GoogleRpcStatus>({
        path: `/v1/organizations`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags UsersService
     * @name UsersServiceInviteOrganizationMember
     * @request POST:/v1/organizations/members/invite
     */
    usersServiceInviteOrganizationMember: (body: V1InviteOrganizationMemberRequest, params: RequestParams = {}) =>
      this.request<V1Invitation, GoogleRpcStatus>({
        path: `/v1/organizations/members/invite`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags UsersService
     * @name UsersServiceCreateWorkspace
     * @summary Workspace
     * @request POST:/v1/organizations/workspaces
     */
    usersServiceCreateWorkspace: (body: V1CreateWorkspaceRequest, params: RequestParams = {}) =>
      this.request<V1WorkspaceResponse, GoogleRpcStatus>({
        path: `/v1/organizations/workspaces`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags UsersService
     * @name UsersServiceRefreshToken
     * @request POST:/v1/user/refresh
     */
    usersServiceRefreshToken: (body: V1RefreshRequest, params: RequestParams = {}) =>
      this.request<V1AuthenticatedUser, GoogleRpcStatus>({
        path: `/v1/user/refresh`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags UsersService
     * @name UsersServiceSignIn
     * @summary Authentication
     * @request POST:/v1/user/signin
     */
    usersServiceSignIn: (body: V1SignInRequest, params: RequestParams = {}) =>
      this.request<V1AuthenticatedUser, GoogleRpcStatus>({
        path: `/v1/user/signin`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags UsersService
     * @name UsersServiceSignOut
     * @request POST:/v1/user/signout
     */
    usersServiceSignOut: (body: V1SignOutRequest, params: RequestParams = {}) =>
      this.request<any, GoogleRpcStatus>({
        path: `/v1/user/signout`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags UsersService
     * @name UsersServiceSignUp
     * @request POST:/v1/user/signup
     */
    usersServiceSignUp: (body: V1SignUpRequest, params: RequestParams = {}) =>
      this.request<V1AuthenticatedUser, GoogleRpcStatus>({
        path: `/v1/user/signup`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags UsersService
     * @name UsersServiceSignUpWithInvite
     * @request POST:/v1/user/signup/invites
     */
    usersServiceSignUpWithInvite: (body: V1SignUpWithInvitationRequest, params: RequestParams = {}) =>
      this.request<V1AuthenticatedUser, GoogleRpcStatus>({
        path: `/v1/user/signup/invites`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags UsersService
     * @name UsersServiceInviteWorkspaceMember
     * @request POST:/v1/workspaces/members/invite
     */
    usersServiceInviteWorkspaceMember: (body: V1InviteWorkspaceMemberRequest, params: RequestParams = {}) =>
      this.request<V1Invitation, GoogleRpcStatus>({
        path: `/v1/workspaces/members/invite`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
}
