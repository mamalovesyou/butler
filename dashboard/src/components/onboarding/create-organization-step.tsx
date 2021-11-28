import type { FC } from 'react';
import * as Yup from 'yup';
import PropTypes from 'prop-types';
import { Box, Button, FormHelperText, TextField, Typography } from '@mui/material';
import { useFormik } from 'formik';
import { ArrowRight as ArrowRightIcon } from '../../icons/arrow-right';
import { useDispatch } from 'react-redux';
import { createOrganizationRequest } from '../../features/workspace';

interface CreateOrganizationStepProps {
  onNext?: () => void;
  onBack?: () => void;
}

export const CreateOrganizationStep: FC<CreateOrganizationStepProps> = (props) => {
  const { onBack, onNext, ...other } = props;

  const dispatch = useDispatch();

  const formik = useFormik({
    initialValues: {
      name: '',
      role: '',
      submit: null
    },
    validationSchema: Yup.object({
      name: Yup
        .string()
        .max(255)
        .required('Organization name is required'),
      role: Yup
        .string()
    }),
    onSubmit: async (values, helpers): Promise<void> => {

      dispatch(createOrganizationRequest({
        name: values.name,
        userRole: values.role
      }))

      // TODO: Display errors
    }
  });



  return (
    <div {...other}>
      <Typography variant="h6">
        Create an Organization
      </Typography>
      <Typography
        color="textSecondary"
        sx={{ mt: 2 }}
        variant="body2"
      >
        What's your company name ? And what is your role or job title ?
      </Typography>
      <form
        noValidate
        onSubmit={formik.handleSubmit}
        {...props}
      >
        <Box sx={{ mt: 3 }}>
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
          {formik.errors.submit && (
            <Box sx={{ mt: 3 }}>
              <FormHelperText error>
                {formik.errors.submit}
              </FormHelperText>
            </Box>
          )}
        </Box>
        <Box sx={{ pt: 2 }}>
          <Button
            endIcon={(<ArrowRightIcon fontSize="small" />)}
            type="submit"
            variant="contained"
            disabled={formik.isSubmitting}
          >
            Continue
          </Button>
        </Box>
      </form>
    </div>
  );
};

CreateOrganizationStep.propTypes = {
  onBack: PropTypes.func,
  onNext: PropTypes.func
};