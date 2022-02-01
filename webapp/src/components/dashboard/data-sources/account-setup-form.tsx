import * as Yup from 'yup';
import { useFormik } from 'formik';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import { useEffect, useState } from "react";
import {
    FormControlLabel, FormGroup, FormHelperText,
    Switch,
    TextField,
    Typography,
} from "@mui/material";
import { useDispatch } from "react-redux";
import { updateConnectorConfigRequest } from "../../../features/connectors";
import { V1ProviderAccount } from "../../../api";


interface AccountSetupFormProps {
    connectorId: string;
    provider: string;
    account?: V1ProviderAccount | undefined;
    onBack?: () => void;
}

export const AccountSetupForm = (props: AccountSetupFormProps) => {
    const dispatch = useDispatch();
    const { connectorId, account, provider, onBack } = props;
    const [readOnly, setReadOnly] = useState(false);

    const formik = useFormik({
        initialValues: {
            id: '',
            name: '',
            isTest: false
        },
        validationSchema: Yup.object({
            id: Yup.string()
                .max(255)
                .required('Account ID is required'),
            name: Yup.string()
                .max(255)
                .required('Account Name is required'),
            isTest: Yup.boolean()
                .required('You must specify if this is a test account'),
        }),
        onSubmit: async (values): Promise<void> => {
            dispatch(updateConnectorConfigRequest({
                workspaceConnectorId: connectorId,
                accountId: values.id,
                accountName: values.name,
                isTestAccount: values.isTest
            }))
        }
    });

    useEffect(() => {
        if (account) {
            setReadOnly(true);
        }
    }, [account])

    return (
        <form noValidate onSubmit={formik.handleSubmit}>
            <Box sx={{ p: 2 }}>

                <Box sx={{ display: 'flex', flexDirection: 'row', my: 2 }}>
                    <Typography sx={{ px: 2 }} variant="h6">Provider:</Typography>
                    <Typography variant="h6">{provider}</Typography>
                </Box>
                <TextField
                    error={Boolean(formik.touched.id && formik.errors.id)}
                    fullWidth
                    disabled={readOnly}
                    helperText={formik.touched.id && formik.errors.id}
                    label="Account ID"
                    margin="normal"
                    name="id"
                    onBlur={formik.handleBlur}
                    onChange={formik.handleChange}
                    value={formik.values.id}
                />
                <TextField
                    error={Boolean(formik.touched.name && formik.errors.name)}
                    fullWidth
                    disabled={readOnly}
                    helperText={formik.touched.name && formik.errors.name}
                    label="Account Name"
                    margin="normal"
                    name="name"
                    onBlur={formik.handleBlur}
                    onChange={formik.handleChange}
                    value={formik.values.name}
                />
                <FormGroup>
                    <FormControlLabel
                        control={
                            <Switch disabled={readOnly} checked={formik.values.isTest} onChange={formik.handleChange} name="isTest" />
                        }
                        label="Test Account"
                    />
                </FormGroup>
                {formik.errors.submit && (
                    <Box sx={{ mt: 3 }}>
                        <FormHelperText error>{formik.errors.submit}</FormHelperText>
                    </Box>
                )}
            </Box>
            <Box sx={{ display: 'flex', flexDirection: 'row', pt: 2 }}>
                <Button
                    disabled={formik.isSubmitting}
                    variant="outlined"
                    color="primary"
                    onClick={onBack}
                    sx={{ mr: 1 }}
                >
                    Back
                </Button>
                <Box sx={{ flex: '1 1 auto' }} />
                <Button
                    type="submit"
                    color="success"
                    variant="contained"
                >
                    Complete
                </Button>
            </Box>
        </form>
    );
};