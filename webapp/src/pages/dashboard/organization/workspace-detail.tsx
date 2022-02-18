import {useState, FC, useEffect} from 'react';
import {Box, CircularProgress, Container, Typography} from '@mui/material';
import {useParams} from "react-router-dom";
import {ORGANIZATION_ROOT_PATH} from "../../../routes";
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
    const {organization} = useWorkspace();

    useEffect(() => {
        (async function() {
            try {
                const response = await Api.v1.usersServiceGetWorkspace({ workspaceId });
                setWorkspace(response.data.workspace);
                setLoading(false);
            } catch (e) {
                console.error(e);
                dispatch(
                    notificationsActions.createAlert({
                        message: e.response?.message || String(e),
                        type: "error",
                    })
                );
                dispatch(push(ORGANIZATION_ROOT_PATH));
            }
        })();
    }, [])

    useEffect(() => {
        if (organization) {
            setOrganizationMembers(organization.members);
        }
    }, [organization]);


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
