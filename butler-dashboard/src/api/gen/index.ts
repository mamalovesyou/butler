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
  id?: string;
  name?: string;
  iconUrl?: string;
  authType?: string;
  authUrl?: string;
}

export interface V1CatalogConnectorList {
  connectors?: V1CatalogConnector[];
}

export interface V1CreateOrganizationRequest {
  name?: string;
  userRole?: string;
}

export interface V1CreateWorkspaceRequest {
  organizationID?: string;
  name?: string;
  description?: string;
}

export interface V1IsValidAccessTokenResponse {
  userID?: string;
}

export interface V1ListUsersResponse {
  users?: V1User[];
}

export interface V1OAuthAuthorizationRequestConnectorCode {
  name?: string;
  code?: string;
}

export interface V1Organization {
  ID?: string;
  name?: string;
  ownerID?: string;
  workspaces?: V1Workspace[];
  members?: V1UserMember[];

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

/**
 * The request message containing the refresh token.
 */
export interface V1RefreshRequest {
  refreshToken?: string;
}

/**
 * The request message containing the email and password.
 */
export interface V1SignInRequest {
  /**
   * option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_schema) = {
   *    json_schema: {
   *      title: "SignIn"
   *      description: "Intentionaly complicated message type to cover many features of Protobuf."
   *      required: ["email", "password"]
   *    }
   *    example: "{\"email\": \"john@heybutler.io\", \"password\": \"AStr0ngP@ssWord!\" }"
   *  };
   */
  email?: string;
  password?: string;
}

/**
 * The request message containing the access token.
 */
export interface V1SignOutRequest {
  accessToken?: string;
}

/**
 * The request message containing the name, email and password.
 */
export interface V1SignUpRequest {
  /**
   * option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_schema) = {
   *    json_schema: {
   *      title: "SignUp"
   *      description: "Welcome to BeyButler! We are glad you want to join us :)"
   *      required: ["name", "email", "password"]
   *    }
   *    example: "{\"companyName\": \"Hey Butler Inc.\", \"firstName\": \"john\", \"lastName \": \"john\", \"email\": \"john@gobaboon.co\", \"password\": \"AStr0ngP@ssWord!\" }"
   *  };
   */
  lastName?: string;
  firstName?: string;
  email?: string;
  password?: string;
  companyName?: string;
  companyRole?: string;
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
  userID?: string;
  firstName?: string;
  lastName?: string;
  role?: string;
}

export interface V1Workspace {
  ID?: string;
  name?: string;
  description?: string;
  organizationID?: string;
  members?: V1UserMember[];

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
 * @title services/proto/auth-service.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  v1 = {
    /**
     * No description
     *
     * @tags AuthService
     * @name AuthServiceRefreshToken
     * @request POST:/v1/auth/refresh
     * @secure
     */
    authServiceRefreshToken: (body: V1RefreshRequest, params: RequestParams = {}) =>
      this.request<V1AuthenticatedUser, GoogleRpcStatus>({
        path: `/v1/auth/refresh`,
        method: "POST",
        body: body,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags AuthService
     * @name AuthServiceSignIn
     * @request POST:/v1/auth/signin
     * @secure
     */
    authServiceSignIn: (body: V1SignInRequest, params: RequestParams = {}) =>
      this.request<V1AuthenticatedUser, GoogleRpcStatus>({
        path: `/v1/auth/signin`,
        method: "POST",
        body: body,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags AuthService
     * @name AuthServiceSignOut
     * @request POST:/v1/auth/signout
     * @secure
     */
    authServiceSignOut: (body: V1SignOutRequest, params: RequestParams = {}) =>
      this.request<any, GoogleRpcStatus>({
        path: `/v1/auth/signout`,
        method: "POST",
        body: body,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags AuthService
     * @name AuthServiceSignUp
     * @request POST:/v1/auth/signup
     * @secure
     */
    authServiceSignUp: (body: V1SignUpRequest, params: RequestParams = {}) =>
      this.request<V1AuthenticatedUser, GoogleRpcStatus>({
        path: `/v1/auth/signup`,
        method: "POST",
        body: body,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags WorkspaceService
     * @name WorkspaceServiceListOrganizations
     * @request GET:/v1/organizations
     * @secure
     */
    workspaceServiceListOrganizations: (params: RequestParams = {}) =>
      this.request<V1OrganizationListResponse, GoogleRpcStatus>({
        path: `/v1/organizations`,
        method: "GET",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags WorkspaceService
     * @name WorkspaceServiceCreateOrganization
     * @request POST:/v1/organizations
     * @secure
     */
    workspaceServiceCreateOrganization: (body: V1CreateOrganizationRequest, params: RequestParams = {}) =>
      this.request<V1OrganizationResponse, GoogleRpcStatus>({
        path: `/v1/organizations`,
        method: "POST",
        body: body,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags WorkspaceService
     * @name WorkspaceServiceCreateWorkspace
     * @request POST:/v1/workspaces
     * @secure
     */
    workspaceServiceCreateWorkspace: (body: V1CreateWorkspaceRequest, params: RequestParams = {}) =>
      this.request<V1WorkspaceResponse, GoogleRpcStatus>({
        path: `/v1/workspaces`,
        method: "POST",
        body: body,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags ConnectorsService
     * @name ConnectorsServiceListCatalogConnectors
     * @request GET:/v1/{workspaceId}/catalogs
     * @secure
     */
    connectorsServiceListCatalogConnectors: (workspaceId: string, params: RequestParams = {}) =>
      this.request<V1CatalogConnectorList, GoogleRpcStatus>({
        path: `/v1/${workspaceId}/catalogs`,
        method: "GET",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags ConnectorsService
     * @name ConnectorsServiceListWorkspaceConnectors
     * @request GET:/v1/{workspaceId}/connectors
     * @secure
     */
    connectorsServiceListWorkspaceConnectors: (workspaceId: string, params: RequestParams = {}) =>
      this.request<V1WorkspaceConnectorList, GoogleRpcStatus>({
        path: `/v1/${workspaceId}/connectors`,
        method: "GET",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags ConnectorsService
     * @name ConnectorsServiceGetOauthConnectorAuthorization
     * @request POST:/v1/{workspaceId}/connectors/oauth
     * @secure
     */
    connectorsServiceGetOauthConnectorAuthorization: (
      workspaceId: string,
      body: V1OAuthAuthorizationRequestConnectorCode,
      params: RequestParams = {},
    ) =>
      this.request<V1WorkspaceConnector, GoogleRpcStatus>({
        path: `/v1/${workspaceId}/connectors/oauth`,
        method: "POST",
        body: body,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
}
