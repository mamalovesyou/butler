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

/**
* `Any` contains an arbitrary serialized protocol buffer message along with a
URL that describes the type of the serialized message.

Protobuf library provides support to pack/unpack Any values in the form
of utility functions or additional generated methods of the Any type.

Example 1: Pack and unpack a message in C++.

    Foo foo = ...;
    Any any;
    any.PackFrom(foo);
    ...
    if (any.UnpackTo(&foo)) {
      ...
    }

Example 2: Pack and unpack a message in Java.

    Foo foo = ...;
    Any any = Any.pack(foo);
    ...
    if (any.is(Foo.class)) {
      foo = any.unpack(Foo.class);
    }

 Example 3: Pack and unpack a message in Python.

    foo = Foo(...)
    any = Any()
    any.Pack(foo)
    ...
    if any.Is(Foo.DESCRIPTOR):
      any.Unpack(foo)
      ...

 Example 4: Pack and unpack a message in Go

     foo := &pb.Foo{...}
     any, err := anypb.New(foo)
     if err != nil {
       ...
     }
     ...
     foo := &pb.Foo{}
     if err := any.UnmarshalTo(foo); err != nil {
       ...
     }

The pack methods provided by protobuf library will by default use
'type.googleapis.com/full.type.name' as the type URL and the unpack
methods only use the fully qualified type name after the last '/'
in the type URL, for example "foo.bar.com/x/y.z" will yield type
name "y.z".


JSON
====
The JSON representation of an `Any` value uses the regular
representation of the deserialized, embedded message, with an
additional field `@type` which contains the type URL. Example:

    package google.profile;
    message Person {
      string first_name = 1;
      string last_name = 2;
    }

    {
      "@type": "type.googleapis.com/google.profile.Person",
      "firstName": <string>,
      "lastName": <string>
    }

If the embedded message type is well-known and has a custom JSON
representation, that representation will be embedded adding a field
`value` which holds the custom JSON in addition to the `@type`
field. Example (for message [google.protobuf.Duration][]):

    {
      "@type": "type.googleapis.com/google.protobuf.Duration",
      "value": "1.212s"
    }
*/
export interface GoogleProtobufAny {
  /**
   * A URL/resource name that uniquely identifies the type of the serialized
   * protocol buffer message. This string must contain at least
   * one "/" character. The last segment of the URL's path must represent
   * the fully qualified name of the type (as in
   * `path/google.protobuf.Duration`). The name should be in a canonical form
   * (e.g., leading "." is not accepted).
   *
   * In practice, teams usually precompile into the binary all types that they
   * expect it to use in the context of Any. However, for URLs which use the
   * scheme `http`, `https`, or no scheme, one can optionally set up a type
   * server that maps type URLs to message definitions as follows:
   * * If no scheme is provided, `https` is assumed.
   * * An HTTP GET on the URL must yield a [google.protobuf.Type][]
   *   value in binary format, or produce an error.
   * * Applications are allowed to cache lookup results based on the
   *   URL, or have them precompiled into a binary to avoid any
   *   lookup. Therefore, binary compatibility needs to be preserved
   *   on changes to types. (Use versioned type names to manage
   *   breaking changes.)
   * Note: this functionality is not currently available in the official
   * protobuf release, and it is not used for type URLs beginning with
   * type.googleapis.com.
   * Schemes other than `http`, `https` (or the empty scheme) might be
   * used with implementation specific semantics.
   */
  "@type"?: string;
}

/**
* `NullValue` is a singleton enumeration to represent the null value for the
`Value` type union.

 The JSON representation for `NullValue` is JSON `null`.

 - NULL_VALUE: Null value.
*/
export enum GoogleProtobufNullValue {
  NULL_VALUE = "NULL_VALUE",
}

export interface GoogleRpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: GoogleProtobufAny[];
}

export enum V1AuthType {
  OAUTH2 = "OAUTH2",
  API_KEY = "API_KEY",
}

export interface V1AuthenticateConnectorRequest {
  connectorId?: string;
  code?: string;
}

/**
 * The response message containing the User.
 */
export interface V1AuthenticatedUser {
  user?: V1User;
  accessToken?: string;
  refreshToken?: string;
}

export interface V1BatchInviteMemberRequest {
  emails?: string[];
  organizationId?: string;
  workspaceId?: string;
}

export interface V1CompleteOnboardingRequest {
  organizationId?: string;
}

export interface V1ConnectWithCodeRequest {
  workspaceId?: string;
  provider?: string;
  airbyteSourceId?: string;
  code?: string;
}

