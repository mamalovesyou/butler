import {FC, useEffect, useState} from 'react';
import {Box, Card, Container, Divider, Typography} from '@mui/material';
import {Link, useLocation} from 'react-router-dom';
import {push} from 'redux-first-history';
import {useDispatch} from 'react-redux';
import {Logo} from '../../components/logo';
import {DASHBOARD_ROOT_PATH, REGISTER_ROOT_PATH} from '../../routes';
import {useAuth} from '../../hooks/use-auth';
import {Api, V1Invitation} from "../../api";
import {AxiosResponse} from "axios";
import Button from "@mui/material/Button";
import JWTRegisterWithInvitation from "../../components/auth/jwt-register-with-invitation";

const JoinInvitation: FC = () => {
    const dispatch = useDispatch();
    const {isAuthenticated, user} = useAuth();
    const [invitation, setInvitation] = useState<V1Invitation>(null);
    const {search} = useLocation();


    // Update invitation on query params
    useEffect(() => {
        const params = new URLSearchParams(search);
        if (params.has("invitationId") && params.has("token")) {
            const body = {invitationId: params.get("invitationId"), token: params.get("token")};
            console.log(body)
            Api.v1.usersServiceGetInvitation(body)
                .then((response: AxiosResponse<V1Invitation>) => setInvitation(response.data))
                .catch((error) => {
                    console.log(error);
                    dispatch(push("/error/403"));
                });

        } else {
            dispatch(push("/error/404"));
        }
    }, [search])

    return (
        <Box
            component="main"
            sx={{
                backgroundColor: 'background.default',
                display: 'flex',
                flexDirection: 'column',
                minHeight: '100%',
                minWidth: '100%'
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
                        <Typography variant="h4">Join {invitation?.organization?.name}</Typography>
                        <Typography color="textSecondary" sx={{mt: 2}} variant="body2">
                            You have been invited to join the workspace {invitation?.workspace?.name}
                        </Typography>
                    </Box>
                    <Box
                        sx={{
                            flexGrow: 1,
                            mt: 3
                        }}
                    >
                        { isAuthenticated ? <Button>Join now</Button> : <JWTRegisterWithInvitation invitation={invitation} /> }
                    </Box>
                    <Divider sx={{my: 3}}/>
                </Card>
            </Container>
        </Box>
    );
};

export default JoinInvitation;
