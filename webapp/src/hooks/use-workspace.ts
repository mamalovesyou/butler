import {useSelector} from "react-redux";
import {V1Organization, V1Workspace} from "../api";
import {ArrayToObject} from "../utils/array";
import {RootState} from "../features";

export const useWorkspace = () =>
    useSelector((state: RootState) => state.workspace);

export const useOrganizationsById = () =>
    useSelector((state: RootState) => ArrayToObject(state.workspace.organizations, 'id'));

export const useCurrentOrganization = () =>
    useSelector((state: RootState) =>
        state.workspace.organizationId ? state.workspace.organizations[state.workspace.organizationId] : null)

export const useCurrentWorkspace = (): {
    workspace: V1Workspace;
    organization: V1Organization;
} =>
    useSelector((state: RootState) => {
        const { workspaceId, organization} = state.workspace;
        return {
            organization,
            workspace: organization?.workspaces?.find((ws: V1Workspace) => ws.id === workspaceId),
        };
    });
