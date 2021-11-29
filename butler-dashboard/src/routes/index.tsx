import React from 'react';
import {
  useLocation,
  useRoutes,
} from 'react-router-dom';
import DataSources from '../pages/dashboard/data-sources';
import { DashboardLayout } from '../components/dashboard/dashboard-layout';
import NotFound from '../pages/404';
import Account from '../pages/dashboard/account';
import {
  ACCOUNT_ROOT_PATH, DASHBOARD_ROOT_PATH, DATA_SOURCES_ROOT_PATH, LOGIN_ROUTE_PATH, OAUTH_CALLBACK, ONBOARDING_ROOT_PATH, REGISTER_ROOT_PATH,
} from './constants';
import Onboarding from '../pages/onboarding';
import { AuthGuard } from '../components/auth/auth-guard';
import Login from '../pages/auth/login';
import Register from '../pages/auth/register';
import { OnboardingGuard } from '../components/onboarding/onboarding-guard';
import OAuthCallback from '../pages/oauth-callback';

export * from './constants';

export const AppRoutes: React.FC = () => {
  const location = useLocation();
  const element = useRoutes([
    {
      path: OAUTH_CALLBACK,
      element: <OAuthCallback />,
    },
    {
      path: LOGIN_ROUTE_PATH,
      element: <Login />,
    },
    {
      path: REGISTER_ROOT_PATH,
      element: <Register />,
    },
    {
      path: ONBOARDING_ROOT_PATH,
      element: <AuthGuard children={<Onboarding />} />,
    },
    // A route object has the same properties as a <Route>
    // element. The `children` is just an array of child routes.
    {
      path: DASHBOARD_ROOT_PATH,
      element: <AuthGuard children={<OnboardingGuard children={<DashboardLayout />} />} />,
      children: [
        { path: DASHBOARD_ROOT_PATH, element: <Account /> },
        { path: ACCOUNT_ROOT_PATH, element: <Account /> },
        { path: DATA_SOURCES_ROOT_PATH, element: <DataSources /> },
      ],
    },
    {
      path: '*',
      element: <NotFound />,
    },
  ], location);

  return element;
};
