import React, {useEffect, useState} from "react";
import type {StepIconProps} from "@mui/material";
import {Avatar, Box, Button, Grid, Step, StepContent, StepLabel, Stepper, Typography,} from "@mui/material";
import {useDispatch} from "react-redux";
import {useCurrentWorkspace} from "../hooks/use-workspace";
import {Check as CheckIcon} from "../icons/check";
import {useAuth} from "../hooks/use-auth";
import {completeOnboardingRequest} from "../features/onboarding";
import OnboardingLeftPannel from "../components/onboarding/onboarding-left-pannel";
import CreateOrganizationForm from "../components/common/forms/create-organization-form";
import CreateWorkspaceForm from "../components/common/forms/create-workspace-form";
import {V1Organization, V1Workspace} from "../api";
import {setCurrentOrganization, setCurrentWorkspace} from "../features/workspace";
import {ANALYTICS_ROOT_PATH, DATA_SOURCES_ROOT_PATH} from "../routes";
import { Link } from "react-router-dom";
import {push} from "redux-first-history";


export enum OnboardingStep {
    CREATE_ORGANIZATION = 0,
    CREATE_WORKSPACE = 1,
    CONNECT_DATA_SOURCE = 2
}


const StepIcon: React.FC<StepIconProps> = (props) => {
    const {active, completed, icon} = props;

    const highlight = active || completed;

    return (
        <Avatar
            sx={{
                backgroundColor: highlight && "secondary.main",
                color: highlight && "secondary.contrastText",
                height: 40,
                width: 40,
            }}
            variant="rounded"
        >
            {completed ? <CheckIcon fontSize="small"/> : icon}
        </Avatar>
    );
};

const Onboarding: React.FC = () => {
    const {user} = useAuth();
    const dispatch = useDispatch();

    const [activeStep, setActiveStep] = useState(OnboardingStep.CREATE_ORGANIZATION);

    const { organization, workspace } = useCurrentWorkspace();

    useEffect(() => {
        console.log(organization)
        if (organization !== null) {
            if (workspace !== null) dispatch(push(ANALYTICS_ROOT_PATH));
            else { setActiveStep(OnboardingStep.CREATE_WORKSPACE) };
        }
    }, [organization, workspace])

    const onOrganizationCreate = (org: V1Organization) => {
        dispatch(setCurrentOrganization(org));
        setActiveStep(OnboardingStep.CREATE_WORKSPACE);
    }

    const onWorkspaceCreate = (ws: V1Workspace) => {
        dispatch(completeOnboardingRequest(ws.organizationId))
        dispatch(setCurrentWorkspace(ws));
        setActiveStep(OnboardingStep.CONNECT_DATA_SOURCE);
    }

    const steps = {
        [OnboardingStep.CREATE_ORGANIZATION]: {
            label: "Organization",
            content: <div>
                <Typography variant="h6">Create an Organization</Typography>
                <Typography color="textSecondary" sx={{mt: 2}} variant="body2">
                    What's your company name ? And what is your role or job title ?
                </Typography>
                <CreateOrganizationForm onSuccess={onOrganizationCreate} />
            </div>,
        },
        [OnboardingStep.CREATE_WORKSPACE]: {
            label: "Workspace",
            content: <div>
                <Typography variant="h6">Create a workspace</Typography>
                <Typography color="textSecondary" sx={{mt: 2}} variant="body2">
                    You can think a workspace as a group. For instance, if you work for
                    multiples companies you will have a different workspace for each of
                    them.
                </Typography>
                { organization ? <CreateWorkspaceForm organizationId={organization.id} onSuccess={onWorkspaceCreate} /> : null }
            </div>,
        },
        [OnboardingStep.CONNECT_DATA_SOURCE]: {
            label: "Connect a source",
            content: <>
                <Typography variant="h6">Connect a Data Source</Typography>
                <Typography color="textSecondary" sx={{mt: 2}} variant="body2">
                    You can think of data sources as channels
                </Typography>
                <Box sx={{pt: 2, display: 'flex', flexDirection: 'row', justifyContent: 'center'}}>
                    <Button component={Link} variant="contained" color="primary" to={DATA_SOURCES_ROOT_PATH}>
                        Connect source
                    </Button>
                </Box>
            </>,
        },
    };

    return (
        <Box
            sx={{
                height: "100vh",
                display: "flex",
                flexGrow: 1,
            }}
        >
            <Grid container sx={{flexGrow: 1}}>
                <Grid
                    item
                    sm={4}
                    xs={12}
                    container
                    justifyContent="center"
                    alignItems="stretch"
                    direction="column"
                    sx={{
                        display: "flex",
                        flexGrow: 1,
                        backgroundPosition: "center",
                        backgroundRepeat: "no-repeat",
                        backgroundSize: "cover",
                        alignContent: "center",
                    }}
                >
                    <OnboardingLeftPannel/>
                </Grid>
                <Grid
                    item
                    xs={12}
                    md={8}
                    sx={{
                        p: {
                            xs: 4,
                            sm: 6,
                            md: 8,
                        },
                    }}
                >
                    <Box maxWidth="sm">
                        <Typography sx={{mb: 3}} variant="h4">
                            Welcome {user.firstName} !
                        </Typography>
                        <Stepper
                            activeStep={activeStep}
                            orientation="vertical"
                            sx={{
                                "& .MuiStepConnector-line": {
                                    ml: 1,
                                    borderLeftColor: "divider",
                                    borderLeftWidth: 2,
                                },
                            }}
                        >
                            {Object.values(steps).map((step, index) => (
                                <Step key={step.label}>
                                    <StepLabel StepIconComponent={StepIcon}>
                                        <Typography sx={{ml: 2}} variant="overline">
                                            {step.label}
                                        </Typography>
                                    </StepLabel>
                                    <StepContent
                                        sx={{
                                            py: activeStep === index && 4,
                                            ml: "20px",
                                            borderLeftColor: "divider",
                                            borderLeftWidth: 2,
                                        }}
                                    >
                                        {step.content}
                                    </StepContent>
                                </Step>
                            ))}
                        </Stepper>
                    </Box>
                </Grid>
            </Grid>
        </Box>
    );
};

export default Onboarding;
