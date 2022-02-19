import {ReactNode, useMemo, useRef, useState} from 'react';
import type {FC} from 'react';
import PropTypes from 'prop-types';
import {Box, Divider, Drawer, Typography, useMediaQuery} from '@mui/material';
import type {Theme} from '@mui/material';
import {Selector as SelectorIcon} from '../../icons/selector';
import {Scrollbar} from '../scrollbar';
import {DashboardSidebarSection} from './dashboard-sidebar-section';
import {WorkspacePopover} from './workspace-popover';
import {useCurrentWorkspace} from '../../hooks/use-workspace';
import {DASHBOARD_SIDEBAR_WIDTH} from "../../theme/constants";
import AssignmentIcon from '@mui/icons-material/Assignment';
import {ORGANIZATION_ROOT_PATH, UTMS_ROOT_PATH} from "../../routes";

interface DashboardSidebarProps {
    onClose: () => void;
    open: boolean;
}

interface Item {
    title: string;
    children?: Item[];
    chip?: ReactNode;
    icon?: ReactNode;
    path?: string;
}

interface Section {
    title: string;
    path: string;
    items: Item[];
}

const getSections = (): Section[] => [
    {
        title: 'Analytics',
        path: '/',
        items: [
            {
                title: 'Dashboard',
                path: ''
            },
        ]
    },
    {
        title: 'Tracking',
        path: '/',
        items: [
            {
                title: 'UTM Automation',
                path: UTMS_ROOT_PATH,
            }
        ]
    },
    {
        title: 'Configuration',
        path: '/',
        items: [
            {
                title: 'Organization',
                path: ORGANIZATION_ROOT_PATH
            },
            {
                title: 'Data Sources',
                path: 'data-sources'
            },
        ]
    }
];

export const DashboardSidebar: FC<DashboardSidebarProps> = (props) => {
    const {onClose, open} = props;
    const lgUp = useMediaQuery((theme: Theme) => theme.breakpoints.up('lg'), {
        noSsr: true
    });
    const {workspace} = useCurrentWorkspace();
    const sections = useMemo(() => getSections(), []);
    const organizationsRef = useRef<HTMLButtonElement | null>(null);
    const [openOrganizationsPopover, setOpenOrganizationsPopover] =
        useState<boolean>(false);

    const toggleOrganizationsPopover = (): void => {
        setOpenOrganizationsPopover(!openOrganizationsPopover);
    };

    const content = (
        <>
            <Scrollbar
                sx={{
                    height: '100%',
                    '& .simplebar-content': {
                        height: '100%'
                    }
                }}
            >
                <Box
                    sx={{
                        display: 'flex',
                        flexDirection: 'column',
                        height: '100%'
                    }}
                >
                    <div>
                        <Box>
                            <Box
                                onClick={toggleOrganizationsPopover}
                                ref={organizationsRef}
                                sx={{
                                    alignItems: 'center',
                                    backgroundColor: 'rgba(255, 255, 255, 0.04)',
                                    cursor: 'pointer',
                                    display: 'flex',
                                    justifyContent: 'space-between',
                                    px: 3,
                                    py: '11px',
                                    borderRadius: 1,
                                    minHeight: 64
                                }}
                            >
                                <div>
                                    <Typography color="inherit" variant="subtitle1">
                                        {workspace?.name}
                                    </Typography>
                                </div>
                                <SelectorIcon
                                    sx={{
                                        color: 'neutral.500',
                                        width: 14,
                                        height: 14
                                    }}
                                />
                            </Box>
                        </Box>
                    </div>
                    <Divider
                        sx={{
                            borderColor: '#2D3748' // dark divider
                        }}
                    />
                    <Box sx={{flexGrow: 1}}>
                        {sections.map((section) => (
                            <DashboardSidebarSection
                                key={section.title}
                                sx={{
                                    mt: 2,
                                    '& + &': {
                                        mt: 2
                                    }
                                }}
                                {...section}
                            />
                        ))}
                    </Box>
                    <Divider
                        sx={{
                            borderColor: '#2D3748' // dark divider
                        }}
                    />
                </Box>
            </Scrollbar>
            <WorkspacePopover
                anchorEl={organizationsRef.current}
                onClose={toggleOrganizationsPopover}
                open={openOrganizationsPopover}
            />
        </>
    );

    if (lgUp) {
        return (
            <Drawer
                anchor="left"
                open
                PaperProps={{
                    sx: {
                        backgroundColor: 'neutral.900',
                        borderRightColor: 'divider',
                        borderRightStyle: 'solid',
                        borderRightWidth: (theme) =>
                            theme.palette.mode === 'dark' ? 1 : 0,
                        color: '#FFFFFF',
                        width: DASHBOARD_SIDEBAR_WIDTH
                    }
                }}
                variant="permanent"
            >
                {content}
            </Drawer>
        );
    }

    return (
        <Drawer
            anchor="left"
            onClose={onClose}
            open={open}
            PaperProps={{
                sx: {
                    backgroundColor: 'neutral.900',
                    color: '#FFFFFF',
                    width: DASHBOARD_SIDEBAR_WIDTH
                }
            }}
            sx={{zIndex: (theme) => theme.zIndex.appBar + 100}}
            variant="temporary"
        >
            {content}
        </Drawer>
    );
};

DashboardSidebar.propTypes = {
    onClose: PropTypes.func,
    open: PropTypes.bool
};
