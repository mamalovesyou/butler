import {Fragment, useState} from 'react';
import type {FC} from 'react';
import PropTypes from 'prop-types';
import {
    Box,
    IconButton,
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableRow,
    Typography
} from '@mui/material';
import {ChevronDown as ChevronDownIcon} from '../../../icons/chevron-down';
import {ChevronRight as ChevronRightIcon} from '../../../icons/chevron-right';
import {DotsHorizontal as DotsHorizontalIcon} from '../../../icons/dots-horizontal';
import {Scrollbar} from '../../scrollbar';
import {SeverityPill} from '../../severity-pill';
import {styled} from "@mui/material/styles";


interface CampaignListTableProps {
    campaigns: any[];
    onCampaignOpen: (campaignId: string) => void;
    sx?: object;
}

export const CampaignListInner = styled(
    'div',
    { shouldForwardProp: (prop) => prop !== 'open' }
)<{ open?: boolean; }>(
    ({ theme, open }) => ({
        flexGrow: 1,
        overflow: 'hidden',
        p: 0,
        zIndex: 1,
        [theme.breakpoints.up('lg')]: {
            marginRight: -500
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

export const CampaignListTable: FC<CampaignListTableProps> = (props) => {
    const {
        onCampaignOpen,
        campaigns,
        ...other
    } = props;

    const [openCampaign, setOpenCampaign] = useState<string>(null);

    const handleOpenCampaign = (campaignId: string): void => {
        setOpenCampaign((prevValue) => (prevValue === campaignId ? null : campaignId));
        onCampaignOpen(campaignId);
    };

    const handleUpdateCampaign = (): void => {
        setOpenCampaign(null);
        // toast.success('Campaign updated');
    };

    const handleCancelEdit = (): void => {
        setOpenCampaign(null);
    };

    const handleDeleteCampaign = (): void => {
        // toast.error('Campaign cannot be deleted');
    };

    return (
        <Box {...other}>
            <Scrollbar>
                <Table sx={{minWidth: '100%'}}>
                    <TableHead>
                        <TableRow>
                            <TableCell/>
                            <TableCell>
                                Campaigns
                            </TableCell>
                            <TableCell>
                                Status
                            </TableCell>
                            <TableCell align="right">
                                Actions
                            </TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {campaigns.map((campaign) => {
                            const open = campaign.id === openCampaign;
                            return (
                                <Fragment key={campaign.id}>
                                    <TableRow
                                        hover
                                        key={campaign.id}
                                    >
                                        <TableCell
                                            padding="checkbox"
                                            sx={{
                                                ...(open && {
                                                    position: 'relative',
                                                    '&:after': {
                                                        position: 'absolute',
                                                        content: '" "',
                                                        top: 0,
                                                        left: 0,
                                                        backgroundColor: 'primary.main',
                                                        width: 5,
                                                        height: 'calc(100% + 1px)'
                                                    }
                                                })
                                            }}
                                        >
                                            <IconButton onClick={() => handleOpenCampaign(campaign.id)}>
                                                {
                                                    open
                                                        ? <ChevronDownIcon fontSize="small"/>
                                                        : <ChevronRightIcon fontSize="small"/>
                                                }
                                            </IconButton>
                                        </TableCell>
                                        <TableCell>
                                            <Box
                                                sx={{
                                                    alignItems: 'center',
                                                    display: 'flex'
                                                }}
                                            >
                                                <Typography variant="subtitle2">
                                                    {campaign.name}
                                                </Typography>
                                            </Box>
                                        </TableCell>
                                        <TableCell>
                                            <SeverityPill color={campaign.status === 'published' ? 'success' : 'info'}>
                                                {campaign.status}
                                            </SeverityPill>
                                        </TableCell>
                                        <TableCell align="right">
                                            <IconButton>
                                                <DotsHorizontalIcon fontSize="small"/>
                                            </IconButton>
                                        </TableCell>
                                    </TableRow>
                                    {open && (
                                        <TableRow>
                                            <TableCell rowSpan={1} />
                                            <TableCell
                                                colSpan={7}
                                                sx={{
                                                    p:0,
                                                    position: 'relative',
                                                    '&:after': {
                                                        position: 'absolute',
                                                        content: '" "',
                                                        top: 0,
                                                        left: 0,
                                                        backgroundColor: 'primary.main',
                                                        width: 5,
                                                        height: 'calc(100% + 1px)'
                                                    }
                                                }}
                                            >
                                                <Table>
                                                    <TableHead>
                                                        <TableRow>
                                                            <TableCell>
                                                                Ad Group: 12345678
                                                            </TableCell>
                                                        </TableRow>
                                                    </TableHead>
                                                    <TableBody>
                                                        <TableRow>
                                                            <TableCell>
                                                                <Box
                                                                    sx={{
                                                                        alignItems: 'center',
                                                                        display: 'flex'
                                                                    }}
                                                                >
                                                                    <Typography variant="subtitle2">
                                                                        Ad: 123456
                                                                    </Typography>
                                                                </Box>
                                                            </TableCell>
                                                        </TableRow>
                                                        <TableRow>
                                                            <TableCell>
                                                                <Box
                                                                    sx={{
                                                                        alignItems: 'center',
                                                                        display: 'flex'
                                                                    }}
                                                                >
                                                                    <Typography variant="subtitle2">
                                                                        Ad: 123456
                                                                    </Typography>
                                                                </Box>
                                                            </TableCell>
                                                        </TableRow>
                                                    </TableBody>
                                                </Table>
                                                <Table>
                                                    <TableHead>
                                                        <TableRow>
                                                            <TableCell>
                                                                Ad Group: 12345678
                                                            </TableCell>
                                                        </TableRow>
                                                    </TableHead>
                                                    <TableBody>
                                                        <TableRow>
                                                            <TableCell>
                                                                <Box
                                                                    sx={{
                                                                        alignItems: 'center',
                                                                        display: 'flex'
                                                                    }}
                                                                >
                                                                    <Typography variant="subtitle2">
                                                                        Ad: 123456
                                                                    </Typography>
                                                                </Box>
                                                            </TableCell>
                                                        </TableRow>
                                                        <TableRow>
                                                            <TableCell>
                                                                <Box
                                                                    sx={{
                                                                        alignItems: 'center',
                                                                        display: 'flex'
                                                                    }}
                                                                >
                                                                    <Typography variant="subtitle2">
                                                                        Ad: 123456
                                                                    </Typography>
                                                                </Box>
                                                            </TableCell>
                                                        </TableRow>
                                                    </TableBody>
                                                </Table>
                                            </TableCell>
                                        </TableRow>
                                    )}
                                </Fragment>
                            );
                        })}
                    </TableBody>
                </Table>
            </Scrollbar>
        </Box>
    );
};

CampaignListTable.propTypes = {
    campaigns: PropTypes.array.isRequired,
};
