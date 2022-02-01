import * as Yup from 'yup';
import { useFormik } from 'formik';
import Box from '@mui/material/Box';
import Stepper from '@mui/material/Stepper';
import Step from '@mui/material/Step';
import StepLabel from '@mui/material/StepLabel';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import { Fragment, useEffect, useState } from "react";
import CloseIcon from '@mui/icons-material/Close';
import {
    Alert,
    AlertTitle,
    CircularProgress, FormControlLabel, FormGroup, FormHelperText, IconButton,
    Switch,
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableRow,
    TextField,
    Typography,
} from "@mui/material";
import { SeverityPill } from "../../severity-pill";
import { useConnectors } from "../../../hooks/use-connectors";
import { useDispatch } from "react-redux";
import {
    listConnectorAccountsRequest,
    setConfigureAccountDialogOpen, updateConnectorConfigRequest
} from "../../../features/connectors";
import { useWorkspace } from "../../../hooks/use-workspace";
import { V1ProviderAccount } from "../../../api";
import ConnectorAccountPreview from "./connector-account-preview";
import { AccountSetupForm } from './account-setup-form';

const steps = [
    'Connect',
    'Select account',
    'Configure',
];


export const ConfigureAccountDialog = () => {
    const dispatch = useDispatch();
    const { workspaceId } = useWorkspace();
    const [activeStep, setActiveStep] = useState(0);
    const { loading, configure, error, accounts } = useConnectors();
    const [account, setAccount] = useState(undefined);


    useEffect(() => {
        if (configure.connected) {
            dispatch(listConnectorAccountsRequest({ provider: configure.provider, workspaceId }))
            setActiveStep(1);
        } else setActiveStep(0);
    }, [configure.connected])

    // If error while loading accounts then force manual setup
    useEffect(() => {
        if (activeStep === 1 && error !== "") {
            setActiveStep(2);
        }
    }, [error])

    const handleBack = () => {
        setActiveStep((prevActiveStep) => prevActiveStep - 1);
    };

    const handleSelectAccount = async (account: V1ProviderAccount) => {
        setAccount(account);
        setActiveStep(2);
    }

    const onClose = (event, reason = '') => {
        if (reason !== 'backdropClick') {
            dispatch(setConfigureAccountDialogOpen(false));
        }
    }

    return (
        <Dialog
            fullWidth
            maxWidth="sm"
            open={configure.open}
            onClose={onClose}
        >
            <DialogTitle>
                Configure {configure.provider}
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
                    <CloseIcon />
                </IconButton>
            </DialogTitle>
            <DialogContent>
                <Box sx={{ width: '100%', p: 2 }}>
                    <Stepper activeStep={activeStep} alternativeLabel>
                        {steps.map((label) => (
                            <Step key={label}>
                                <StepLabel>{label}</StepLabel>
                            </Step>
                        ))}
                    </Stepper>
                </Box>
                {loading ? <Box sx={{ display: 'flex', justifyContent: 'center', width: '100%', p: 4 }}><CircularProgress /></Box> : null}
                {/* {(!loading && error !== "") ? <Box sx={{ width: '100%', p: 4 }}>
                    <Alert severity="error">
                        <AlertTitle>Error</AlertTitle>
                        {error}
                    </Alert>
                </Box> : null
                } */}
                {(!loading && error !== "" && activeStep === 1) ? <Box sx={{ width: '100%' }}>
                    <Table sx={{ minWidth: 400 }}>
                        <TableHead>
                            <TableRow>
                                <TableCell>Account name</TableCell>
                                <TableCell>Account ID</TableCell>
                                <TableCell />
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {accounts.length === 0 ? <TableRow>No accounts found</TableRow> : null}
                            {accounts.map((account: V1ProviderAccount, index: number) =>
                                <TableRow sx={{ cursor: 'pointer' }} key={index} hover onClick={() => handleSelectAccount(account)}>
                                    <TableCell>{account.name}</TableCell>
                                    <TableCell>{account.id}</TableCell>
                                    <TableCell>{account.test ??
                                        <SeverityPill>TEST ACCOUNT</SeverityPill>}</TableCell>
                                </TableRow>
                            )}
                        </TableBody>
                    </Table>
                </Box> : null}
                {(activeStep === 2) ?
                    <AccountSetupForm provider={configure.provider} connectorId={configure.connectorId} account={account} onBack={handleBack} /> : null
                }
            </DialogContent>
        </Dialog>
    );
}

export default ConfigureAccountDialog;