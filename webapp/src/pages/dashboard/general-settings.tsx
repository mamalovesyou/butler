import { FC } from 'react';
import { Box, Container, Divider, Typography } from '@mui/material';
import {AccountGeneralSettings} from "../../components/dashboard/account/account-general-settings";

const GeneralSettings: FC = () => {

  return (
    <>
      <Box
        component="main"
        sx={{
          flexGrow: 1,
          py: 8
        }}
      >
        <Container maxWidth="md">
          <Typography variant="h4">My Account</Typography>
          <Divider sx={{ mb: 3 }} />
          <AccountGeneralSettings />
        </Container>
      </Box>
    </>
  );
};

export default GeneralSettings;
