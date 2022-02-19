import { ChangeEvent, useState } from 'react';
import type { FC, MutableRefObject } from 'react';
import PropTypes from 'prop-types';
import {
    Box,
    Button,
    Drawer,
    IconButton, Switch,
    TextField, Theme,
    Typography, useMediaQuery
} from '@mui/material';
import { styled } from '@mui/material/styles';
import { X as XIcon } from '../../../icons/x';


interface IUTMConfigurationFormProps {
    autopilot: boolean;
    onSave: () => void;
}

const UTMConfigurationForm: FC<IUTMConfigurationFormProps> = (props) => {
    const { onSave, autopilot } = props;

    return (
        <>
            <Typography
                sx={{ my: 3 }}
                variant="h6"
            >
                Details
            </Typography>
            <TextField
                fullWidth
                label="URL"
                margin="normal"
                name="url"
            />
            <TextField
                disabled={autopilot}
                fullWidth
                label="UTM Source"
                margin="normal"
                name="utm_source"
            />
            <TextField
                disabled={autopilot}
                fullWidth
                label="UTM Medium"
                margin="normal"
                name="utm_medium"
            />
            <TextField
                disabled={autopilot}
                fullWidth
                label="UTM Campaign"
                margin="normal"
                name="utm_campaign"
            />
            <TextField
                disabled={autopilot}
                fullWidth
                label="UTM Content"
                margin="normal"
                name="utm_content"
            />
            <TextField
                disabled={autopilot}
                fullWidth
                label="UTM Source"
                margin="normal"
                name="utm_source"
            />
            <Button
                color="primary"
                sx={{ mt: 3 }}
            >
                Save
            </Button>
        </>
    );
};


export const UTMDetails = () => {
    const [checked, setChecked] = useState(true);

    const handleAutoPilotChange = (event: ChangeEvent<HTMLInputElement>) => {
        setChecked(event.target.checked);
    };

    return <Box
        sx={{
            px: 3,
            py: 4
        }}
    >
        <Box
            sx={{
                alignItems: 'center',
                backgroundColor: (theme) => theme.palette.mode === 'dark'
                    ? 'neutral.800'
                    : 'neutral.100',
                borderRadius: 1,
                display: 'flex',
                flexWrap: 'wrap',
                justifyContent: 'space-between',
                px: 3,
                py: 2.5
            }}
        >
            <Typography
                color="textSecondary"
                sx={{ mr: 2 }}
                variant="overline"
            >
                Autopilot Mode
            </Typography>
            <Box
                sx={{
                    alignItems: 'right',
                    display: 'flex',
                    flexWrap: 'wrap',
                    m: -1,
                    '& > button': {
                        m: 1
                    }
                }}
            >
                <Switch checked={checked} onChange={handleAutoPilotChange} />
            </Box>
        </Box>
        <UTMConfigurationForm autopilot={checked} onSave={() => { }} />
    </Box>
};

export default UTMDetails;
