import type {FC, ReactNode} from 'react';
import {useEffect} from 'react';
import PropTypes from 'prop-types';
import {push} from 'redux-first-history';
import {ONBOARDING_ROOT_PATH} from '../../routes';
import {useDispatch} from 'react-redux';
import {useWorkspace} from '../../hooks/use-workspace';
import {Loading} from "../loading";
import {listOrganizationsRequest} from "../../features/workspace";

interface OnboardingGuardProps {
    children: ReactNode;
}

export const OnboardingGuard: FC<OnboardingGuardProps> = (props) => {

    const {children} = props;
    const dispatch = useDispatch();
    const {selected, loading} = useWorkspace();

    useEffect(() => {
        console.log(loading, selected);
        if (!loading && (selected.organization === null || selected.workspace === null)) {
            dispatch(push(ONBOARDING_ROOT_PATH));
        }
    }, [loading]);

    useEffect(() => {
        if (selected.organization && !selected.organization.onboarded) {
            dispatch(push(ONBOARDING_ROOT_PATH));
        }
    }, [selected])

    return loading ? <Loading /> : <>{children}</>;
};

OnboardingGuard.propTypes = {
    children: PropTypes.node
};

export default OnboardingGuard;
