import React from "react";
import {Box, Typography} from "@mui/material";
import OnboardingLogo from "./onboarding-logo";

export const OnboardingLeftPannel = () => {
    return <Box sx={{ display: 'flex', flexDirection: 'column', p:2, alignItems: 'center' }}>
        <OnboardingLogo width={200} />
        <Typography variant="h2" textAlign="center" sx={{ py: 3 }}>
            Hey Butler
        </Typography>
        <Typography variant="h6" textAlign="center">
            Welcome to the butler family. We&apos;re here to guide you through
            this incredible journey.
        </Typography>
    </Box>
}

export default OnboardingLeftPannel;