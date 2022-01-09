import { useSelector } from "react-redux";
import { V1Organization, V1Workspace } from "../api";
import { ArrayToObject } from "../utils/array";
import { RootState } from "../features";

export const useWorkspace = () =>
  useSelector((state: RootState) => state.workspace);

export const useCurrentWorkspace = (): {
  workspace: V1Workspace;
  organization: V1Organization;
} =>
  useSelector((state: RootState) => {
    const { organizationId, workspaceId, organizations } = state.workspace;
    return {
      organization: organizationId ? organizations[organizationId] : null,
      workspace:
        workspaceId && organizationId
          ? ArrayToObject(organizations[organizationId].workspaces)[
              workspaceId
            ]
          : null,
    };
  });
