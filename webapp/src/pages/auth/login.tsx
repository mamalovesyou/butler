import { FC, useEffect } from 'react';
import { Box, Card, Container, Divider, Typography } from '@mui/material';
import { Link } from 'react-router-dom';
import { push } from 'redux-first-history';
import { useDispatch } from 'react-redux';
import { JWTLogin } from '../../components/auth/jwt-login';
import { Logo } from '../../components/logo';
import { DASHBOARD_ROOT_PATH, REGISTER_ROOT_PATH } from '../../routes';
import { useAuth } from '../../hooks/use-auth';

const Login: FC = () => {
  const dispatch = useDispatch();
  const { isAuthenticated } = useAuth();

  useEffect(() => {
    if (isAuthenticated) {
      dispatch(push(DASHBOARD_ROOT_PATH));
    }
  }, []);

  return (
    <Box
      component="main"
      sx={{
        backgroundColor: 'background.default',
        display: 'flex',
        flexDirection: 'column',
        minHeight: '100vh'
      }}
    >
      <Container
        maxWidth="sm"
        sx={{
          py: {
            xs: '60px',
            md: '120px'
          }
        }}
      >
        <Card elevation={16} sx={{ p: 4 }}>
          <Box
            sx={{
              alignItems: 'center',
              display: 'flex',
              flexDirection: 'column',
              justifyContent: 'center'
            }}
          >
            <Link to={DASHBOARD_ROOT_PATH}>
              <Logo
                sx={{
                  height: 40,
                  width: 40
                }}
              />
            </Link>
            <Typography variant="h4">Login</Typography>
            <Typography color="textSecondary" sx={{ mt: 2 }} variant="body2">
              Welcome back !
            </Typography>
          </Box>
          <Box
            sx={{
              flexGrow: 1,
              mt: 3
            }}
          >
            <JWTLogin />
          </Box>
          <Divider sx={{ my: 3 }} />

          <Link to={REGISTER_ROOT_PATH} style={{ textDecoration: 'none' }}>
            <Typography color="textSecondary" variant="body2">
              Create new account
            </Typography>
          </Link>
        </Card>
      </Container>
    </Box>
  );
};

export default Login;
