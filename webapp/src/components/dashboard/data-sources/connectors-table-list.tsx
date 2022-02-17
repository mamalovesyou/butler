import {Fragment, useEffect, useState} from 'react';
import {
    Box, Chip, Divider, FormControlLabel, List, ListItem, ListItemAvatar, ListItemButton, ListItemText, Switch,
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
        console.log("click: ", connectorId)
        dispatch(push(DATA_SOURCES_ROOT_PATH + '/' + connectorId))
    }

    return (
        <List dense sx={{width: '100%', p: 0, bgcolor: 'background.paper'}}>
            {enhancedConnectors.map((connector: ConnectorWithSource, index) => {
                return (
                    <Box key={connector.id}
                         onClick={() => goToConnectorDetails(connector.id)}>
                        <ListItem
                            secondaryAction={!connector.isActive ?
                                <Chip color="error" label="INACTIVE"/>
                                : <Chip color="primary" label="ACTIVE"/>}
                            disablePadding
                        >
                            <ListItemButton sx={{py: 2}}>
                                <ListItemAvatar>
                                    <SourceIcon
                                        xml={connector.source.iconSvg}
                                        name={connector.source.name}/>
                                </ListItemAvatar>
                                <ListItemText id={connector.id} sx={{paddingLeft: 2}} primary={connector.source.name} primaryTypographyProps={{variant: "h6"}}/>
                            </ListItemButton>
                        </ListItem>
                        { (index === enhancedConnectors.length - 1) ? null : <Divider component="li"/> }
                    </Box>
                );
            })}
        </List>
    );
};
