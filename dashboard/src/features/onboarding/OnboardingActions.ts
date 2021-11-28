import { OnboardingStep } from '.';
import * as ActionTypes from './OnboardingAction.types';

export const setOnboardingStep = (step: OnboardingStep) => ({
    type: ActionTypes.SET_ONBOARDING_STEP,
    payload: step
});