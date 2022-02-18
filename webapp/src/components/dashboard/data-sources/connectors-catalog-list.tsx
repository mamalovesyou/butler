import {FC, Fragment, useEffect, useState} from 'react';
import {
    Table, TableBody,
    TableCell,
    TableHead,
    TableRow
} from '@mui/material';
import {useConnectors} from '../../../hooks/use-connectors';
import {
    connectOAuthConnectorRequest,
    listWorkspaceConnectorsRequest
} from '../../../features/connectors';
import {useDispatch} from 'react-redux';
import {useWorkspace} from "../../../hooks/use-workspace";
import { V1WorkspaceConnector } from "../../../api";
import CatalogTableRow from "./catalog-table-row";
import ConfigureAccountDialog from "./configure-account-dialog";

export const ConnectorsCatalogList: FC = () => {

    const dispatch = useDispatch();
    const {workspaceId} = useWorkspace();
    const {connectors, configure} = useConnectors();
    const [workspaceConnectorByName, setWorkspaceConnectorByName] = useState({});


    useEffect(() => {
        if (workspaceId) {
            dispatch(listWorkspaceConnectorsRequest({workspaceId}));
        }
    }, [workspaceId]);

    // useEffect(() => {
    //     const wsConnectorByName = {};
    //     Object.values(connectors).forEach((connector: V1WorkspaceConnector) => wsConnectorByName[connector.name] = connector);
    //     setWorkspaceConnectorByName(wsConnectorByName);
    // }, [, connectors]);

    const handleConnectOAuth = async (params: { code: string, provider: string }) => {
        dispatch(connectOAuthConnectorRequest({workspaceId, ...params}));
    }


    return (
        <Fragment>
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
                    {/*{catalog.map((connector: V1CatalogConnector, index) =>*/}
                    {/*    <CatalogTableRow key={`${connector.name}-${index}`}*/}
                    {/*                     workspaceConnector={workspaceConnectorByName[connector.name]}*/}
                    {/*                     isConnected={(connector.name in workspaceConnectorByName)}*/}
                    {/*                     connector={connector} onOAuthConnect={handleConnectOAuth}/>*/}
                    {/*)}*/}
                </TableBody>
            </Table>
        <ConfigureAccountDialog />
        </Fragment>
    );
};
