import {  takeEvery, put, fork, select } from 'redux-saga/effects';
import * as ActionTypes from './WorkspaceActions.types';
import * as Actions from './WorkspaceActions';
import {
  GoogleRpcStatus, V1Organization,
  V1OrganizationListResponse,
  V1OrganizationResponse,
  V1WorkspaceResponse
} from '../../api';
import { AxiosResponse } from 'axios';
import { Api } from '../configureEffects';
import { Location, useLocation } from 'react-router-dom';
import {ONBOARDING_ROOT_PATH} from '../../routes';
import { OnboardingStep, setOnboardingStep } from '../onboarding';
import {push} from "redux-first-history";

export function* onCreateOrganizationRequest() {
  yield takeEvery(
    ActionTypes.CREATE_ORGANIZATION_REQUEST,
    function* ({ payload }: ActionTypes.ICreateOrganizationRequest) {
      try {
        const response: AxiosResponse<V1OrganizationResponse> =
          yield Api.v1.usersServiceCreateOrganization(payload);
        yield put(Actions.createOrganizationSuccess(response.data));


      } catch (error) {
        console.log(error)
        const rpcError: GoogleRpcStatus = error.response.data;
        yield put(Actions.createOrganizationFailure(rpcError));
      }
    }
  );
}

export function* onCreateOrganizationSuccess() {
  yield takeEvery(
      ActionTypes.CREATE_ORGANIZATION_SUCCESS,
      function* () {
        // Next step if this is onboarding page
        const location: Location = yield select(useLocation);
        if (location.pathname === ONBOARDING_ROOT_PATH) {
          yield put(setOnboardingStep(OnboardingStep.CREATE_WORKSPACE));
        }
      }
  );
}

export function* onListOrganizationsRequest() {
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

export function* onCreateWorkspaceSuccess() {
  yield takeEvery(
      ActionTypes.CREATE_WORKSPACE_SUCCESS,
      function* () {
        // Next step if this is onboarding page
        const location: Location = yield select(useLocation);
        if (location.pathname === ONBOARDING_ROOT_PATH) {
          yield put(setOnboardingStep(OnboardingStep.CONNECT_DATA_SOURCE));
        }
      }
  );
}

export function* onGetOrganizationRequest() {
  yield takeEvery(
      ActionTypes.GET_ORGANIZATION_REQUEST,
      function* ({ payload }: ActionTypes.IGetOrganizationRequest) {
        try {
          const response: AxiosResponse<V1Organization> =
              yield Api.v1.usersServiceGetOrganization(payload);
          yield put(Actions.getOrganizationSuccess(response.data));
        } catch (error) {
          const rpcError: GoogleRpcStatus = error.response.data;
          yield put(Actions.getOrganizationFailure(rpcError));
        }
      }
  );
}

export function* onGetOrganizationFailure() {
  yield takeEvery(
      ActionTypes.GET_ORGANIZATION_FAILURE,
      function* ({ error }: ActionTypes.IGetOrganizationFailure) {
          yield put(push("/error/404"));
      }
  );
}


export const workspaceEffects = [
  fork(onCreateOrganizationRequest),
  fork(onGetOrganizationRequest),
  fork(onGetOrganizationFailure),
  fork(onCreateOrganizationSuccess),
  fork(onCreateWorkspaceSuccess),
  fork(onListOrganizationsRequest),
  fork(onCreateWorkspaceRequest),
];

export default workspaceEffects;
