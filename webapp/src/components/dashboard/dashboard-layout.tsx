import { useState } from 'react';
import type { FC } from 'react';
import { styled } from '@mui/material/styles';
import { DashboardNavbar } from './dashboard-navbar';
import { DashboardSidebar } from './dashboard-sidebar';
import {Box} from '@mui/material';
import { Outlet } from 'react-router-dom';
import {DASHBOARD_NAVBAR_HEIGHT, DASHBOARD_SIDEBAR_WIDTH} from "../../theme/constants";

const DashboardLayoutRoot = styled('div')(({ theme }) => ({
  display: 'flex',
  flex: '1 1 auto',
  maxWidth: '100%',
  paddingTop: 64,
  [theme.breakpoints.up('lg')]: {
    paddingLeft: DASHBOARD_SIDEBAR_WIDTH
  }
}));

export const DashboardLayout: FC = () => {
  const [isSidebarOpen, setIsSidebarOpen] = useState<boolean>(false);

  return (
    <>
      <DashboardLayoutRoot>
        <Box
          sx={{
            display: 'flex',
            flex: '1 1 auto',
            flexDirection: 'column',
            width: '100%',
          }}
        >
            <Outlet />
        </Box>
      </DashboardLayoutRoot>
      <DashboardNavbar onOpenSidebar={(): void => setIsSidebarOpen(true)} />
      <DashboardSidebar
        onClose={(): void => setIsSidebarOpen(false)}
        open={isSidebarOpen}
      />
    </>
  );
};
