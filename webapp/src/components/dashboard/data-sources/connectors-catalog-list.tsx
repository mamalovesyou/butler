import {FC, Fragment, useEffect, useState} from 'react';
import {
    Avatar,
    Box,
    Button,
    Card,
    CardContent,
    CardHeader,
    Chip,
    Grid, IconButton,
    Tab,
    Table, TableBody,
    TableCell,
    TableHead,
    TableRow, Typography
} from '@mui/material';
import {useCatalog, useConnectors, useDataSources} from '../../../hooks/use-connectors';
import {
    connectOAuthConnectorRequest,
    listCatalogConnectorsRequest,
    listWorkspaceConnectorsRequest
} from '../../../features/connectors';
import {useDispatch} from 'react-redux';
import {useWorkspace} from "../../../hooks/use-workspace";
import {V1CatalogConnector, V1WorkspaceConnector} from "../../../api";
import CatalogTableRow from "./catalog-table-row";
import ConfigureAccountDialog from "./configure-account-dialog";

export const ConnectorsCatalogList: FC = () => {

    const dispatch = useDispatch();
    const {workspaceId} = useWorkspace();
    const {catalog, connectors} = useConnectors();
    const [workspaceConnectorByName, setWorkspaceConnectorByName] = useState({});
    const [openConfigureAccount, setOpenConfigureAccount] = useState(false);
    const [configureAccountProvider, setConfigureAccountProvider] = useState("");

    useEffect(() => {
        dispatch(listCatalogConnectorsRequest());
    }, []);

    useEffect(() => {
        console.log("worksapce id")
        if (workspaceId) {
            dispatch(listWorkspaceConnectorsRequest({workspaceId}));
        }
    }, [workspaceId]);

    useEffect(() => {
        const wsConnectorByName = {};
        Object.values(connectors).forEach((connector: V1WorkspaceConnector) => wsConnectorByName[connector.name] = connector);
        console.log("connector by bame", wsConnectorByName);
        setWorkspaceConnectorByName(wsConnectorByName);
    }, [catalog, connectors]);

    const handleConnectOAuth = (params: { code: string, provider: string }) => {
        dispatch(connectOAuthConnectorRequest({workspaceId, ...params}));
        setOpenConfigureAccount(true);
        setConfigureAccountProvider(params.provider);
    }


    return (
        <Fragment>
        <Box sx={{width: '100%', p: 1}}>
            <Table sx={{minWidth: 400}}>
                <TableHead>
                    <TableRow>
                        <TableCell>Connect a new data sources</TableCell>
                        <TableCell/>
                        <TableCell/>
                        <TableCell/>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {catalog.map((connector: V1CatalogConnector, index) =>
                        <CatalogTableRow key={`${connector.name}-${index}`}
                                         workspaceConnector={workspaceConnectorByName[connector.name]}
                                         isConnected={(connector.name in workspaceConnectorByName)}
                                         connector={connector} onOAuthConnect={handleConnectOAuth}/>
                    )}
                </TableBody>
            </Table>
        </Box>
        <ConfigureAccountDialog provider={configureAccountProvider} show={openConfigureAccount} />
        </Fragment>
    );
};
