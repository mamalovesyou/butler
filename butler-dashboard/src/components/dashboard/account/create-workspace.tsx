import { FC } from 'react';
import * as Yup from 'yup';
import { useFormik } from 'formik';
import { Box, Button, FormHelperText, TextField } from '@mui/material';
import { useDispatch } from 'react-redux';
import { createWorkspaceRequest } from '../../../features/workspace';

export const CreateWorkspace: FC = (props) => {

  const dispatch = useDispatch();

  const formik = useFormik({
    initialValues: {
      name: '',
      description: '',
      submit: null
    },
    validationSchema: Yup.object({
      name: Yup
        .string()
        .max(255)
        .required('Name is required'),
        description: Yup
        .string()
    }),
    onSubmit: async (values, helpers): Promise<void> => {
      
      // dispatch(createWorkspaceRequest({
      //   name: values.name,
      //   description: values.description
      // }))

      // TODO: Display errors
    }
  });

  return (
    <form
      noValidate
      onSubmit={formik.handleSubmit}
      {...props}
    >
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
        label="Password"
        margin="normal"
        name="password"
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
      <Box sx={{ mt: 2 }}>
        <Button
          disabled={formik.isSubmitting}
          fullWidth
          size="large"
          type="submit"
          variant="contained"
        >
          Create
        </Button>
      </Box>
    </form>
  );
};
