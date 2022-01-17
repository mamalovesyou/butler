import {FC, useEffect, useState} from 'react';
import {Box, Card, Container, Divider, Typography} from '@mui/material';
import {Logo} from '../../components/logo';
import {
    DASHBOARD_ROOT_PATH,
    LOGIN_ROUTE_PATH, NOT_FOUND_ROUTE_PATH,
    REGISTER_ROOT_PATH
} from '../../routes';
import {Link, useLocation} from 'react-router-dom';
import {useAuth} from '../../hooks/use-auth';
import {push} from 'redux-first-history';
import {useDispatch} from 'react-redux';
import {JWTRegister} from '../../components/auth/jwt-register';
import {Api, V1Invitation} from "../../api";
import {AxiosResponse} from "axios";
import JWTRegisterWithInvitation from "../../components/auth/jwt-register-with-invitation";

const Register: FC = () => {
    const dispatch = useDispatch();
    const {isAuthenticated, user} = useAuth();
    const {search} = useLocation();
    const params = new URLSearchParams(search);

    const withInvitation = (): boolean => {
        return params.get("invitationId") !== null && params.get("token") !== null;
    }

    const processInvite = () => {
        const invitationId = params.get("invitationId");
        const invitationToken = params.get("token");
        Api.v1.usersServiceGetInvitation({
            invitationId,
            token: invitationToken
        }).then((response: AxiosResponse<V1Invitation>) => {
            const invitation = response.data;
            console.log("got invitation", invitation);
            if (isAuthenticated && user.email !== invitation.email) dispatch(push(DASHBOARD_ROOT_PATH));
            if (isAuthenticated && user.email === invitation.email) {
                // TODO: JOIN

            }
        }).catch((error) => {
            console.log(error);
            dispatch(push(NOT_FOUND_ROUTE_PATH));
        });
    }

    useEffect(() => {
        if (!withInvitation() && isAuthenticated) {
            console.log("not invitation and authanticated");
            dispatch(push(DASHBOARD_ROOT_PATH));
        } else {
            processInvite();
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
                            { withInvitation() ? <JWTRegisterWithInvitation invitationId={params.get("invitationId")}
                                                                            token={params.get("token")}
                            /> : <JWTRegister /> }
                        </Box>
                        <Divider sx={{my: 3}}/>

                        <Link to={LOGIN_ROUTE_PATH} style={{textDecoration: 'none'}}>
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
