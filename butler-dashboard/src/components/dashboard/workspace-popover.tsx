import { FC, useEffect, useState } from 'react';
import { MenuItem, Popover } from '@mui/material';
import { useDispatch } from 'react-redux';
import { useWorkspace } from '../../hooks/use-workspace';
import { V1Workspace } from '../../api';
import { setCurrentWorkspace } from '../../features/workspace';

interface WorkspacePopoverProps {
  anchorEl: null | Element;
  onClose?: () => void;
  onClick?: () => void;
  open?: boolean;
}

interface WorkspaceName {
  ID: string;
  organizationID: string;
  name: string;
}

export const WorkspacePopover: FC<WorkspacePopoverProps> = (props) => {
  const {
    anchorEl, onClose, open, ...other
  } = props;
  const dispatch = useDispatch();
  const [workspaceNames, setworkspaceNames] = useState<WorkspaceName[]>([])
  const { organizations, organizationId } = useWorkspace();

  const handleChange = (workspace: WorkspaceName): void => {
    dispatch(setCurrentWorkspace({
      organizationId: workspace.organizationID,
      workspaceId: workspace.ID,
    }));
    onClose();
  };

  useEffect(() => {
    const names = organizationId
      ? organizations[organizationId].workspaces?.map((w: V1Workspace) => w as WorkspaceName)
      : [];
    setworkspaceNames(names);
  }, [organizationId]);

  return (
    <Popover
      anchorEl={anchorEl}
      anchorOrigin={{
        horizontal: 'left',
        vertical: 'bottom',
      }}
      keepMounted
      onClose={onClose}
      open={open}
      PaperProps={{ sx: { width: 248 } }}
      transitionDuration={0}
      {...other}
    >
      {workspaceNames.map((workspace) => (
        <MenuItem
          key={workspace.ID}
          onClick={() => handleChange(workspace)}
        >
          {workspace.name}
        </MenuItem>
      ))}
    </Popover>
  );
};
