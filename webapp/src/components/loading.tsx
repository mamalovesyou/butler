import {Box, CircularProgress, Container} from '@mui/material';
import { Logo } from '../components/logo';

export const Loading = () => {

    return (
        <Box
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
            <Box sx={{p: 2}}> <CircularProgress /> </Box>
        </Box>
    );
};

export default Loading;
