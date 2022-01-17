import type {FC} from 'react';
import {
    Avatar,
    Box,
    Button,
    Card,
    CardContent,
    Grid,
    Typography
} from '@mui/material';
import {useDispatch} from 'react-redux';
import {push} from 'redux-first-history';
import {DASHBOARD_ROOT_PATH} from '../../routes';
import {ConnectorsCatalogList} from '../dashboard/data-sources/connectors-catalog-list';
import {completeOnboardingRequest, setOnboardingStep} from "../../features/onboarding";
import {useDataSources} from "../../hooks/use-connectors";
import {useWorkspace} from "../../hooks/use-workspace";

export const ConnectDataSourceStep: FC = (props) => {
    const {...other} = props;
    const dispatch = useDispatch();

    const { organizationId } = useWorkspace();
    const dataSources = useDataSources();

    const handleComplete = () => {
        dispatch(completeOnboardingRequest(organizationId));
        dispatch(push(DASHBOARD_ROOT_PATH));
    };

    return (
        <div {...other}>
            <Typography variant="h6">Connect a Data Source</Typography>
            <Typography color="textSecondary" sx={{mt: 2}} variant="body2">
                You can think a workspace as a group. For instance, if you work for
                multiples companies you will have a different workspace for each of
                them.
            </Typography>
            <ConnectorsCatalogList/>
            <Box sx={{pt: 2}}>
                <Button variant="contained" onClick={handleComplete}>
                    {Object.keys(dataSources).length > 0 ? "Complete" : "Skip for now"}
                </Button>
            </Box>
        </div>
    );
};
