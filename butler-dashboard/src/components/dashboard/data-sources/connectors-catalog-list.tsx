import { FC, useEffect } from 'react';
import { Button, Card, CardContent, CardHeader, Chip, Grid } from '@mui/material';
import { useCatalog, useDataSources } from '../../../hooks/use-connectors';
import { connectOAuthConnectorRequest } from '../../../features/connectors';
import OAuthPopup from '../../oauth-popup';
import { useDispatch } from 'react-redux';
import DoneIcon from '@mui/icons-material/Done';

export const ConnectorsCatalogList: FC = () => {

    const dispatch = useDispatch();
    const catalog = useCatalog();
    const dataSources = useDataSources();

    console.log(catalog)

    const handleConnectOAuth = (params: { code: string, name: string }) => {
        dispatch(connectOAuthConnectorRequest(params))
    }

    const connectedProviders = Object.values(dataSources).reduce((acc, source) => {
        let { name } = source;
        return { ...acc, [name]: true };
    }, {});

    return (
        <Grid container spacing={2}>
            {Object.entries(catalog).map(([key, item]) => <Grid key={key} item xs={12} md={6} lg={4}>
                <Card>
                    <CardHeader
                        title={item.name}
                    />
                    <CardContent>
                        {connectedProviders[item.name]
                            ? <Chip label="Connected" color="success" deleteIcon={<DoneIcon />} />
                            : <OAuthPopup
                                url={item.authUrl}
                                title={`Connect ${item.name}`}
                                onCode={(code, params) => {
                                    handleConnectOAuth({
                                        code,
                                        name: item.name
                                    })
                                }}
                                onClose={() => {
                                    console.log("window closed")
                                }}
                            >
                                <Button variant="contained" fullWidth size="large">
                                    Try OAuth
                                </Button>
                            </OAuthPopup>

                        }
                    </CardContent>
                </Card>

            </Grid>
            )}
        </Grid>
    );
};
