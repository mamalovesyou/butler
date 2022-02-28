import {useEffect, useState} from "react";
import * as Yup from 'yup';
import {useFormik} from 'formik';
import {Box, Button, CircularProgress, TextField, Alert} from '@mui/material';
import {ArrowRight as ArrowRightIcon} from "../../../icons/arrow-right";
import {Api, V1Organization} from "../../../api";

interface CreateOrganizationFormProps {
    onSuccess?: (workspace: V1Organization) => void;
}

export const CreateOrganizationForm = (props: CreateOrganizationFormProps) => {
    const { onSuccess, ...other } = props;
    const [error, setError] = useState('');

    const formik = useFormik({
        initialValues: {
            name: '',
            role: '',
        },
        validationSchema: Yup.object({
            name: Yup.string().max(255).required('Organization name is required'),
            role: Yup.string()
        }),
        onSubmit: async (values, helpers): Promise<void> => {
            console.log("clicked submit")
            try {
                const response = await Api.v1.usersServiceCreateOrganization({
                    name: values.name,
                    userRole: values.role
                });
                if (onSuccess) onSuccess(response.data.organization);
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
                label="Organization Name"
                margin="normal"
                name="name"
                onBlur={formik.handleBlur}
                onChange={formik.handleChange}
                type="text"
                value={formik.values.name}
            />
            <TextField
                error={Boolean(formik.touched.role && formik.errors.role)}
                fullWidth
                helperText={formik.touched.role && formik.errors.role}
                label="Role"
                margin="normal"
                name="role"
                onBlur={formik.handleBlur}
                onChange={formik.handleChange}
                type="text"
                value={formik.values.role}
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

export default CreateOrganizationForm;
