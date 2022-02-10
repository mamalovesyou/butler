import * as React from 'react';
import Typography from '@mui/material/Typography';
import {Button, Chip, Grid, Paper} from "@mui/material";
import CheckIcon from '@mui/icons-material/Check';
import OAuthPopup from "../../oauth-popup";
import {useMemo} from "react";
import {styled} from "@mui/material/styles";
import {V1Connector} from "../../../api";
import {ConnectorWithSource} from "./connectors-table-list";
import {ConfigInputForm} from "./forms/config-input-form";

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

    return (
        <Paper sx={{p: 1}}>
            <Typography>Configuration</Typography>
            <ConfigInputForm connectorId={connector.id} source={connector.source} onComplete={() => {}} />
        </Paper>
    );
}


export default ConnectorDetailsEdit;