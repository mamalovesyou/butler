import { takeEvery, put, fork, select } from 'redux-saga/effects';
import * as ActionTypes from './WorkspaceActions.types';
import * as Actions from './WorkspaceActions';
import {
  GoogleRpcStatus,
  V1OrganizationListResponse,
  V1OrganizationResponse,
  V1AuthenticatedUser,
  V1WorkspaceResponse
} from '../../api';
import { AxiosResponse } from 'axios';
import { Api } from '../configureEffects';
import { Location, useLocation } from 'react-router-dom';
import { ONBOARDING_ROOT_PATH } from '../../routes';
import { OnboardingStep, setOnboardingStep } from '../onboarding';

export function* onCreateOrganizationRequest() {
  yield takeEvery(
    ActionTypes.CREATE_ORGANIZATION_REQUEST,
    function* ({ payload }: ActionTypes.ICreateOrganizationRequest) {
      try {
        const response: AxiosResponse<V1OrganizationResponse> =
          yield Api.v1.usersServiceCreateOrganization(payload);
        yield put(Actions.createOrganizationSuccess(response.data));

        // Next step if this is onboarding page
        const location: Location = yield select(useLocation);
        if (location.pathname === ONBOARDING_ROOT_PATH) {
          yield put(setOnboardingStep(OnboardingStep.CREATE_WORKSPACE));
        }
      } catch (error) {
        console.log(error)
        const rpcError: GoogleRpcStatus = error.response.data;
        yield put(Actions.createOrganizationFailure(rpcError));
      }
    }
  );
}

export function* onListWorkspacesRequest() {
  yield takeEvery(ActionTypes.LIST_ORGANIZATIONS_REQUEST, function* () {
    try {
      const response: AxiosResponse<V1OrganizationListResponse> =
        yield Api.v1.usersServiceListOrganizations();
      yield put(Actions.listOrganizationsSuccess(response.data));
    } catch (error) {
      const rpcError: GoogleRpcStatus = error.response.data;
      yield put(Actions.listOrganizationsFailure(rpcError));
    }
  });
}

export function* onCreateWorkspaceRequest() {
  yield takeEvery(
    ActionTypes.CREATE_WORKSPACE_REQUEST,
    function* ({ payload }: ActionTypes.ICreateWorkspaceRequest) {
      try {
        const response: AxiosResponse<V1WorkspaceResponse> =
          yield Api.v1.usersServiceCreateWorkspace(payload);
        yield put(Actions.createWorkspaceSuccess(response.data));

        // Next step if this is onboarding page
        const location: Location = yield select(useLocation);
        if (location.pathname === ONBOARDING_ROOT_PATH) {
          yield put(setOnboardingStep(OnboardingStep.CONNECT_DATA_SOURCE));
        }
      } catch (error) {
        const rpcError: GoogleRpcStatus = error.response.data;
        yield put(Actions.createWorkspaceFailure(rpcError));
      }
    }
  );
}

export function* onAddWorkspaceMembersRequest() {
  yield takeEvery(
    ActionTypes.ADD_WORKSPACE_MEMBER_REQUEST,
    function* ({ payload }: ActionTypes.IAddWorkspaceMembersRequest) {
      try {
        const response: AxiosResponse<V1AuthenticatedUser> =
          yield Api.v1.usersServiceInviteWorkspaceMember(payload);
        yield put(Actions.addWorkspaceMembersSuccess(response.data));
      } catch (error) {
        const rpcError: GoogleRpcStatus = error.response.data;
        yield put(Actions.addWorkspaceMembersFailure(rpcError));
      }
    }
  );
}

export const workspaceEffects = [
  fork(onCreateOrganizationRequest),
  fork(onListWorkspacesRequest),
  fork(onCreateWorkspaceRequest),
  fork(onAddWorkspaceMembersRequest)
];

export default workspaceEffects;
