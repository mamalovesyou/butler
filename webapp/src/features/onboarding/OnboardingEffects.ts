import { takeEvery, put, fork, select } from 'redux-saga/effects';
import * as ActionTypes from './OnboardingAction.types';
import * as Actions from './OnboardingActions';
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
import {completeOnboardingSuccess} from "./OnboardingActions";

export function* onCompleteOnboarding() {
  yield takeEvery(
    ActionTypes.COMPLETE_ONBOARDING_REQUEST,
    function* ({ payload }: ActionTypes.ICompleteOnboardingRequest) {
      try {
        const response: AxiosResponse<V1OrganizationResponse> =
          yield Api.v1.usersServiceCompleteOnboarding(payload);
          yield put(completeOnboardingSuccess(response.data));
      } catch (error) {
        // TODO: Add generic error request failure
        console.log(error)
        const rpcError: GoogleRpcStatus = error.response?.data;
      }
    }
  );
}

export const onboardingEffects = [
  fork(onCompleteOnboarding),
];

export default onboardingEffects;
