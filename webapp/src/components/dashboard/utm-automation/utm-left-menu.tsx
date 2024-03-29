import type {FC} from 'react';
import {Box, Divider, List, ListItemButton} from '@mui/material';
import {V1Connector, V1WorkspaceConnector} from "../../../api";
import {ConnectorIcon} from "../../connector-icon";
import {useEffect, useState} from "react";

interface UTMLeftMenuProps {
    onClickConnector?: (connectorId: string) => void;
    onSelectConnector?: (connector: V1Connector) => void;
    connectors: V1Connector[];
}

export const UTMLeftMenu = (props: UTMLeftMenuProps) => {

    const [selectedConnectorId, setSelectedConnectorId] = useState("")
    const {connectors, onClickConnector, onSelectConnector} = props;

    useEffect(() => {
        if (connectors.length) {
            setSelectedConnectorId(connectors[0].id)
        }
    }, [connectors])

    const selectConnector = (connectorId: string) => {
        setSelectedConnectorId(connectorId);
        onSelectConnector ? onSelectConnector(connectors[connectorId]) : null;
        onClickConnector ? onClickConnector(connectorId) : null;
    }

    return (
        <Box
            sx={{
                display: 'flex',
                color: '#FFFFFF',
                width: 75,

                borderRightColor: 'divider',
                borderRightStyle: 'solid',
                borderRightWidth: (theme) =>
                    theme.palette.mode === 'dark' ? 1 : 0,
            }}
        >

            <List sx={{
                maxHeight: '100%',
                overflow: 'auto',
                p: 0
            }}>
                {connectors.map((connector: V1WorkspaceConnector) => (
                    <div key={connector.id}>
                        <ListItemButton sx={{height: 75, width: 75}}
                                        onClick={() => selectConnector(connector.id)}
                                        selected={connector.id === selectedConnectorId}>
                            <ConnectorIcon name={connector.name} height={75} width={75}/>
                        </ListItemButton>
                        <Divider/>
                    </div>
                ))}
            </List>
        </Box>
    );
};

export default UTMLeftMenu;
