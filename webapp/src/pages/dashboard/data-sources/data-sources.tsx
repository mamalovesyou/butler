import {FC, useEffect, useState} from 'react';
import {Box, Button, CircularProgress, Container, Divider, Typography} from '@mui/material';
import {ConnectorsCatalogList} from '../../../components/dashboard/data-sources/connectors-catalog-list';
import {useDispatch} from 'react-redux';
import {listAvailableSourcesRequest} from "../../../features/data-sources";
import {listWorkspaceConnectorsRequest} from "../../../features/connectors";
import {useCurrentWorkspace} from "../../../hooks/use-workspace";
import AddIcon from '@mui/icons-material/Add';
import {NewSourceDialog} from "../../../components/dashboard/data-sources/new-source-dialog";
import {useConnectors} from "../../../hooks/use-connectors";
import {useDataSources} from "../../../hooks/use-sources";
import {ConnectorsTableList} from "../../../components/dashboard/data-sources/connectors-table-list";

const DataSources: FC = () => {

    const dispatch = useDispatch();
    const {workspace} = useCurrentWorkspace();
    const {connectors, loading: connectorsLoading} = useConnectors();
    const {sources, loading: sourcesLoading} = useDataSources();
    const [sourceDialogOpen, setSourceDialogOpen] = useState(false);

    useEffect(() => {
        if (workspace?.id) {
            dispatch(listAvailableSourcesRequest());
            dispatch(listWorkspaceConnectorsRequest({workspaceId: workspace.id}));
        }
    }, [workspace]);

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
                    <Typography variant="h4">Data Sources</Typography>
                    <NewSourceDialog/>
                </Box>
                <Divider sx={{mb: 3}}/>
                {connectorsLoading || sourcesLoading ? <CircularProgress/> :
                    <ConnectorsTableList sources={sources} connectors={connectors}/>}
            </Container>
        </Box>
    );
};

export default DataSources;
