import {FC, useEffect, useState} from 'react';
import {Box, Card, Container, Divider, Typography} from '@mui/material';
import {Logo} from '../../components/logo';
import {
    DASHBOARD_ROOT_PATH,
    LOGIN_ROOT_PATH, NOT_FOUND_ROUTE_PATH,
    REGISTER_ROOT_PATH
} from '../../routes';
import {Link, useLocation} from 'react-router-dom';
import {useAuth} from '../../hooks/use-auth';
import {push} from 'redux-first-history';
import {useDispatch} from 'react-redux';
import JWTRegisterWithInvitation from "../../components/auth/jwt-register-with-invitation";
import {Api, V1Invitation} from "../../api";
import {AxiosResponse} from "axios";

const RegisterWithInvitation: FC = () => {
    const dispatch = useDispatch();
    const {search} = useLocation();
    const params = new URLSearchParams(search);
    const {isAuthenticated, user} = useAuth();

    useEffect(() => {

        Api.v1.usersServiceGetInvitation({
            invitationId: params.get("invitationId"),
            token: params.get("token")
        }).then((response: AxiosResponse<V1Invitation>) => {
            const invitation = response.data;
            if (isAuthenticated && user.email !== invitation.email) dispatch(push(DASHBOARD_ROOT_PATH));
            if (isAuthenticated && user.email === invitation.email) {
                // TODO: JOIN
            }
        }).catch((error) => {
            console.log(error);
            dispatch(push(NOT_FOUND_ROUTE_PATH));
        });

        if (isAuthenticated) {

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
                    <Card elevation={16} sx={{p: 4}}>
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
                            <Typography color="textSecondary" sx={{mt: 2}} variant="body2">
                                Welcome to the HeyButler familly !
                            </Typography>
                        </Box>
                        <Box
                            sx={{
                                flexGrow: 1,
                                mt: 3
                            }}
                        >
                            <JWTRegisterWithInvitation invitationId={params.get("invitationId")}
                                                       token={params.get("token")}/>
                        </Box>
                        <Divider sx={{my: 3}}/>

                        <Link to={LOGIN_ROOT_PATH} style={{textDecoration: 'none'}}>
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

export default RegisterWithInvitation;
