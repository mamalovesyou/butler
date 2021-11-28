import React, { FC, useEffect, useState } from 'react';
import {
  Avatar,
  Box,
  Button,
  Card,
  Grid,
  Step,
  StepContent,
  StepLabel,
  Stepper,
  Typography,
} from '@mui/material';
import type { StepIconProps } from '@mui/material';
import { useWorkspace } from '../hooks/use-workspace';
import { useOnboarding } from '../hooks/use-onboarding';
import { useDispatch } from 'react-redux';
import { CreateWorkspaceStep } from '../components/onboarding/create-workspace-step';
import { CreateOrganizationStep } from '../components/onboarding/create-organization-step';
import { Check as CheckIcon } from '../icons/check';
import { useAuth } from '../hooks/use-auth';
import { OnboardingStep, setOnboardingStep } from '../features/onboarding';
import { ConnectDataSourceStep } from '../components/onboarding/connect-data-source';

const StepIcon: React.FC<StepIconProps> = (props) => {
  const { active, completed, icon } = props;

  const highlight = active || completed;

  return (
    <Avatar
      sx={{
        backgroundColor: highlight && 'secondary.main',
        color: highlight && 'secondary.contrastText',
        height: 40,
        width: 40,
      }}
      variant="rounded"
    >
      {
        completed
          ? <CheckIcon fontSize="small" />
          : icon
      }
    </Avatar>
  );
};

const Onboarding: React.FC = () => {
  const { activeStep } = useOnboarding();
  const { user } = useAuth();
  const dispatch = useDispatch();
  const { workspaceId, organizationId } = useWorkspace();

  useEffect(() => {
    if (organizationId) {
      if (workspaceId) {
        dispatch(setOnboardingStep(OnboardingStep.CONNECT_DATA_SOURCE));
      } else {
        dispatch(setOnboardingStep(OnboardingStep.CREATE_WORKSPACE));
      }
    }
  }, [organizationId, workspaceId]);

  const steps = {
    [OnboardingStep.CREATE_ORGANIZATION]: {
      label: 'Organization',
      content: (
        <CreateOrganizationStep />
      ),
    },
    [OnboardingStep.CREATE_WORKSPACE]: {
      label: 'Workspace',
      content: (
        <CreateWorkspaceStep />
      ),
    },
    [OnboardingStep.CONNECT_DATA_SOURCE]: {
      label: 'Connect a source',
      content: (
        <ConnectDataSourceStep />
      ),
    },
  };

  return (
    <Box
      sx={{
        height: '100vh',
        display: 'flex',
        flexGrow: 1,
      }}
    >
      <Grid
        container
        sx={{ flexGrow: 1 }}
      >
        <Grid
          item
          sm={4}
          xs={12}
          container
          justifyContent="center"
          alignItems="stretch"
          direction="column"
          sx={{
            display: 'flex',
            flexGrow: 1,
            backgroundColor: 'red',
            backgroundPosition: 'center',
            backgroundRepeat: 'no-repeat',
            backgroundSize: 'cover',
            alignContent: 'center',

          }}
        >
          <Typography variant="h2" textAlign="center" sx={{ py: 3 }}>
            Hey Butler
          </Typography>
          <Typography variant="h6" textAlign="center">
            Welcome to the butler family. We&apos;re here to guide you through this incredible journey.
          </Typography>
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
            <Typography
              sx={{ mb: 3 }}
              variant="h4"
            >
              Welcome
              {' '}
              {user.firstName}
              {' '}
              !
            </Typography>
            <Stepper
              activeStep={activeStep}
              orientation="vertical"
              sx={{
                '& .MuiStepConnector-line': {
                  ml: 1,
                  borderLeftColor: 'divider',
                  borderLeftWidth: 2,
                },
              }}
            >
              {Object.values(steps).map((step, index) => (
                <Step key={step.label}>
                  <StepLabel StepIconComponent={StepIcon}>
                    <Typography
                      sx={{ ml: 2 }}
                      variant="overline"
                    >
                      {step.label}
                    </Typography>
                  </StepLabel>
                  <StepContent
                    sx={{
                      py: (activeStep === index) && 4,
                      ml: '20px',
                      borderLeftColor: 'divider',
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
