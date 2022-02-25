import {useState} from "react";
import * as Yup from 'yup';
import {useFormik} from 'formik';
import {Box, Button, CircularProgress, TextField, Alert} from '@mui/material';
import {ArrowRight as ArrowRightIcon} from "../../../icons/arrow-right";
import {Api, V1Workspace} from "../../../api";

interface CreateWorkspaceFormProps {
    organizationId: string;
    onSuccess?: (workspace: V1Workspace) => void;
}

export const CreateWorkspaceForm = (props: CreateWorkspaceFormProps) => {
    const { organizationId, onSuccess, ...other } = props;
    const [error, setError] = useState('');

    const formik = useFormik({
        initialValues: {
            name: '',
            description: '',
        },
        validationSchema: Yup.object({
            name: Yup.string().max(255).required('Name is required'),
            description: Yup.string()
        }),
        onSubmit: async (values, helpers): Promise<void> => {
            try {
                const response = await Api.v1.usersServiceCreateWorkspace({
                    organizationId,
                    workspace: values
                });
                if (onSuccess) onSuccess(response.data.workspace);
            } catch (e) {
                if (e.response) {
                    setError(`Error ${e.response.status} - ${e.response.data}`)
                } else {
                    setError(String(e));
                }
            }
        }
    });

    return <form noValidate onSubmit={formik.handleSubmit} {...other}>
        { error != '' ? <Box sx={{mt: 3}}>
            <Alert severity="error">{error}</Alert>
        </Box> : null }
        <Box sx={{mt: 3}}>
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
                error={Boolean(
                    formik.touched.description && formik.errors.description
                )}
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
        </Box>
        <Box sx={{pt: 2}}>
            { formik.isSubmitting ? <CircularProgress /> : <Button
                type="submit"
                variant="contained"
                disabled={!formik.dirty || !formik.isValid || formik.isValidating}
            >
                Create
            </Button> }
        </Box>
    </form>
};

export default CreateWorkspaceForm;
