import React, {useEffect, useState} from "react";
import {useLocation, useRoutes, Navigate} from "react-router-dom";
import DataSources from "../pages/dashboard/data-sources/data-sources";
import {DashboardLayout} from "../components/dashboard/dashboard-layout";
import WorkspaceDetail from "../pages/dashboard/organization/workspace-detail";

import {
    ACCOUNT_ROOT_PATH, ACCOUNT_SETTINGS_PATH, ANALYTICS_ROOT_PATH, CAMPAIGNS_ROOT_PATH, CONNECTORS_DETAIL_PATH,
    DASHBOARD_ROOT_PATH,
    DATA_SOURCES_ROOT_PATH, ERROR_ROOT_PATH, JOIN_ROOT_PATH,
    LOGIN_ROUTE_PATH,
    OAUTH_CALLBACK,
    ONBOARDING_ROOT_PATH,
    ORGANIZATION_ROOT_PATH, ORGANIZATION_TEAM_PATH, ORGANIZATION_WORKSPACES_DETAIL_PATH, ORGANIZATION_WORKSPACES_PATH,
    REGISTER_ROOT_PATH,
} from "./constants";
import Onboarding from "../pages/onboarding";
import {AuthGuard} from "../components/auth/auth-guard";
import Login from "../pages/auth/login";
import Register from "../pages/auth/register";
import {OnboardingGuard} from "../components/onboarding/onboarding-guard";
import OAuthCallback from "../pages/oauth-callback";
import {OrganizationLayout} from "../components/dashboard/organization/organization-layout";
import WorkspacesList from "../components/dashboard/organization/workspaces-list";
import Organization from "../pages/dashboard/organization/organization";
import {TeamMembersTabs} from "../components/dashboard/organization/team-members-tabs";
import {useWorkspace} from "../hooks/use-workspace";
import GeneralSettings from "../pages/dashboard/general-settings";
import UTMAutomation from "../pages/dashboard/utm-automation";
import ErrorPage from "../pages/errors";
import JoinInvitation from "../pages/auth/join-invitation";
import Analytics from "../pages/dashboard/analytics";
import ConnectorDetailsEdit from "../components/dashboard/data-sources/connector-details-edit";
import DataSourceDetail from "../pages/dashboard/data-sources/data-source-detail";

export * from "./constants";

export const AppRoutes: React.FC = () => {

    const location = useLocation();

    const [organization, setOrganization] = useState(null);
    const {organizationId, organizations} = useWorkspace();

    useEffect(() => {
        if (organizationId) {
            setOrganization(organizations[organizationId]);
        }
    }, [organizationId]);

    const element = useRoutes(
        [
            {
                path: OAUTH_CALLBACK,
                element: <OAuthCallback/>,
            },
            {
                path: LOGIN_ROUTE_PATH,
                element: <Login/>,
            },
            {
                path: REGISTER_ROOT_PATH,
                element: <Register/>,
            },
            {
                path: JOIN_ROOT_PATH,
                element: <JoinInvitation/>,
            },
            {
                path: ONBOARDING_ROOT_PATH,
                element: <AuthGuard children={
                    <OnboardingGuard children={<Onboarding/>}/>
                }/>,
            },
            // A route object has the same properties as a <Route>
            // element. The `children` is just an array of child routes.
            {
                path: DASHBOARD_ROOT_PATH,
                element: (
                    <AuthGuard
                        children={<OnboardingGuard children={<DashboardLayout/>}/>}
                    />
                ),
                children: [
                    {path: "", element: <Navigate replace to={ANALYTICS_ROOT_PATH}/>},
                    {path: ANALYTICS_ROOT_PATH, element: <Analytics/>},
                    {path: `${ACCOUNT_ROOT_PATH}/${ACCOUNT_SETTINGS_PATH}`, element: <GeneralSettings/>},
                    {
                        path: ORGANIZATION_ROOT_PATH, element: <OrganizationLayout/>, children: [
                            {
                                path: "", element: <Organization/>, children: [
                                    {path: "", element: <Navigate to={ORGANIZATION_WORKSPACES_PATH}/>},
                                    {path: ORGANIZATION_WORKSPACES_PATH, element: <WorkspacesList/>},
                                    {
                                        path: ORGANIZATION_TEAM_PATH, element: <>{
                                            organization
                                                ? <TeamMembersTabs mode={"organization"} members={organization.members}
                                                                   invitations={organization.invitations}/>
                                                : null
                                        }
                                        </>
                                    }
                                ]
                            },
                            {path: ORGANIZATION_WORKSPACES_DETAIL_PATH, element: <WorkspaceDetail/>},
                        ]
                    },
                    { path: DATA_SOURCES_ROOT_PATH, element: <DataSources/> },
                    {path: CONNECTORS_DETAIL_PATH, element: <DataSourceDetail/>},
                    {
                        path: CAMPAIGNS_ROOT_PATH, element: <UTMAutomation/>,
                        // children: [
                        //     { path: "", element: <Campaigns />},
                        // ]
                    }
                ],
            },
            {
                path: ERROR_ROOT_PATH,
                element: <ErrorPage/>,
                children: [
                    {
                        path: ":code",
                        element: <ErrorPage/>
                    }
                ]
            },
        ],
        location
    );

    return element;
};
