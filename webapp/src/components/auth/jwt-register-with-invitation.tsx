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
import {signupRequest, signupWithInviteRequest} from '../../features/auth';
import {useLocation} from "react-router-dom";
import {useEffect, useState} from "react";
import {Api} from "../../api";

interface IJWTRegisterWithInvitationProps {
    token: string;
    invitationId: string;
}

export const JWTRegisterWithInvitation: FC<IJWTRegisterWithInvitationProps> = (props: IJWTRegisterWithInvitationProps) => {
  const {invitationId, token} = props;
  const dispatch = useDispatch();
  const formik = useFormik({
    initialValues: {
      lastName: '',
      firstName: '',
      password: '',
      policy: false,
      submit: null
    },
    validationSchema: Yup.object({
      firstName: Yup.string().max(255).required('Firstname is required'),
      lastName: Yup.string().max(255).required('Lastname is required'),
      password: Yup.string().min(7).max(255).required('Password is required'),
      policy: Yup.boolean().oneOf([true], 'This field must be checked')
    }),
    onSubmit: async (values, helpers): Promise<void> => {
        dispatch(
            signupWithInviteRequest({
                firstName: values.firstName,
                lastName: values.lastName,
                password: values.password,
                invitationId,
                token
            })
        );
    }
  });

  return (
    <form noValidate onSubmit={formik.handleSubmit}>
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
          Join Now
        </Button>
      </Box>
    </form>
  );
};

export default JWTRegisterWithInvitation;
