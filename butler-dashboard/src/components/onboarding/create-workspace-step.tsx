import type { FC } from 'react';
import * as Yup from 'yup';
import PropTypes from 'prop-types';
import { Box, Button, FormHelperText, TextField, Typography } from '@mui/material';
import { useFormik } from 'formik';
import { ArrowRight as ArrowRightIcon } from '../../icons/arrow-right';
import { useDispatch } from 'react-redux';
import { createWorkspaceRequest } from '../../features/workspace';
import { useWorkspace } from '../../hooks/use-workspace';

interface CreateWorkspaceStepProps {
  onNext?: () => void;
  onBack?: () => void;
}

export const CreateWorkspaceStep: FC<CreateWorkspaceStepProps> = (props) => {
  const { onBack, onNext, ...other } = props;

  const { organizationId } = useWorkspace();
  const dispatch = useDispatch();

  const formik = useFormik({
    initialValues: {
      organizationId: organizationId,
      name: '',
      description: '',
      submit: null
    },
    validationSchema: Yup.object({
      organizationId: Yup.string().required(),
      name: Yup
        .string()
        .max(255)
        .required('Name is required'),
      description: Yup
        .string()
    }),
    onSubmit: async (values, helpers): Promise<void> => {

      dispatch(createWorkspaceRequest({
        organizationID: values.organizationId,
        name: values.name,
        description: values.description
      }))

      // TODO: Display errors
    }
  });



  return (
    <div {...other}>
      <Typography variant="h6">
        Create a workspace
      </Typography>
      <Typography
        color="textSecondary"
        sx={{ mt: 2 }}
        variant="body2"
      >
        You can think a workspace as a group. For instance, if you work for multiples companies you will have a different workspace for each of them.
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

CreateWorkspaceStep.propTypes = {
  onBack: PropTypes.func,
  onNext: PropTypes.func
};