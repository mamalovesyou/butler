import {FC, useEffect, useState} from 'react';
import {Box, Button, CircularProgress, Container, Divider, Typography} from '@mui/material';
import {ConnectorsCatalogList} from '../../../components/dashboard/data-sources/connectors-catalog-list';
import {useDispatch} from 'react-redux';
import {listAvailableSourcesRequest} from "../../../features/data-sources";
import {listWorkspaceConnectorsRequest} from "../../../features/connectors";
import {useCurrentWorkspace} from "../../../hooks/use-workspace";
import {useConnectors} from "../../../hooks/use-connectors";
import {useDataSources} from "../../../hooks/use-sources";
import {useParams} from "react-router-dom";
import ConnectorDetailsEdit from "../../../components/dashboard/data-sources/connector-details-edit";
import {ArrayToObject} from "../../../utils/array";
import TestConnection from "../../../components/dashboard/data-sources/test-connection";


const DataSourceDetail: FC = () => {

    const dispatch = useDispatch();
    const {workspace} = useCurrentWorkspace();
    const { connectorId } = useParams();
    const { connectors } = useConnectors();
    const { sources } = useDataSources();
    const [connectorWithSource, setConnectorWithSource] = useState(null);

    useEffect(() => {
        if (workspace?.id) {
            dispatch(listAvailableSourcesRequest());
            dispatch(listWorkspaceConnectorsRequest({workspaceId: workspace.id}));
        }
    }, [workspace]);

    useEffect(() => {
        if (sources.length && connectors.length && connectorId) {
            const connectorsById = ArrayToObject(connectors, 'id');
            const sourcesByDefinitionId = ArrayToObject(sources, 'airbyteSourceDefinitionId');
            const connector = connectorsById[connectorId];
            setConnectorWithSource({ ...connector, source: sourcesByDefinitionId[connector.airbyteSourceDefinitionId]})
        }
    }, [connectorId, sources, connectors]);

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
                    { connectorWithSource ?  <Typography variant="h4"> { connectorWithSource.source.name }</Typography> : null }
                </Box>
                <Divider sx={{mb: 3}}/>
                { connectorWithSource ?
                    <Box>
                        <ConnectorDetailsEdit connector={connectorWithSource} />
                    </Box>
                    : null }
            </Container>
        </Box>
    );
};

export default DataSourceDetail;
