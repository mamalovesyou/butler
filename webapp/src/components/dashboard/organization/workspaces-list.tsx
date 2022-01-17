import type {FC} from 'react';
import {
    Avatar,
    Box,
    Button,
    Card,
    CardContent, Chip,
    Divider,
    IconButton,
    InputAdornment, Tab,
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableRow,
    TextField,
    Typography
} from '@mui/material';
import {Scrollbar} from '../../scrollbar';
import {SeverityPill} from '../../severity-pill';
import {useWorkspace} from "../../../hooks/use-workspace";
import {useEffect, useState} from "react";
import {V1Invitation, V1UserMember, V1Workspace} from "../../../api";
import {ChevronRight} from "@mui/icons-material";
import {CreateWorkspaceDialog} from "./create-workspace-dialog";
import {useNavigate} from "react-router-dom";
import {ORGANIZATION_WORKSPACES_PATH} from "../../../routes";

export const WorkspacesList: FC = () => {

    const navigate = useNavigate();
    const [workspaces, setWorkspaces] = useState<V1Workspace[]>([]);
    const {organizationId, workspaceId, organizations} = useWorkspace();


    useEffect(() => {
        if (organizationId) {
            const orga = organizations[organizationId];
            setWorkspaces(orga.workspaces);
        }
    }, [organizationId])

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
                    <Typography variant="h6">Workspaces</Typography>
                </div>

                <CreateWorkspaceDialog />
            </Box>
        </CardContent>
        <Scrollbar>
            <Box sx={{width: '100%', typography: 'body1', p: 0}}>
                <Table sx={{minWidth: 400}}>
                    <TableHead>
                        <TableRow>
                            <TableCell>Name</TableCell>
                            <TableCell/>
                            <TableCell/>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {workspaces.map((ws: V1Workspace) => <TableRow key={ws.id} hover>
                                <TableCell>
                                    <Box
                                        sx={{
                                            alignItems: 'center',
                                            display: 'flex'
                                        }}
                                    >
                                        <div>
                                            <Typography
                                                variant="subtitle2">{ws.name}
                                            </Typography>
                                            <Typography color="textSecondary" variant="body2">
                                                {ws.description}
                                            </Typography>
                                        </div>
                                    </Box>
                                </TableCell>
                                <TableCell>
                                    { ws.id === workspaceId && <SeverityPill>Current workspace</SeverityPill>}
                                </TableCell>
                                <TableCell align="right">
                                    <IconButton aria-label="delete" size="large" onClick={() => navigate(ws.id)}>
                                        <ChevronRight/>
                                    </IconButton>
                                </TableCell>
                            </TableRow>
                        )}

                    </TableBody>
                </Table>
            </Box>
        </Scrollbar>
    </Card>
};

export default WorkspacesList;
