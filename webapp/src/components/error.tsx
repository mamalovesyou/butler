import { FC, useEffect } from 'react';
import { Box, Button, Container, Typography, useMediaQuery } from '@mui/material';
import { useTheme } from '@mui/material/styles';
import {Link, Navigate} from "react-router-dom";
import {DASHBOARD_ROOT_PATH} from "../routes";
import ErrorIcon from '@mui/icons-material/Error';

interface ErrorProps {
    title?: string;
    description?: string;
    redirectPath?: string;
    redirectTitle?: string;
}

export const Error: FC<ErrorProps> = (props) => {
    const theme = useTheme();
    const mobileDevice = useMediaQuery(theme.breakpoints.down('sm'));

    return <Box
        component="main"
        sx={{
            alignItems: 'center',
            backgroundColor: 'background.paper',
            display: 'flex',
            flexGrow: 1,
            py: '80px'
        }}
    >
        <Container maxWidth="lg">
            <Box
                sx={{
                    display: 'flex',
                    justifyContent: 'center',
                    mt: 6
                }}
            >
                <Box
                    sx={{
                        height: 'auto',
                        maxWidth: '100%',
                    }}
                >
                    <ErrorIcon  color="error" sx={{ width: 300, height: 300 }}/>
                </Box>
            </Box>
            <Typography
                align="center"
                variant={mobileDevice ? 'h4' : 'h1'}
            >
                {props.title}
            </Typography>
            <Typography
                align="center"
                color="textSecondary"
                sx={{ mt: 0.5 }}
                variant="subtitle2"
            >
                {props.description}
            </Typography>
            <Box
                sx={{
                    display: 'flex',
                    justifyContent: 'center',
                    mt: 6
                }}
            >
                <Button
                    component={Link}
                    to={props.redirectPath}
                    replace
                    variant="outlined"
                >
                    {props.redirectTitle}
                </Button>
            </Box>
        </Container>
    </Box>
};

Error.defaultProps = {
    title: "Something wrong happened!",
    description: "",
    redirectPath: DASHBOARD_ROOT_PATH,
    redirectTitle: "Back to Dashboard"
}

export default Error;

