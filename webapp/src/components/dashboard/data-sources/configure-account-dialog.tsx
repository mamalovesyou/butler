import * as React from 'react';
import Box from '@mui/material/Box';
import Stepper from '@mui/material/Stepper';
import Step from '@mui/material/Step';
import StepLabel from '@mui/material/StepLabel';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import {useEffect} from "react";
import {Table, TableBody, TableCell, TableHead, TableRow} from "@mui/material";
import {SeverityPill} from "../../severity-pill";

const steps = [
    'Select account',
    'Verify configuration',
];

interface IConfigureAccountDialogProps {
    provider: string;
    show: boolean;
}

export const ConfigureAccountDialog = (props: IConfigureAccountDialogProps) => {
    const [open, setOpen] = React.useState(false);
    const handleClose = () => setOpen(false);
    const [activeStep, setActiveStep] = React.useState(0);

    useEffect(() => {
        setOpen(props.show);
    }, [props.show])

    const handleBack = () => {
        setActiveStep((prevActiveStep) => prevActiveStep - 1);
    };

    const handleFinish = () => {
        setOpen(false);
    }

    return (
        <Dialog
            fullWidth
            maxWidth="sm"
            open={open}
            onClose={handleClose}
        >
            <DialogTitle>Configure {props.provider}</DialogTitle>
            <DialogContent>
                <Box sx={{width: '100%', p: 2}}>
                    <Stepper activeStep={activeStep} alternativeLabel>
                        {steps.map((label) => (
                            <Step key={label}>
                                <StepLabel>{label}</StepLabel>
                            </Step>
                        ))}
                    </Stepper>
                </Box>
                <Box sx={{width: '100%'}}>
                    <Table sx={{minWidth: 400}}>
                        <TableHead>
                            <TableRow>
                                <TableCell>Account name</TableCell>
                                <TableCell>Account ID</TableCell>
                                <TableCell/>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            <TableRow>
                                <TableCell>Heybutler.io</TableCell>
                                <TableCell>00000000</TableCell>
                                <TableCell/>
                            </TableRow>
                            <TableRow>
                                <TableCell>Heybutler.io Test</TableCell>
                                <TableCell>12</TableCell>
                                <TableCell><SeverityPill>Test</SeverityPill></TableCell>
                            </TableRow>
                        </TableBody>
                    </Table>
                </Box>
                {activeStep > 0 ?? <Box sx={{display: 'flex', flexDirection: 'row', pt: 2}}>
                    <Button
                        color="inherit"
                        disabled={activeStep === 0}
                        onClick={handleBack}
                        sx={{mr: 1}}
                    >
                        Back
                    </Button>
                    <Box sx={{flex: '1 1 auto'}}/>
                    <Button onClick={handleFinish}>
                        Finish
                    </Button>
                </Box>
                }
            </DialogContent>
            <DialogActions>
                <Button onClick={handleClose}>Close</Button>
            </DialogActions>
        </Dialog>
    );
}

export default ConfigureAccountDialog;