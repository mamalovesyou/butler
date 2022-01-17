import { FC, useEffect } from 'react';
import { Box, Card, Container, Divider, Typography } from '@mui/material';
import { Logo } from '../../components/logo';
import {
  DASHBOARD_ROOT_PATH,
  LOGIN_ROUTE_PATH,
  REGISTER_ROOT_PATH
} from '../../routes';
import { Link } from 'react-router-dom';
import { useAuth } from '../../hooks/use-auth';
import { push } from 'redux-first-history';
import { useDispatch } from 'react-redux';
import { JWTRegister } from '../../components/auth/jwt-register';

const RegisterWithInvitation: FC = () => {
  const dispatch = useDispatch();
  const { isAuthenticated } = useAuth();

  useEffect(() => {
    if (isAuthenticated) {
      dispatch(push(DASHBOARD_ROOT_PATH));
    }
  }, []);

  return (
    <>
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
              <Typography variant="h4">Register</Typography>
              <Typography color="textSecondary" sx={{ mt: 2 }} variant="body2">
                Welcome to the HeyButler familly !
              </Typography>
            </Box>
            <Box
              sx={{
                flexGrow: 1,
                mt: 3
              }}
            >
              <JWTRegister />
            </Box>
            <Divider sx={{ my: 3 }} />

            <Link to={LOGIN_ROUTE_PATH} style={{ textDecoration: 'none' }}>
              <Typography color="textSecondary" variant="body2">
                Signin with existing account
              </Typography>
            </Link>
          </Card>
        </Container>
      </Box>
    </>
  );
};

export default Register;
