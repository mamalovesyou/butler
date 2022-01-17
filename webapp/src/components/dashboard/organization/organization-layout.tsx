import type { FC } from 'react';
import { Box } from '@mui/material';
import { Outlet } from 'react-router-dom';

export const OrganizationLayout: FC = () => {
  return (
      <Box
          sx={{
              display: 'flex',
              flex: '1 1 auto',
              flexDirection: 'column',
              width: '100%'
          }}
      >
          <Outlet />
      </Box>
  );
};
