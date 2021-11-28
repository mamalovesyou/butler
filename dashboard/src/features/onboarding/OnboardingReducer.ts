
import * as OnboardingActions from './OnboardingAction.types';

export type OnboardingStateType = {
    activeStep: OnboardingActions.OnboardingStep;
}

const initialState: OnboardingStateType = {
    activeStep: OnboardingActions.OnboardingStep.CREATE_ORGANIZATION,
}

const onboardingReducer = (state: OnboardingStateType = initialState, action: OnboardingActions.OnboardingActionType) => {
    switch (action.type) {
        case OnboardingActions.SET_ONBOARDING_STEP:
            return {
                ...state,
                activeStep: action.payload
            }
        default:
            return state;
    }
}

export default onboardingReducer;