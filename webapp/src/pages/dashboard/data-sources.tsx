import { FC, useEffect } from 'react';
import { Box, Container, Divider, Typography } from '@mui/material';
import { ConnectorsCatalogList } from '../../components/dashboard/data-sources/connectors-catalog-list';
import { listCatalogConnectorsRequest } from '../../features/connectors';
import { useDispatch } from 'react-redux';
import {AccountGeneralSettings} from "../../components/dashboard/account/account-general-settings";

const DataSources: FC = () => {

  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(listCatalogConnectorsRequest());
  }, []);

  return (
      <Box
          component="main"
          sx={{
              flexGrow: 1,
              py: 8
          }}
      >
          <Container maxWidth="md">
              <Typography variant="h4">Data Sources</Typography>
              <Divider sx={{ mb: 3 }} />
              <ConnectorsCatalogList />
          </Container>
      </Box>
  );
};

export default DataSources;
