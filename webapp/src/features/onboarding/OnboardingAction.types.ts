import {V1CompleteOnboardingRequest, V1Organization, V1OrganizationResponse} from '../../api';

export enum OnboardingStep {
  CREATE_ORGANIZATION = 0,
  CREATE_WORKSPACE = 1,
  CONNECT_DATA_SOURCE = 2
}

export const SET_ONBOARDING_STEP = 'SET_ONBOARDING_STEP';

export const COMPLETE_ONBOARDING_REQUEST = 'COMPLETE_ONBOARDING_REQUEST';
export const COMPLETE_ONBOARDING_SUCCESS = 'COMPLETE_ONBOARDING_SUCCESS';

export interface ISetOnboardingStep {
  type: typeof SET_ONBOARDING_STEP;
  payload: OnboardingStep;
}

export interface ICompleteOnboardingRequest {
  type: typeof COMPLETE_ONBOARDING_REQUEST;
  payload: V1CompleteOnboardingRequest;
}

export interface ICompleteOnboardingSuccess {
  type: typeof COMPLETE_ONBOARDING_SUCCESS;
  payload: V1OrganizationResponse;
}

export type OnboardingActionType = ISetOnboardingStep | ICompleteOnboardingRequest | ICompleteOnboardingSuccess;
