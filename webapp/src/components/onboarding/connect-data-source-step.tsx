import type {FC} from 'react';
import {
    Box,
    Button,
    Typography
} from '@mui/material';
import {useDispatch} from 'react-redux';
import {push} from 'redux-first-history';
import {DASHBOARD_ROOT_PATH, DATA_SOURCES_ROOT_PATH} from '../../routes';
import {completeOnboardingRequest} from "../../features/onboarding";
import {useWorkspace} from "../../hooks/use-workspace";

export const ConnectDataSourceStep: FC = (props) => {
    const {...other} = props;
    const dispatch = useDispatch();

    const { organizationId } = useWorkspace();

    const handleCreateSource = () => {
        dispatch(completeOnboardingRequest(organizationId));
        dispatch(push(DATA_SOURCES_ROOT_PATH));
    };

    const handleSkip = () => {
        dispatch(completeOnboardingRequest(organizationId));
        dispatch(push(DASHBOARD_ROOT_PATH));
    };

    return (
        <div {...other}>
            <Typography variant="h6">Connect a Data Source</Typography>
            <Typography color="textSecondary" sx={{mt: 2}} variant="body2">
                You can think of data sources as channels
            </Typography>
            <Box sx={{pt: 2, display: 'flex', flexDirection: 'row', justifyContent: 'space-between'}}>
                <Button variant="contained" color="primary" onClick={handleCreateSource}>
                    Connect source
                </Button>
                <Button variant="outlined" color="secondary" onClick={handleSkip}>
                    Skip
                </Button>
            </Box>
        </div>
    );
};
