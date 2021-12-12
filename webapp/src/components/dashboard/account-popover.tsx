import type { FC } from 'react';
import PropTypes from 'prop-types';
import {
  Box,
  Divider,
  ListItemIcon,
  ListItemText,
  MenuItem,
  Popover,
  Typography
} from '@mui/material';
import LogoutIcon from '@mui/icons-material/Logout';
import { Cog as CogIcon } from '../../icons/cog';
import { useDispatch } from 'react-redux';
import { ACCOUNT_ROOT_PATH } from '../../routes';
import { Link } from 'react-router-dom';
import { logout } from '../../features/auth';
import { UserAvatar } from '../user-avatar';
import { useAuth } from '../../hooks/use-auth';

interface AccountPopoverProps {
  anchorEl: null | Element;
  onClose?: () => void;
  open?: boolean;
}

export const AccountPopover: FC<AccountPopoverProps> = (props) => {
  const { anchorEl, onClose, open, ...other } = props;
  const dispatch = useDispatch();

  const { user } = useAuth();
  const fullName = `${user?.firstName} ${user?.lastName}`;

  const handleLogout = async (): Promise<void> => {
    dispatch(logout());
  };

  return (
    <Popover
      anchorEl={anchorEl}
      anchorOrigin={{
        horizontal: 'center',
        vertical: 'bottom'
      }}
      keepMounted
      onClose={onClose}
      open={open}
      PaperProps={{ sx: { width: 300 } }}
      transitionDuration={0}
      {...other}
    >
      <Box
        sx={{
          alignItems: 'center',
          p: 2,
          display: 'flex'
        }}
      >
        <UserAvatar name={fullName} />
        <Box
          sx={{
            ml: 1
          }}
        >
          <Typography variant="body1">{user.firstName}</Typography>
        </Box>
      </Box>
      <Divider />
      <Box sx={{ my: 1 }}>
        <MenuItem to={ACCOUNT_ROOT_PATH} component={Link}>
          <ListItemIcon>
            <CogIcon fontSize="small" />
          </ListItemIcon>
          <ListItemText
            primary={<Typography variant="body1">Settings</Typography>}
          />
        </MenuItem>
        <Divider />
        <MenuItem onClick={handleLogout}>
          <ListItemIcon>
            <LogoutIcon fontSize="small" />
          </ListItemIcon>
          <ListItemText
            primary={<Typography variant="body1">Logout</Typography>}
          />
        </MenuItem>
      </Box>
    </Popover>
  );
};

AccountPopover.propTypes = {
  anchorEl: PropTypes.any,
  onClose: PropTypes.func,
  open: PropTypes.bool
};