export interface V1Connector {
  id?: string;
  workspaceId?: string;
  name?: string;
  airbyteSourceDefinitionId?: string;
  authScheme?: V1AuthType;
  config?: object;

  /** @format date-time */
  updatedAt?: string;
}

export interface V1ConnectorList {
  connectors?: V1Connector[];
}

export interface V1ConnectorSecret {
  value?: object;
}

export interface V1ConnectorSecretPair {
  connector?: V1WorkspaceConnector;
  credentials?: V1ConnectorSecret;
}

export interface V1CreateConnectorRequest {
  workspaceId?: string;
  airbyteWorkspaceId?: string;
  airbyteSourceDefinitionId?: string;
  secrets?: object;
  config?: object;
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

export interface V1DataSource {
  name?: string;
  iconSvg?: string;
  authType?: V1AuthType;
  authUrl?: string;
  configurationInputJSONSchema?: string;
  airbyteSourceDefinitionId?: string;
}

export interface V1DataSourceList {
  sources?: V1DataSource[];
}

export interface V1GetConnectorRequest {
  connectorId?: string;
}

export interface V1GetInvitationRequest {
  invitationId?: string;
  token?: string;
}

export interface V1GetOrganizationRequest {
  organizationId?: string;
}

export interface V1GetWorkspaceRequest {
  workspaceId?: string;
}

export interface V1Invitation {
  id?: string;
  email?: string;
  token?: string;
  organization?: V1Organization;
  workspace?: V1Workspace;

