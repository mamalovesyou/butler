import type { FC, ReactNode } from 'react';
import { useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import { useAuth } from '../../hooks/use-auth';
import { push } from 'redux-first-history';
import { LOGIN_ROUTE_PATH } from '../../routes';
import { useDispatch } from 'react-redux';

interface AuthGuardProps {
  children: ReactNode;
}

export const AuthGuard: FC<AuthGuardProps> = (props) => {
  const { children } = props;
  const dispatch = useDispatch();
  const { isAuthenticated } = useAuth();
  const [checked, setChecked] = useState(false);

  useEffect(() => {
    if (!isAuthenticated) {
      console.log('Not authenticated');
      dispatch(push(LOGIN_ROUTE_PATH));
    } else {
      setChecked(true);
    }
  }, [isAuthenticated]);

  if (!checked) {
    return null;
  }

  // If got here, it means that the redirect did not occur, and that tells us that the user is
  // authenticated / authorized.

  return <>{children}</>;
};

AuthGuard.propTypes = {
  children: PropTypes.node
};
