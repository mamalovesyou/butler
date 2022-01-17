import {useState, FC, useEffect} from 'react';
import {Box, CircularProgress, Container, Typography} from '@mui/material';
import {useParams} from "react-router-dom";
import {ACCOUNT_ROOT_PATH, DASHBOARD_ROOT_PATH} from "../../../routes";
import {push} from "redux-first-history";
import {useDispatch} from "react-redux";
import {Api} from "../../../api";
import {notificationsActions} from '../../../features/notifications'
import {TeamMembersTabs} from "../../../components/dashboard/organization/team-members-tabs";
import {useWorkspace} from "../../../hooks/use-workspace";

const WorkspaceDetail: FC = () => {

    const dispatch = useDispatch();
    const {workspaceId} = useParams();
    const [loading, setLoading] = useState(true);
    const [workspace, setWorkspace] = useState(null);

    const [organizationMembers, setOrganizationMembers] = useState([]);
    const {organizationId, organizations} = useWorkspace();

    useEffect(() => {
        if (organizationId) {
            setOrganizationMembers(organizations[organizationId].members);
        }
    }, [organizationId]);

    useEffect(() => {
        if (workspaceId) {
            try {
                Api.v1.usersServiceGetWorkspace({workspaceId}).then((response) => {
                    setWorkspace(response.data.workspace);
                    setLoading(false);
                });
            } catch (err) {
                dispatch(
                    notificationsActions.createAlert({
                        message: err.response?.message || String(err),
                        type: "error",
                    })
                );
                dispatch(push(`/${ACCOUNT_ROOT_PATH}`));
            }
        }
    }, [workspaceId]);


    return (
        <>
            <Box
                component="main"
                sx={{
                    flexGrow: 1,
                    py: 8
                }}
            >
                <Container maxWidth="md">
                    {loading
                        ? <CircularProgress/>
                        : <>
                            <Typography sx={{paddingBottom: 2}} variant="h5">Workspace: {workspace.name}</Typography>
                            <TeamMembersTabs mode="workspace"
                                             workspaceId={workspaceId}
                                             members={workspace.members.concat(organizationMembers)}
                                             invitations={workspace.invitations}
                            /></>
                    }
                </Container>
            </Box>
        </>
    );
};

export default WorkspaceDetail;
