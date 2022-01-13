import {FC, useEffect} from 'react';
import {Button, Card, CardContent, CardHeader, Chip, Grid} from '@mui/material';
import {useCatalog, useDataSources} from '../../../hooks/use-connectors';
import {connectOAuthConnectorRequest} from '../../../features/connectors';
import {useDispatch} from 'react-redux';
import {useCurrentWorkspace} from "../../../hooks/use-workspace";
import ConnectorItem from "./connector-item";

export const ConnectorsCatalogList: FC = () => {

    const dispatch = useDispatch();
    const {workspace} = useCurrentWorkspace();
    const catalog = useCatalog();
    const dataSources = useDataSources();

    console.log("catalog", catalog)

    const handleConnectOAuth = (params: { code: string, provider: string }) => {
        dispatch(connectOAuthConnectorRequest({workspaceId: workspace.id, ...params}))
    }

    const connectedProviders = Object.values(dataSources).reduce((acc, source) => {
        let {name} = source;
        return {...acc, [name]: true};
    }, {});

    return (
        <Grid container spacing={3}>
            {catalog.map((item, index) =>
                <Grid item key={`${item.name}-${index}`}>
                    <ConnectorItem
                        name={item.name}
                        authUrl={item.authUrl}
                        svgIcon={item.iconSvg}
                        onOAuthConnect={handleConnectOAuth}
                        isConnected={(item.name in connectedProviders)}
                    />
                </Grid>
            )}
        </Grid>
    );
};