  /** @format date-time */
  expiresAt?: string;
}

export interface V1InvitationListResponse {
  invitations?: V1Invitation[];
}

export interface V1ListConnectorsRequest {
  workspaceId?: string;
}

export interface V1MutateConnectorRequest {
  connectorId?: string;
  secrets?: object;
  config?: object;
}

export interface V1Organization {
  id?: string;
  name?: string;
  ownerId?: string;
  onboarded?: boolean;
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

export interface V1ProviderAccount {
  name?: string;
  id?: string;
  test?: boolean;
  currency?: string;
}

export interface V1RefreshRequest {
  refreshToken?: string;
}

export interface V1SelectAccountRequest {
  workspaceConnectorId?: string;
  accountId?: string;
  accountName?: string;
  isTestAccount?: boolean;
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
  lastName?: string;
  firstName?: string;
  password?: string;
  invitationId?: string;
  token?: string;
}

export interface V1TestConnectionRequest {
  connectorId?: string;
}

export interface V1TestConnectionResponse {
  status?: string;
  message?: string;
  logs?: string[];
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
  airbyteWorkspaceId?: string;
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
  authScheme?: V1AuthType;
  accountConfig?: V1ProviderAccount;

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
     * @tags ConnectorsService
     * @name ConnectorsServiceCreateConnector
     * @request POST:/v1/connectors
     */
    connectorsServiceCreateConnector: (body: V1CreateConnectorRequest, params: RequestParams = {}) =>
      this.request<V1Connector, GoogleRpcStatus>({
        path: `/v1/connectors`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags ConnectorsService
     * @name ConnectorsServiceMutateConnector
     * @request PUT:/v1/connectors
     */
    connectorsServiceMutateConnector: (body: V1MutateConnectorRequest, params: RequestParams = {}) =>
      this.request<V1Connector, GoogleRpcStatus>({
        path: `/v1/connectors`,
        method: "PUT",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags ConnectorsService
     * @name ConnectorsServiceGetConnector
     * @request POST:/v1/connectors/get
     */
    connectorsServiceGetConnector: (body: V1GetConnectorRequest, params: RequestParams = {}) =>
      this.request<V1Connector, GoogleRpcStatus>({
        path: `/v1/connectors/get`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags ConnectorsService
     * @name ConnectorsServiceListConnectors
     * @request POST:/v1/connectors/list
     */
    connectorsServiceListConnectors: (body: V1ListConnectorsRequest, params: RequestParams = {}) =>
      this.request<V1ConnectorList, GoogleRpcStatus>({
        path: `/v1/connectors/list`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags ConnectorsService
     * @name ConnectorsServiceAuthenticateOAuthConnector
     * @request POST:/v1/connectors/oauth
     */
    connectorsServiceAuthenticateOAuthConnector: (body: V1AuthenticateConnectorRequest, params: RequestParams = {}) =>
      this.request<V1Connector, GoogleRpcStatus>({
        path: `/v1/connectors/oauth`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags ConnectorsService
     * @name ConnectorsServiceTestConnection
     * @request POST:/v1/connectors/test
     */
    connectorsServiceTestConnection: (body: V1TestConnectionRequest, params: RequestParams = {}) =>
      this.request<V1TestConnectionResponse, GoogleRpcStatus>({
        path: `/v1/connectors/test`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags DataSourcesService
     * @name DataSourcesServiceListAvailableSources
     * @request GET:/v1/data-sources
     */
    dataSourcesServiceListAvailableSources: (params: RequestParams = {}) =>
      this.request<V1DataSourceList, GoogleRpcStatus>({
        path: `/v1/data-sources`,
        method: "GET",
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags OctopusService
     * @name OctopusServiceSelectAccount
     * @request POST:/v1/data-sources/accounts/setup
     */
    octopusServiceSelectAccount: (body: V1SelectAccountRequest, params: RequestParams = {}) =>
      this.request<V1WorkspaceConnector, GoogleRpcStatus>({
        path: `/v1/data-sources/accounts/setup`,
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
     * @name UsersServiceGetInvitation
     * @request POST:/v1/invitations
     */
    usersServiceGetInvitation: (body: V1GetInvitationRequest, params: RequestParams = {}) =>
      this.request<V1Invitation, GoogleRpcStatus>({
        path: `/v1/invitations`,
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
     * @name UsersServiceSendBatchInvitations
     * @request POST:/v1/invitations/send
     */
    usersServiceSendBatchInvitations: (body: V1BatchInviteMemberRequest, params: RequestParams = {}) =>
      this.request<V1InvitationListResponse, GoogleRpcStatus>({
        path: `/v1/invitations/send`,
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
     * @name UsersServiceGetOrganization
     * @request POST:/v1/organizations
     */
    usersServiceGetOrganization: (body: V1GetOrganizationRequest, params: RequestParams = {}) =>
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
     * @name UsersServiceCreateOrganization
     * @request POST:/v1/organizations/create
     */
    usersServiceCreateOrganization: (body: V1CreateOrganizationRequest, params: RequestParams = {}) =>
      this.request<V1OrganizationResponse, GoogleRpcStatus>({
        path: `/v1/organizations/create`,
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
     * @name UsersServiceCompleteOnboarding
     * @request POST:/v1/organizations/onboarding
     */
    usersServiceCompleteOnboarding: (body: V1CompleteOnboardingRequest, params: RequestParams = {}) =>
      this.request<V1OrganizationResponse, GoogleRpcStatus>({
        path: `/v1/organizations/onboarding`,
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
     * @request POST:/v1/users/refresh
     */
    usersServiceRefreshToken: (body: V1RefreshRequest, params: RequestParams = {}) =>
      this.request<V1AuthenticatedUser, GoogleRpcStatus>({
        path: `/v1/users/refresh`,
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
     * @request POST:/v1/users/signin
     */
    usersServiceSignIn: (body: V1SignInRequest, params: RequestParams = {}) =>
      this.request<V1AuthenticatedUser, GoogleRpcStatus>({
        path: `/v1/users/signin`,
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
     * @request POST:/v1/users/signout
     */
    usersServiceSignOut: (body: V1SignOutRequest, params: RequestParams = {}) =>
      this.request<any, GoogleRpcStatus>({
        path: `/v1/users/signout`,
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
     * @request POST:/v1/users/signup
     */
    usersServiceSignUp: (body: V1SignUpRequest, params: RequestParams = {}) =>
      this.request<V1AuthenticatedUser, GoogleRpcStatus>({
        path: `/v1/users/signup`,
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
     * @request POST:/v1/users/signup/invites
     */
    usersServiceSignUpWithInvite: (body: V1SignUpWithInvitationRequest, params: RequestParams = {}) =>
      this.request<V1AuthenticatedUser, GoogleRpcStatus>({
        path: `/v1/users/signup/invites`,
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
     * @name UsersServiceGetWorkspace
     * @request POST:/v1/workspaces
     */
    usersServiceGetWorkspace: (body: V1GetWorkspaceRequest, params: RequestParams = {}) =>
      this.request<V1WorkspaceResponse, GoogleRpcStatus>({
        path: `/v1/workspaces`,
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
     * @request POST:/v1/workspaces/create
     */
    usersServiceCreateWorkspace: (body: V1CreateWorkspaceRequest, params: RequestParams = {}) =>
      this.request<V1WorkspaceResponse, GoogleRpcStatus>({
        path: `/v1/workspaces/create`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
}
