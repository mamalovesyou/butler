import {Fragment, useEffect, useState} from 'react';
import {
    Box, Chip, FormControlLabel, Switch,
    Table, TableBody,
    TableCell,
    TableHead,
    TableRow
} from '@mui/material';
import {V1Connector, V1DataSource} from "../../../api";
import {ArrayToObject} from "../../../utils/array";
import * as React from "react";
import Typography from "@mui/material/Typography";
import SourceIcon from "./source-icon";
import {TableRowProps} from "@mui/material/TableRow/TableRow";
import {useDispatch} from "react-redux";
import {push} from "redux-first-history";
import {DATA_SOURCES_ROOT_PATH} from "../../../routes";

export type ConnectorWithSource = V1Connector & {
    source: V1DataSource
}

type ConnectorTableRowProps = TableRowProps & {
    connector: ConnectorWithSource;
};

export const ConnectorTableRow = (props: ConnectorTableRowProps) => {
    const {connector, ...others} = props;
    return (
        <TableRow {...others} hover>
            <TableCell>
                <Box
                    sx={{
                        alignItems: 'center',
                        display: 'flex'
                    }}
                >
                    <SourceIcon
                        xml={connector.source.iconSvg}
                        name={connector.source.name}/>
                    <Typography
                        variant="subtitle2">{connector.source.name}</Typography>
                </Box>
            </TableCell>
            <TableCell sx={{display: 'flex', justifyContent: 'flex-end'}}>
                <Chip label="Uncomplete"/>
            </TableCell>
        </TableRow>
    );
}

interface ConnectorsTableListProps {
    sources: V1DataSource[]
    connectors: V1Connector[]
}

export const ConnectorsTableList = (props: ConnectorsTableListProps) => {

    const dispatch = useDispatch();
    const {sources, connectors} = props;
    const [enhancedConnectors, setEnhancedConnectors] = useState([]);

    useEffect(() => {
        console.log("sources or connectors changed", sources, connectors)
        const sourcesByDefinitionId = ArrayToObject(sources, 'airbyteSourceDefinitionId');
        const result: ConnectorWithSource[] = connectors.map(
            (c: V1Connector) => ({...c, source: sourcesByDefinitionId[c.airbyteSourceDefinitionId]}))
        setEnhancedConnectors(result);
    }, [sources, connectors]);

    const goToConnectorDetails = (connectorId: string) => {
        dispatch(push(DATA_SOURCES_ROOT_PATH + '/' + connectorId))
    }

    return (
        <Table sx={{minWidth: 400}}>
            <TableHead>
                <TableRow>
                    <TableCell>Connect a new data sources</TableCell>
                    <TableCell/>
                </TableRow>
            </TableHead>
            <TableBody>
                {enhancedConnectors.map((connector: ConnectorWithSource, index) =>
                    <ConnectorTableRow
                        onClick={() => goToConnectorDetails(connector.id)}
                        sx={{
                            cursor: 'pointer'
                        }}
                        key={`${connector.source.name}-${index}`}
                        connector={connector}/>
                )}
            </TableBody>
        </Table>
    );
};
