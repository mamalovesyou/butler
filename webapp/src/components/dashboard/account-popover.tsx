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
import {ACCOUNT_ROOT_PATH, ACCOUNT_SETTINGS_PATH, LOGOUT_ROOT_PATH} from '../../routes';
import { Link } from 'react-router-dom';
import {push} from "redux-first-history";

interface AccountPopoverProps {
  anchorEl: null | Element;
  onClose?: () => void;
  open?: boolean;
}

export const AccountPopover: FC<AccountPopoverProps> = (props) => {
  const { anchorEl, onClose, open, ...other } = props;
  const dispatch = useDispatch();

  const handleLogout = async (): Promise<void> => {
    dispatch(push(LOGOUT_ROOT_PATH));
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
      <Box sx={{ my: 1 }}>
        <MenuItem to={`${ACCOUNT_ROOT_PATH}/${ACCOUNT_SETTINGS_PATH}`} component={Link}>
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
