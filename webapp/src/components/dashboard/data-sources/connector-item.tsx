import * as React from 'react';
import Typography from '@mui/material/Typography';
import {Button, Chip, Grid, Paper} from "@mui/material";
import CheckIcon from '@mui/icons-material/Check';
import OAuthPopup from "../../oauth-popup";
import {useMemo} from "react";
import {styled} from "@mui/material/styles";

const Img = styled('img')({
    margin: 'auto',
    display: 'block',
    maxWidth: '100%',
    maxHeight: '100%',
});

type IConnectorItemProps = {
    svgIcon: string;
    name: string;
    authUrl: string;
    isConnected: boolean;
    onOAuthConnect: (params: { code: string, provider: string }) => void;
};

export const ConnectorItem: React.FC<IConnectorItemProps> = (props) => {

    const imageURL = useMemo(() => {
        const blob = new Blob([props.svgIcon], {type: 'image/svg+xml'});
        return URL.createObjectURL(blob);
    }, [props.svgIcon]);

    return (
        <Paper sx={{ p: 1, margin: 'auto', maxWidth: 260, minHeight: 106, maxHeight: 106, flexGrow: 1 }}>
            <Grid container>
                <Grid item>
                    <Grid sx={{ width: 90, height: 90 }}>
                        <Img alt={`${props.name} logo`} src={imageURL} />
                    </Grid>
                </Grid>
                <Grid item xs={12} sm container>
                    <Grid item xs container direction="column" spacing={2}>
                        <Grid item xs>
                            <Typography gutterBottom variant="h5" component="div">
                                {props.name}
                            </Typography>
                            {props.isConnected ? <Chip label="Connected" color="success" deleteIcon={<CheckIcon/>}/>
                                : <OAuthPopup
                                    url={props.authUrl}
                                    title={`Connect ${props.name}`}
                                    onCode={(code, params) => props.onOAuthConnect({
                                        code,
                                        provider: props.name
                                    })
                                    }
                                    onClose={() => {
                                        console.log("window closed")
                                    }}
                                >
                                    <Button variant="contained" fullWidth size="large">
                                        Connect
                                    </Button>
                                </OAuthPopup>}
                        </Grid>
                    </Grid>
                </Grid>
            </Grid>
        </Paper>
    );
}

export default ConnectorItem;