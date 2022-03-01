import {FC} from 'react';
import {Box, Container, Divider, Typography} from '@mui/material';
import {Routes, Route, Outlet} from "react-router-dom";

const URLBuilder: FC = () => {
    console.log('url buidler')
    return (
        <Box
            component="main"
            sx={{
                flexGrow: 1,
                py: 8
            }}
        >
            <Container maxWidth="md">
                <Box
                    sx={{
                        paddingBottom: 2,
                        alignItems: 'center',
                        display: 'flex',
                        justifyContent: 'space-between',
                    }}
                >
                    <Typography variant="h4">URL Builder</Typography>
                </Box>
                <Divider sx={{mb: 3}}/>
                <Outlet />
            </Container>
        </Box>
    );
};

export default URLBuilder;
