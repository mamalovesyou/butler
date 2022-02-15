import * as React from 'react';
import isEqual from 'lodash/isEqual'
import Typography from '@mui/material/Typography';
import {Alert, Box, Button, AlertTitle, CircularProgress, Paper} from "@mui/material";
import {useMemo, useState} from "react";
import {styled} from "@mui/material/styles";
import {Api, V1AuthType, V1Connector} from "../../../api";
import {ConnectorWithSource} from "./connectors-table-list";
import {ConfigInputForm} from "./forms/config-input-form";
import OAuthPopup from "../../oauth-popup";
import AlertSlice, {notificationsActions} from "../../../features/notifications/AlertSlice";

const Img = styled('img')({
    margin: 'auto',
    display: 'block',
    maxWidth: '100%',
    maxHeight: '100%',
});

type ConnectorDetailsEditProps = {
    connector: ConnectorWithSource
};

export const ConnectorDetailsEdit = ({connector}: ConnectorDetailsEditProps) => {

    const [error, setError] = useState('');


    const onOAuthConnect = (code: string) => {
        Api.v1.connectorsServiceAuthenticateOAuthConnector({
            connectorId: connector.id,
            code
        }).then(() => {
            console.log()
            notificationsActions.createAlert({ message: "Successfully connected", type: "success"});
        }).catch((error) => {
            setError(String(error))
        })
    };

    return (
        <>
            <Paper sx={{p: 1, my: 1}}>
                <Box sx={{width: '100%', display: 'flex', flexDirection: 'row', alignItems: 'center'}}>
                    <Typography variant="h6" sx={{paddingRight: 2}}>Authentication: </Typography>
                    {connector.source?.authType === V1AuthType.OAUTH2 ?
                        <OAuthPopup onCode={onOAuthConnect}
                                    title={"OAuth 2"}
                                    url={connector.source?.authUrl}>
                            <Button variant="contained" size="large">
                                Authenticate
                            </Button>
                        </OAuthPopup> : <Typography>Not supported</Typography>}
                </Box>
                {error !== "" && (
                    <Box sx={{ py: 1 }}>
                        <Alert severity="error">
                            <AlertTitle>Error</AlertTitle>{error}
                        </Alert>
                    </Box>
                )}
            </Paper>
            <Paper sx={{p: 2, my: 1}}>
                <Typography variant="h6">Configuration</Typography>
                <ConfigInputForm connectorId={connector.id}
                                 initialValues={connector.config}
                                 inputJSONSchema={connector.source.configurationInputJSONSchema} />
            </Paper>
        </>

    );
}


export default ConnectorDetailsEdit;