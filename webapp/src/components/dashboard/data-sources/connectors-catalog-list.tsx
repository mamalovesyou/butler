import {FC, useEffect, useState} from 'react';
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
import {useCurrentWorkspace, useWorkspace} from "../../../hooks/use-workspace";
import ConnectorItem from "./connector-item";
import {put} from "redux-saga/effects";
import * as WorkspaceActions from "../../../features/workspace/WorkspaceActions";
import TabContext from "@mui/lab/TabContext";
import TabList from "@mui/lab/TabList";
import TabPanel from "@mui/lab/TabPanel";
import {V1CatalogConnector, V1Invitation, V1UserMember, V1WorkspaceConnector} from "../../../api";
import {UserCircle as UserCircleIcon} from "../../../icons/user-circle";
import {SeverityPill} from "../../severity-pill";
import {DotsHorizontal as DotsHorizontalIcon} from "../../../icons/dots-horizontal";
import CatalogTableRow from "./catalog-table-row";

export const ConnectorsCatalogList: FC = () => {

    const dispatch = useDispatch();
    const {workspaceId} = useWorkspace();
    const {catalog, connectors} = useConnectors();
    const [workspaceConnectorByName, setWorkspaceConnectorByName] = useState({});

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
        dispatch(connectOAuthConnectorRequest({workspaceId, ...params}))
    }


    return (
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
    );
};
