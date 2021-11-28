import { GoogleRpcStatus } from '../../api';

export enum OnboardingStep {
    CREATE_ORGANIZATION = 0,
    CREATE_WORKSPACE = 1,
    CONNECT_DATA_SOURCE = 2
}

export const SET_ONBOARDING_STEP = "SET_ONBOARDING_STEP";

export interface ISetOnboardingStep {
    type: typeof SET_ONBOARDING_STEP;
    payload: OnboardingStep
}

export type OnboardingActionType = ISetOnboardingStep;