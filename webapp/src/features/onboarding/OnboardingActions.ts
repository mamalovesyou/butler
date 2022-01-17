import { OnboardingStep } from '.';
import * as ActionTypes from './OnboardingAction.types';
import {V1OrganizationResponse} from "../../api";

export const setOnboardingStep = (step: OnboardingStep) => ({
  type: ActionTypes.SET_ONBOARDING_STEP,
  payload: step
});

export const completeOnboardingRequest = (organizationId: string) => ({
  type: ActionTypes.COMPLETE_ONBOARDING_REQUEST,
  payload: { organizationId }
});

export const completeOnboardingSuccess = (organization: V1OrganizationResponse) => ({
  type: ActionTypes.COMPLETE_ONBOARDING_SUCCESS,
  payload: organization
});
