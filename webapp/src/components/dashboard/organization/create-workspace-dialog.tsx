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
    FormHelperText,
    TextField
} from '@mui/material';
import {Api, GoogleRpcStatus, V1WorkspaceResponse} from "../../../api";
import {useWorkspace} from "../../../hooks/use-workspace";
import {useDispatch} from "react-redux";

export const CreateWorkspaceDialog: FC = () => {
    const dispatch = useDispatch();
    const [open, setOpen] = useState(false);
    const {organizationId} = useWorkspace();

    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const formik = useFormik({
        initialValues: {
            name: '',
            description: '',
            submit: null
        },
        validationSchema: Yup.object({
            name: Yup.string().max(255).required('Name is required'),
            description: Yup.string()
        }),
        onSubmit: async (values, { setErrors }): Promise<void> => {
            try {
                await Api.v1.usersServiceCreateWorkspace({
                    organizationId,
                    ...values
                })
                setOpen(false);
            } catch (err) {
                const rpcError: GoogleRpcStatus = err.response?.data;
                setErrors({submit: rpcError.message})
            }
        }
    });

    return (
        <div>
            <Button variant="contained" onClick={handleClickOpen}>
                Add Workspace
            </Button>
            <Dialog open={open} onClose={handleClose}>
                <DialogTitle>Create a new workspace</DialogTitle>
                <form noValidate onSubmit={formik.handleSubmit}>
                    <DialogContent sx={{paddingTop: 0}}>
                        <DialogContentText>
                            {formik.errors.submit &&
                            <Alert severity="error">{formik.errors.submit}</Alert>
                            }
                        </DialogContentText>
                        <TextField
                            autoFocus
                            error={Boolean(formik.touched.name && formik.errors.name)}
                            fullWidth
                            helperText={formik.touched.name && formik.errors.name}
                            label="Workspace Name"
                            margin="normal"
                            name="name"
                            onBlur={formik.handleBlur}
                            onChange={formik.handleChange}
                            type="text"
                            value={formik.values.name}
                        />
                        <TextField
                            error={Boolean(formik.touched.description && formik.errors.description)}
                            fullWidth
                            helperText={formik.touched.description && formik.errors.description}
                            label="Description"
                            margin="normal"
                            name="description"
                            onBlur={formik.handleBlur}
                            onChange={formik.handleChange}
                            type="text"
                            value={formik.values.description}
                        />
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={handleClose}>Cancel</Button>
                        <Button
                            disabled={formik.isSubmitting}
                            type="submit"
                            variant="contained"
                        >
                            Create
                        </Button>
                    </DialogActions>
                </form>
            </Dialog>
        </div>
    );
};
