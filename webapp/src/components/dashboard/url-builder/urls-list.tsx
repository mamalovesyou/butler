import { List } from '@mui/material';
import * as React from "react";
import {useDispatch} from "react-redux";


export const URLsList = () => {

    const dispatch = useDispatch();

    return (
        <List dense sx={{width: '100%', p: 0, bgcolor: 'background.paper'}}>
            {/*{enhancedConnectors.map((connector: ConnectorWithSource, index) => {*/}
            {/*    return (*/}
            {/*        <Box key={connector.id}*/}
            {/*             onClick={() => goToConnectorDetails(connector.id)}>*/}
            {/*            <ListItem*/}
            {/*                secondaryAction={!connector.isActive ?*/}
            {/*                    <Chip color="error" label="INACTIVE"/>*/}
            {/*                    : <Chip color="primary" label="ACTIVE"/>}*/}
            {/*                disablePadding*/}
            {/*            >*/}
            {/*                <ListItemButton sx={{py: 2}}>*/}
            {/*                    <ListItemAvatar>*/}
            {/*                        <SourceIcon*/}
            {/*                            xml={connector.source.iconSvg}*/}
            {/*                            name={connector.source.name}/>*/}
            {/*                    </ListItemAvatar>*/}
            {/*                    <ListItemText id={connector.id} sx={{paddingLeft: 2}} primary={connector.source.name} primaryTypographyProps={{variant: "h6"}}/>*/}
            {/*                </ListItemButton>*/}
            {/*            </ListItem>*/}
            {/*            { (index === enhancedConnectors.length - 1) ? null : <Divider component="li"/> }*/}
            {/*        </Box>*/}
            {/*    );*/}
            {/*})}*/}
        </List>
    );
};
