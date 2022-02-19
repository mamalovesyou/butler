import {FC, useState} from 'react';
import * as Yup from 'yup';
import {useFormik} from 'formik';
import {
    Alert,
    Box,
    Button,
    Dialog, DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
    IconButton, List, ListItem, ListItemIcon, ListItemText,
    TextField
} from '@mui/material';
import SendIcon from '@mui/icons-material/Send';
import AddIcon from '@mui/icons-material/Add';
import CloseIcon from '@mui/icons-material/Close';
import {LoadingButton} from "@mui/lab";
import {Api, GoogleRpcStatus} from "../../../api";

interface BaseProps {
    mode: "organization" | "workspace";
    organizationId: string;
}

interface IOrganizationModeProps extends BaseProps {
    mode: "organization";
}

interface IWorkspaceModeProps extends BaseProps {
    mode: "workspace";
    workspaceId: string;
}

export const InviteMembersDialog = (props: IOrganizationModeProps | IWorkspaceModeProps) => {

    const [open, setOpen] = useState(false);
    const [loading, setLoading] = useState(false);
    const [emails, setEmails] = useState([]);
    const [error, setError] = useState(null);

    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const removeEmail = (value) => {
        setEmails(emails.filter((e: string) => e !== value))
    }

    const sendInvites = async (): Promise<void> => {
        try {
            setLoading(true);
            let payload: any = { organizationId: props.organizationId, emails };
            if (props.mode === "workspace") {
                payload.workspaceId = props.workspaceId;
            }
            await Api.v1.usersServiceSendBatchInvitations(payload);
            setEmails([]);
            setLoading(false);
            setOpen(false);
        } catch (err) {
            setLoading(false);
            console.log(err)
            const rpcError: GoogleRpcStatus = err.response?.data;
            setError(rpcError.message);
        }
    }


    const formik = useFormik({
        initialValues: {
            email: '',
        },
        validationSchema: Yup.object({
            email: Yup.string().email(),
        }),
        onSubmit: async (values, {resetForm}): Promise<void> => {
            if (values.email != '') {
                setEmails([...emails, values.email]);
                resetForm();
            }
        }
    });

    return (
        <div>
            <Button variant="contained" onClick={handleClickOpen}>
                Add member
            </Button>
            <Dialog open={open} fullWidth
                    maxWidth="sm" onClose={handleClose}>
                <DialogTitle>Invite member</DialogTitle>
                <DialogContent sx={{p: 0}}>
                    <DialogContentText>
                        {error &&
                        <Alert severity="error">{error}</Alert>
                        }
                    </DialogContentText>
                    <List dense>
                        {emails.map((email: string, index: number) => <ListItem key={index}>
                            <ListItemIcon>
                                <IconButton onClick={() => removeEmail(email)}><CloseIcon/></IconButton>
                            </ListItemIcon>
                            <ListItemText
                                primary={email}
                            />
                        </ListItem>)}
                    </List>
                    <form noValidate onSubmit={formik.handleSubmit}>
                        <Box
                            sx={{
                                alignItems: 'center',
                                display: 'flex',
                                p: 1
                            }}
                        >
                            <TextField
                                sx={{m: 2, width: '50vh'}}
                                autoFocus
                                error={Boolean(formik.touched.email && formik.errors.email)}
                                helperText={formik.touched.email && formik.errors.email}
                                label="Email"
                                margin="normal"
                                name="email"
                                onBlur={formik.handleBlur}
                                onChange={formik.handleChange}
                                type="text"
                                value={formik.values.email}
                            />
                            <Button variant="outlined" startIcon={<AddIcon/>} type="submit">
                                Add
                            </Button>
                        </Box>
                    </form>
                </DialogContent>

                <DialogActions>
                    <Button onClick={handleClose}>Cancel</Button>
                    <LoadingButton
                        disabled={emails.length === 0}
                        loading={loading}
                        loadingPosition="start"
                        startIcon={<SendIcon/>}
                        variant="outlined"
                        onClick={sendInvites}
                    >
                        Send
                    </LoadingButton>
                </DialogActions>
            </Dialog>
        </div>
    );
};
