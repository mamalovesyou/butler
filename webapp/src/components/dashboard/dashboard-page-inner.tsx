import React from "react";
import {styled} from "@mui/material/styles";

export const DashboardPageInner = styled(
    'div',
    { shouldForwardProp: (prop) => prop !== 'open' && prop !== 'drawerWidth' }
)<{ open?: boolean; drawerWidth?: number }>(
    ({ theme, open, drawerWidth }) => ({
        flexGrow: 1,
        overflow: 'hidden',
        p: 0,
        zIndex: 1,
        [theme.breakpoints.up('lg')]: {
            marginRight: -drawerWidth
        },
        transition: theme.transitions.create('margin', {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.leavingScreen
        }),
        ...(open && {
            [theme.breakpoints.up('lg')]: {
                marginRight: 0
            },
            transition: theme.transitions.create('margin', {
                easing: theme.transitions.easing.easeOut,
                duration: theme.transitions.duration.enteringScreen
            })
        })
    })
);

DashboardPageInner.defaultProps = {
    drawerWidth: 500
}

export default DashboardPageInner;