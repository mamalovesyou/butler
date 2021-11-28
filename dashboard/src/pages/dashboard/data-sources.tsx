import type { FC } from 'react';
import { Box, Container, Divider, Typography } from '@mui/material';

const DataSources: FC = () => {

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
          <Typography variant="h4">
            Data Sources
          </Typography>
          <Divider sx={{ mb: 3 }} />
          <Typography variant="h5">Comming soon</Typography>
        </Container>
      </Box>
    </>
  );
};

export default DataSources;
