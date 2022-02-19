import { FC, useEffect, useState } from 'react';
import { MenuItem, Popover } from '@mui/material';
import { useDispatch } from 'react-redux';
import {useOrganizationsById, useWorkspace} from '../../hooks/use-workspace';
import { V1Workspace } from '../../api';
import {getOrganizationRequest, setCurrentWorkspace} from '../../features/workspace';
import {onGetOrganizationRequest} from "../../features/workspace/WorkspaceEffects";

interface WorkspacePopoverProps {
  anchorEl: null | Element;
  onClose?: () => void;
  onClick?: () => void;
  open?: boolean;
}

interface WorkspaceName {
  id: string;
  organizationid: string;
  name: string;
}

export const WorkspacePopover: FC<WorkspacePopoverProps> = (props) => {
  const { anchorEl, onClose, open, ...other } = props;
  const dispatch = useDispatch();
  const [workspaceNames, setWorkspaceNames] = useState<WorkspaceName[]>([]);
  const { organizationId } = useWorkspace();
  const organizationsById = useOrganizationsById();

  const handleChange = (workspace: WorkspaceName): void => {
    dispatch(
      setCurrentWorkspace({
        organizationId: workspace.organizationid,
        workspaceId: workspace.id
      })
    );
    dispatch(
        getOrganizationRequest({
          organizationId: workspace.organizationid,
        })
    );
    onClose();
  };

  useEffect(() => {

    const names = organizationId
      ? organizationsById[organizationId].workspaces?.map(
          (w: V1Workspace) => w as WorkspaceName
        )
      : [];
    setWorkspaceNames(names);
  }, [organizationId]);

  return (
    <Popover
      anchorEl={anchorEl}
      anchorOrigin={{
        horizontal: 'left',
        vertical: 'bottom'
      }}
      keepMounted
      onClose={onClose}
      open={open}
      PaperProps={{ sx: { width: 248 } }}
      transitionDuration={0}
      {...other}
    >
      {workspaceNames.map((workspace) => (
        <MenuItem key={workspace.id} onClick={() => handleChange(workspace)}>
          {workspace.name}
        </MenuItem>
      ))}
    </Popover>
  );
};
