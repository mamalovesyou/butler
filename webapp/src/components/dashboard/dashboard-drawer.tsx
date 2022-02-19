import type { FC, MutableRefObject } from 'react';
import {
    Box,
    Drawer,
    IconButton,
    Theme,
    Typography, useMediaQuery
} from '@mui/material';
import CloseIcon from '@mui/icons-material/Close';
import { styled } from '@mui/material/styles';
import {DASHBOARD_NAVBAR_HEIGHT} from "../../theme/constants";


interface DashboardDrawerProps {
    containerRef?: MutableRefObject<HTMLDivElement>;
    open?: boolean;
    onClose?: () => void;
    title?: string;
    width?: number;
}

const DashboardDrawerDesktop = styled(Drawer, {
    shouldForwardProp: (prop) => prop !== "width",
})<{ width?: number; }>(({ width }) => ({
    width: width,
    flexShrink: 0,
    '& .MuiDrawer-paper': {
        position: 'relative',
        width: width
    }
}));

const DashboardDrawerMobile = styled(Drawer, {
    shouldForwardProp: (prop) => prop !== "width",
})<{ width?: number; }>(({ width }) => ({
    flexShrink: 0,
    maxWidth: '100%',
    height: `calc(100% - ${DASHBOARD_NAVBAR_HEIGHT}px)`,
    width: width,
    '& .MuiDrawer-paper': {
        height: `calc(100% - ${DASHBOARD_NAVBAR_HEIGHT}px)`,
        maxWidth: '100%',
        top: DASHBOARD_NAVBAR_HEIGHT,
        width: width
    }
}));

export const DashboardDrawer: FC<DashboardDrawerProps> = (props) => {
    const { containerRef, width, onClose, open, title, children, ...other } = props;

    const lgUp = useMediaQuery((theme: Theme) => theme.breakpoints.up('lg'));

    // The reason for doing this, is that the persistent drawer has to be rendered, but not it's
    // content if an order is not passed.
    const content = open
        ? (
            <>
                <Box
                    sx={{
                        alignItems: 'center',
                        backgroundColor: 'primary.main',
                        color: 'primary.contrastText',
                        display: 'flex',
                        justifyContent: 'space-between',
                        px: 3,
                        py: 2
                    }}
                >
                    <Typography
                        color="inherit"
                        variant="h6"
                    >
                        {title}
                    </Typography>
                    <IconButton
                        color="inherit"
                        onClick={onClose}
                    >
                        <CloseIcon />
                    </IconButton>
                </Box>
                <Box sx={{ p: 2 }}>
                    {children}
                </Box>
            </>
        ) : null;

    return lgUp ?
            <DashboardDrawerDesktop
                anchor="right"
                open={open}
                SlideProps={{ container: containerRef?.current }}
                variant="persistent"
                width={width}
                {...other}
            >
                {content}
            </DashboardDrawerDesktop>
        : <DashboardDrawerMobile
            anchor="right"
            ModalProps={{ container: containerRef?.current }}
            onClose={onClose}
            open={open}
            SlideProps={{ container: containerRef?.current }}
            variant="temporary"
            width={width}
            {...other}
        >
            {content}
        </DashboardDrawerMobile>
};

DashboardDrawer.defaultProps = {
    onClose: () => {},
    open: false,
    title: "",
    width: 500,
};

export default DashboardDrawer;