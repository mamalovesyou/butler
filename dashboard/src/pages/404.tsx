import type { FC } from 'react';
import {
  Box, Button, Container, Typography, useMediaQuery,
} from '@mui/material';
import { useTheme } from '@mui/material/styles';
import { Link } from 'react-router-dom';
import { DASHBOARD_ROOT_PATH } from '../routes/constants';

const NotFound: FC = () => {
  const theme = useTheme();
  const mobileDevice = useMediaQuery(theme.breakpoints.down('sm'));

  return (
    <Box
      component="main"
      sx={{
        alignItems: 'center',
        backgroundColor: 'background.paper',
        display: 'flex',
        flexGrow: 1,
        py: '80px',
      }}
    >
      <Container maxWidth="lg">
        <Typography
          align="center"
          variant={mobileDevice ? 'h4' : 'h1'}
        >
          404: The page you are looking for isn’t here
        </Typography>
        <Typography
          align="center"
          color="textSecondary"
          sx={{ mt: 0.5 }}
          variant="subtitle2"
        >
          You either tried some shady route or you
          came here by mistake. Whichever it is, try using the
          navigation.
        </Typography>
        <Box
          sx={{
            display: 'flex',
            justifyContent: 'center',
            mt: 6,
          }}
        >
          <Box
            alt="Under development"
            component="img"
            src={`/static/error/error404_${theme.palette.mode}.svg`}
            sx={{
              height: 'auto',
              maxWidth: '100%',
              width: 400,
            }}
          />
        </Box>
        <Box
          sx={{
            display: 'flex',
            justifyContent: 'center',
            mt: 6,
          }}
        >
          <Button
            component={Link}
            to={DASHBOARD_ROOT_PATH}
            variant="outlined"
          >
            Back to Dashboard
          </Button>
        </Box>
      </Container>
    </Box>
  );
};

export default NotFound;
