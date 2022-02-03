import type {FC, ReactNode} from 'react';
import {useEffect} from 'react';
import PropTypes from 'prop-types';
import {push} from 'redux-first-history';
import {ANALYTICS_ROOT_PATH, ONBOARDING_ROOT_PATH} from '../../routes';
import {useDispatch} from 'react-redux';
import {useWorkspace} from '../../hooks/use-workspace';

interface OnboardingGuardProps {
    children: ReactNode;
}

export const OnboardingGuard: FC<OnboardingGuardProps> = (props) => {
    const {children} = props;
    const dispatch = useDispatch();
    const {organization, organizations, attempts} = useWorkspace();

    useEffect(() => {

        const redirectOnboarding = (organization && !organization.onboarded)
            || (attempts > 0 && organizations.length === 0)

        if (redirectOnboarding) {
            dispatch(push(ONBOARDING_ROOT_PATH));
        } else if (organization && organization.onboarded) {
            dispatch(push(ANALYTICS_ROOT_PATH));
        }
    }, [organization, organizations, attempts]);

    return <>{children}</>;
};

OnboardingGuard.propTypes = {
    children: PropTypes.node
};

export default OnboardingGuard;
