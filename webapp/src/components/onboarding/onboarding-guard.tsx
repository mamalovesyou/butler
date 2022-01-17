import type { FC, ReactNode } from 'react';
import { useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import { push } from 'redux-first-history';
import { DASHBOARD_ROOT_PATH, ONBOARDING_ROOT_PATH } from '../../routes';
import { useDispatch } from 'react-redux';
import { useWorkspace } from '../../hooks/use-workspace';

interface OnboardingGuardProps {
  children: ReactNode;
}

export const OnboardingGuard: FC<OnboardingGuardProps> = (props) => {
  const { children } = props;
  const dispatch = useDispatch();
  const { organizationId, organizations } = useWorkspace();

  useEffect(() => {
    if (organizationId) {
      const currentOrg = organizations[organizationId]
      if (!currentOrg.onboarded) {
        dispatch(push(ONBOARDING_ROOT_PATH));
      }
    }
  }, [organizationId]);

  return <>{children}</>;
};

OnboardingGuard.propTypes = {
  children: PropTypes.node
};

export default OnboardingGuard;
