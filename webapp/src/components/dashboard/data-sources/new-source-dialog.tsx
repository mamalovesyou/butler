import Box from '@mui/material/Box';
import Stepper from '@mui/material/Stepper';
import Step from '@mui/material/Step';
import StepLabel from '@mui/material/StepLabel';
import Dialog from '@mui/material/Dialog';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import * as React from "react";
import {useEffect, useState} from "react";
import CloseIcon from '@mui/icons-material/Close';
import {
    Alert,
    AlertTitle,
    Button,
    Divider,
    FormControl,
    IconButton,
    InputLabel,
    MenuItem,
    Select,
    Typography,
} from "@mui/material";
import {useCurrentWorkspace} from "../../../hooks/use-workspace";
import {useDataSources} from "../../../hooks/use-sources";
import {Api, V1AuthType, V1Connector, V1DataSource} from "../../../api";
import AddIcon from "@mui/icons-material/Add";
import {SelectChangeEvent} from '@mui/material/Select';
import {ArrayToObject} from "../../../utils/array";
import {ConfigInputForm} from "./forms/config-input-form";
import OAuthPopup from "../../oauth-popup";
import {listAvailableSourcesRequest} from "../../../features/data-sources";
import {useDispatch} from "react-redux";

interface NewSourceStepperProps {
    sources: Record<string, V1DataSource>;
}

const NewSourceStepper = (props: NewSourceStepperProps) => {
    const steps = ["Choose a data source", "Authentication", "Configuration"];
    const {sources} = props;
    const {workspace} = useCurrentWorkspace();
    const [error, setError] = useState("");
    const [source, setSource] = useState<V1DataSource>(null);
    const [connector, setConnector] = useState<V1Connector>(null);
    const [activeStep, setActiveStep] = useState(0);

    const createConnector = () => {
        Api.v1.connectorsServiceCreateConnector({
            workspaceId: workspace.id,
            airbyteWorkspaceId: workspace.airbyteWorkspaceId,
            airbyteSourceDefinitionId: source.airbyteSourceDefinitionId,
            airbyteDestinationId: workspace.airbyteDestinationId,
        }).then((response) => {
            setConnector(response.data);
            setActiveStep(1);
        }).catch((error) => setError(String(error)));
    }

    const handleSourceChange = (event: SelectChangeEvent) => {
        const dataSource = sources[event.target.value];
        setSource(dataSource);
    };

    const onOAuthConnect = (code: string) => {
        Api.v1.connectorsServiceAuthenticateOAuthConnector({
            connectorId: connector.id,
            code
        }).then((response) => setActiveStep(activeStep + 1)).catch((error) => {
            setError(String(error))
        })
    };

    return <>
        <Stepper activeStep={activeStep} alternativeLabel>
            {steps.map((step) => (
                <Step key={step}>
                    <StepLabel> {step} </StepLabel>
                </Step>
            ))}
        </Stepper>
        {activeStep == 0 ?
            <Box sx={{display: 'flex', flexDirection: 'column', py: 5, px: 10, alignItems: 'center', width: '100%'}}>
                <FormControl fullWidth>
                    <InputLabel>Data Source</InputLabel>
                    <Select
                        label="Data Source"
                        value={source ? source.airbyteSourceDefinitionId : ''}
                        onChange={handleSourceChange}
                    >
                        {Object.values(sources).map((source: V1DataSource, index: number) =>
                            <MenuItem key={index} value={source.airbyteSourceDefinitionId}>
                                {/*<ListItemIcon>*/}
                                {/*    /!*<SourceIcon xml={source.iconSvg} name={source.name} />*!/*/}
                                {/*</ListItemIcon>*/}
                                {/*<ListItemText>{source.name}</ListItemText>*/}
                                {source.name}
                            </MenuItem>)}
                    </Select>
                </FormControl>
                <Box sx={{display: 'flex', py: 4}}>
                    <Button
                        type="submit"
                        variant="contained"
                        onClick={createConnector}
                    >
                        Next
                    </Button>
                </Box>
            </Box> : null}
        {activeStep === 1 ?
            <Box sx={{display: 'flex', flexDirection: 'column', py: 5, px: 10, alignItems: 'center', width: '100%'}}>
                {source?.authType === V1AuthType.OAUTH2 ?
                    <OAuthPopup onCode={onOAuthConnect}
                                title={"OAuth 2"}
                                url={source?.authUrl}>
                        <Button variant="contained" size="large">
                            Authenticate
                        </Button>
                    </OAuthPopup> : <Typography>Not supported</Typography>}
            </Box> : null}
        {(activeStep === 2 && connector) ? <Box>
            <ConfigInputForm connectorId={connector.id}
                             initialValues={connector.config}
                             configJSONSchema={source.configInputJSONSchema}
                             secretsJSONSchema={source.secretsInputJSONSchema}
            />
        </Box> : null}
    </>
}

interface NewSourceDialogProps {
    onClose?: () => void;
}

export const NewSourceDialog = (props: NewSourceDialogProps) => {
    const dispatch = useDispatch();
    const {sources, loading, error} = useDataSources();
    const [sourcesById, setSourcesById] = useState<Record<string, V1DataSource>>({});
    const [open, setOpen] = useState(false);

    useEffect(() => {
        dispatch(listAvailableSourcesRequest())
    }, [])

    useEffect(() => {
        setSourcesById(ArrayToObject(sources, 'airbyteSourceDefinitionId'))
        console.log("sources by id:", sourcesById);
    }, [sources])

    const onClose = (event, reason = '') => {
        if (reason !== 'backdropClick') {
            setOpen(false);
        }
        props.onClose();
    }

    const onComplete = () => {
        setOpen(false);
    };

    return (
        <>
            <Button variant="contained" startIcon={<AddIcon/>} onClick={() => setOpen(true)}>
                Add Source
            </Button>
            {open ? <Dialog
                fullWidth
                maxWidth="md"
                open={open}
                onClose={onClose}
            >
                <DialogTitle>
                    Add a new data source
                    <IconButton
                        aria-label="close"
                        onClick={onClose}
                        sx={{
                            position: 'absolute',
                            right: 8,
                            top: 8,
                            color: (theme) => theme.palette.grey[500],
                        }}
                    >
                        <CloseIcon/>
                    </IconButton>
                </DialogTitle>
                <Divider/>
                <DialogContent>
                    {(!loading && error !== "") ? <Box sx={{width: '100%', p: 4}}>
                            <Alert severity="error">
                                <AlertTitle>Error</AlertTitle>
                                {error}
                            </Alert>
                        </Box> :
                        <Box sx={{width: '100%'}}>
                            <NewSourceStepper sources={sourcesById}/>
                        </Box>}
                </DialogContent>
            </Dialog> : null}
        </>
    );
}

NewSourceDialog.defaultProps = {
    open: false,
    onClose: () => {
    },
}

export default NewSourceDialog;