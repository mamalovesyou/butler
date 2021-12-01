import type { FC } from 'react';
import * as Yup from 'yup';
import { useFormik } from 'formik';
import {
  Box,
  Button,
  Checkbox,
  FormHelperText,
  TextField,
  Typography,
  Link
} from '@mui/material';
import { useDispatch } from 'react-redux';
import { signupRequest } from '../../features/auth';

export const JWTRegister: FC = (props) => {
  const dispatch = useDispatch();
  const formik = useFormik({
    initialValues: {
      email: '',
      lastName: '',
      firstName: '',
      password: '',
      policy: false,
      submit: null
    },
    validationSchema: Yup.object({
      firstName: Yup.string().max(255).required('Firstname is required'),
      lastName: Yup.string().max(255).required('Lastname is required'),
      email: Yup.string()
        .email('Must be a valid email')
        .max(255)
        .required('Email is required'),
      password: Yup.string().min(7).max(255).required('Password is required'),
      policy: Yup.boolean().oneOf([true], 'This field must be checked')
    }),
    onSubmit: async (values, helpers): Promise<void> => {
      dispatch(
        signupRequest({
          firstName: values.firstName,
          lastName: values.lastName,
          email: values.email,
          password: values.password
        })
      );

      // TODO: Display errors
    }
  });

  return (
    <form noValidate onSubmit={formik.handleSubmit} {...props}>
      <TextField
        error={Boolean(formik.touched.firstName && formik.errors.firstName)}
        fullWidth
        helperText={formik.touched.firstName && formik.errors.firstName}
        label="Firstname"
        margin="normal"
        name="firstName"
        onBlur={formik.handleBlur}
        onChange={formik.handleChange}
        value={formik.values.firstName}
      />
      <TextField
        error={Boolean(formik.touched.lastName && formik.errors.lastName)}
        fullWidth
        helperText={formik.touched.lastName && formik.errors.lastName}
        label="Lastname"
        margin="normal"
        name="lastName"
        onBlur={formik.handleBlur}
        onChange={formik.handleChange}
        value={formik.values.lastName}
      />
      <TextField
        error={Boolean(formik.touched.email && formik.errors.email)}
        fullWidth
        helperText={formik.touched.email && formik.errors.email}
        label="Email Address"
        margin="normal"
        name="email"
        onBlur={formik.handleBlur}
        onChange={formik.handleChange}
        type="email"
        value={formik.values.email}
      />
      <TextField
        error={Boolean(formik.touched.password && formik.errors.password)}
        fullWidth
        helperText={formik.touched.password && formik.errors.password}
        label="Password"
        margin="normal"
        name="password"
        onBlur={formik.handleBlur}
        onChange={formik.handleChange}
        type="password"
        value={formik.values.password}
      />
      <Box
        sx={{
          alignItems: 'center',
          display: 'flex',
          ml: -1,
          mt: 2
        }}
      >
        <Checkbox
          checked={formik.values.policy}
          name="policy"
          onChange={formik.handleChange}
        />
        <Typography color="textSecondary" variant="body2">
          I have read the{' '}
          <Link component="a" href="#">
            Terms and Conditions
          </Link>
        </Typography>
      </Box>
      {Boolean(formik.touched.policy && formik.errors.policy) && (
        <FormHelperText error>{formik.errors.policy}</FormHelperText>
      )}
      {formik.errors.submit && (
        <Box sx={{ mt: 3 }}>
          <FormHelperText error>{formik.errors.submit}</FormHelperText>
        </Box>
      )}
      <Box sx={{ mt: 2 }}>
        <Button
          disabled={formik.isSubmitting || !formik.values.policy}
          fullWidth
          size="large"
          type="submit"
          variant="contained"
        >
          Register
        </Button>
      </Box>
    </form>
  );
};
