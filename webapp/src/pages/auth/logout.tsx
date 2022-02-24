import {FC, useEffect} from 'react';
import {useDispatch} from 'react-redux';
import {logout} from "../../features/auth";
import {Box, Button, Container, Typography} from "@mui/material";
import {Link} from "react-router-dom";
import {LOGIN_ROOT_PATH} from "../../routes";

const Logout: FC = () => {
    const dispatch = useDispatch();

    useEffect(() => {
        console.log('logout');
        dispatch(logout());
    }, []);

    return <Box
        component="main"
        sx={{
            backgroundColor: 'background.default',
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            justifyContent: 'center',
            minHeight: '100%',
            minWidth: '100%'
        }}
    >
        <Typography variant="h4">Sorry to see you go :(</Typography>
        <Box sx={{p: 2}}>
            <Button
                component={Link}
                to={LOGIN_ROOT_PATH}
                type="submit"
                variant="contained"
            >
                Log In
            </Button>
        </Box>
    </Box>;
};

export default Logout;
