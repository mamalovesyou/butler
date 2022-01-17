import {useState} from "react";
import type {FC} from 'react';
import {
    Avatar,
    Box,
    Card,
    CardContent,
    IconButton,
    Tab,
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableRow,
    Typography
} from '@mui/material';
import TabContext from '@mui/lab/TabContext';
import TabList from '@mui/lab/TabList';
import TabPanel from '@mui/lab/TabPanel';
import {DotsHorizontal as DotsHorizontalIcon} from '../../../icons/dots-horizontal';
import {UserCircle as UserCircleIcon} from '../../../icons/user-circle';
import {Scrollbar} from '../../scrollbar';
import {SeverityPill} from '../../severity-pill';
import {V1Invitation, V1UserMember} from "../../../api";
import {InviteMembersDialog} from "./invite-members-dialog";
import {useWorkspace} from "../../../hooks/use-workspace";
import {UserAvatar} from "../../user-avatar";

interface ITeamMembersProps {
    members: V1UserMember[];
    invitations: V1Invitation[];
}

interface IOrganizationTeamMembersProps extends ITeamMembersProps {
    mode: "organization";
}

interface IWorkspaceTeamMembersProps extends ITeamMembersProps {
    mode: "workspace";
    workspaceId: string;
}


export const TeamMembersTabs: FC<IOrganizationTeamMembersProps | IWorkspaceTeamMembersProps> = (props) => {

    const {mode, members, invitations} = props;
    const {organizationId} = useWorkspace();
    const [value, setValue] = useState('1');
    const handleChange = (event: React.SyntheticEvent, newValue: string) => {
        setValue(newValue);
    };

    return <Card>
        <CardContent sx={{p: 2}}>
            <Box
                sx={{
                    p: 0,
                    alignItems: 'center',
                    display: 'flex',
                    justifyContent: 'space-between',
                }}
            >
                <div>
                    <Typography variant="h6">Team members</Typography>
                </div>
                {mode === "organization" ? <InviteMembersDialog mode={mode} organizationId={organizationId}/> :
                    <InviteMembersDialog mode={mode} workspaceId={props.workspaceId}/>}
            </Box>
        </CardContent>
        <Scrollbar>
            <Box sx={{width: '100%', typography: 'body1', p: 0}}>
                <TabContext value={value}>
                    <Box sx={{px: 1}}>
                        <TabList onChange={handleChange}>
                            <Tab label="Team members" value="1"/>
                            <Tab label="Pending invitations" value="2"/>
                        </TabList>
                    </Box>
                    <TabPanel value="1" sx={{px: 0, py: 1}}>
                        <Table sx={{minWidth: 400}}>
                            <TableHead>
                                <TableRow>
                                    <TableCell>Member</TableCell>
                                    <TableCell>Role</TableCell>
                                    <TableCell/>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {members.map((member: V1UserMember) => <TableRow key={member.userId}>
                                        <TableCell>
                                            <Box
                                                sx={{
                                                    alignItems: 'center',
                                                    display: 'flex'
                                                }}
                                            >
                                                <Avatar
                                                    sx={{
                                                        height: 40,
                                                        width: 40,
                                                        mr: 1
                                                    }}
                                                >
                                                    <UserAvatar name={`${member.firstName} ${member.lastName}`}/>
                                                </Avatar>
                                                <Typography sx={{pl: 2}} variant="h6">{member.firstName} {member.lastName}</Typography>
                                            </Box>
                                        </TableCell>
                                        <TableCell>
                                            <SeverityPill>{member.role}</SeverityPill>
                                        </TableCell>
                                        <TableCell align="right">
                                            {(member.role != "owner") && <IconButton>
                                                <DotsHorizontalIcon fontSize="small"/>
                                            </IconButton>}
                                        </TableCell>
                                    </TableRow>
                                )}

                            </TableBody>
                        </Table>
                    </TabPanel>
                    <TabPanel value="2" sx={{px: 0, py: 1}}>
                        {invitations.length > 0 ? <Table sx={{minWidth: 400}}>
                                <TableHead>
                                    <TableRow>
                                        <TableCell>Invitations</TableCell>
                                        <TableCell/>
                                    </TableRow>
                                </TableHead>
                                <TableBody>
                                    {invitations.map((invite: V1Invitation) => <TableRow key={invite.id}>
                                            <TableCell>
                                                <Box
                                                    sx={{
                                                        alignItems: 'center',
                                                        display: 'flex'
                                                    }}
                                                >
                                                    <Typography variant="subtitle2">
                                                        {invite.email}
                                                    </Typography>
                                                </Box>
                                            </TableCell>
                                        </TableRow>
                                    )}

                                </TableBody>
                            </Table> :
                            <Box sx={{p: 1}}>
                                <Typography variant="subtitle2">No pending invitations</Typography>
                            </Box>
                        }
                    </TabPanel>
                </TabContext>
            </Box>
        </Scrollbar>
    </Card>
};
