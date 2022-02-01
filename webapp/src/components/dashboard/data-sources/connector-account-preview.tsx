import {Box, Chip, List, ListItem, ListItemText, Typography} from "@mui/material";
import {V1ProviderAccount} from "../../../api";

interface ConnectorAccountPreviewProps {
    provider: string;
    account: V1ProviderAccount;
}

export const ConnectorAccountPreview = (props: ConnectorAccountPreviewProps) => {

    return (
        <>
            {props.account ? <Box sx={{display: 'flex', flexDirection: 'column'}}>
                <Box sx={{display: 'flex', flexDirection: 'row', my: 2 }}>
                    <Typography sx={{px: 2}} variant="h6">Provider:</Typography>
                    <Typography variant="h6">{props.provider}</Typography>
                </Box>
                <Box sx={{display: 'flex', flexDirection: 'row', my: 2}}>
                    <Typography sx={{px: 2}} variant="h6">Account ID:</Typography>
                    <Typography variant="h6">{props.account.id}</Typography>
                </Box>
                <Box sx={{display: 'flex', flexDirection: 'row', my: 2}}>
                    <Typography sx={{px: 2}} variant="h6">Account Name:</Typography>
                    <Typography variant="h6">{props.account.name}</Typography>
                </Box>
                <Box sx={{display: 'flex', flexDirection: 'row', my: 2}}>
                    <Typography sx={{px: 2}} variant="h6">Account Type:</Typography>
                    { props.account.test ? <Chip label="TEST" color="primary"/> : <Chip label="PROD" color="primary"/> }
                </Box>
            </Box> : null}
        </>
    );
};

export default ConnectorAccountPreview;