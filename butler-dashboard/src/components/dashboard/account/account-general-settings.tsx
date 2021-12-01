import type { FC } from 'react';
import {
  Box,
  Button,
  Card,
  CardContent,
  Divider,
  Grid,
  Switch,
  TextField,
  Typography
} from '@mui/material';
import LightThemeIcon from './lightTheme.svg';
import DarkThemeIcon from './darkTheme.svg';
import { useSettings } from '../../../hooks/use-settings';
import { useEffect, useState } from 'react';
import { useAuth } from '../../../hooks/use-auth';

const themes = [
  {
    label: 'Light',
    value: 'light',
    icon: LightThemeIcon
  },
  {
    label: 'Dark',
    value: 'dark',
    icon: DarkThemeIcon
  }
];

const getValues = (settings) => ({
  direction: settings.direction,
  responsiveFontSizes: settings.responsiveFontSizes,
  theme: settings.theme
});

export const AccountGeneralSettings: FC = (props) => {
  const { user } = useAuth();

  // Getting settings
  const { settings, saveSettings } = useSettings();
  const [values, setValues] = useState(getValues(settings));

  useEffect(() => {
    setValues(getValues(settings));
  }, [settings]);

  const handleChange = (field, value): void => {
    setValues({
      ...values,
      [field]: value
    });
  };

  const handleSave = (): void => saveSettings(values);

  return (
    <Box sx={{ mt: 4 }} {...props}>
      <Card>
        <CardContent>
          <Grid container spacing={3}>
            <Grid item md={4} xs={12}>
              <Typography variant="h6">Basic details</Typography>
            </Grid>
            <Grid item md={8} xs={12}>
              <Box
                sx={{
                  display: 'flex',
                  alignItems: 'center'
                }}
              >
                <TextField
                  disabled
                  defaultValue={user.firstName}
                  label="Firstname"
                  size="small"
                  sx={{
                    flexGrow: 1,
                    mr: 3
                  }}
                />
                <TextField
                  disabled
                  defaultValue={user.lastName}
                  label="Lastname"
                  size="small"
                  sx={{
                    flexGrow: 1,
                    mr: 3
                  }}
                />
              </Box>
              <Box
                sx={{
                  display: 'flex',
                  mt: 3,
                  alignItems: 'center'
                }}
              >
                <TextField
                  defaultValue={user.email}
                  disabled
                  label="Email Address"
                  required
                  size="small"
                  sx={{
                    flexGrow: 1,
                    mr: 3
                  }}
                />
              </Box>
            </Grid>
          </Grid>
        </CardContent>
      </Card>
      <Card sx={{ mt: 4 }}>
        <CardContent>
          <Grid container spacing={3}>
            <Grid item md={4} xs={12}>
              <Typography variant="h6">Dashboard Settings</Typography>
            </Grid>
            <Grid item md={8} sm={12} xs={12}>
              <Box
                sx={{
                  alignItems: 'center',
                  display: 'flex',
                  justifyContent: 'space-between',
                  mb: 3
                }}
              >
                <Box
                  sx={{
                    alignItems: 'center',
                    display: 'flex',
                    justifyContent: 'space-between',
                    mb: 3
                  }}
                  className="test"
                >
                  {themes.map((theme) => {
                    const { label, icon: Icon, value } = theme;

                    return (
                      <div key={value}>
                        <Box
                          onClick={() => handleChange('theme', value)}
                          sx={{
                            borderColor:
                              values.theme === value
                                ? 'primary.main'
                                : 'divider',
                            borderRadius: 1,
                            borderStyle: 'solid',
                            borderWidth: 2,
                            cursor: 'pointer',
                            flexGrow: 1,
                            fontSize: 0,
                            m: 1,
                            overflow: 'hidden',
                            p: 1,
                            '& svg': {
                              height: 'auto',
                              width: '100%'
                            }
                          }}
                        >
                          <Icon />
                        </Box>
                        <Typography
                          align="center"
                          sx={{ mt: 1 }}
                          variant="subtitle2"
                        >
                          {label}
                        </Typography>
                      </div>
                    );
                  })}
                </Box>
              </Box>
              <Divider />
              <Box
                sx={{
                  alignItems: 'center',
                  display: 'flex',
                  justifyContent: 'space-between',
                  mt: 3
                }}
              >
                <div>
                  <Typography variant="subtitle1">
                    Responsive font sizes
                  </Typography>
                  <Typography
                    color="textSecondary"
                    sx={{ mt: 1 }}
                    variant="body2"
                  >
                    Toggling this will make fonts to be responsive to your
                    screen size.
                  </Typography>
                </div>
                <Switch
                  checked={values.responsiveFontSizes}
                  name="direction"
                  onChange={(event): void =>
                    handleChange('responsiveFontSizes', event.target.checked)
                  }
                />
              </Box>
              <Box
                sx={{
                  alignItems: 'center',
                  display: 'flex',
                  justifyContent: 'left',
                  mt: 4
                }}
              >
                <Button
                  color="primary"
                  onClick={handleSave}
                  variant="contained"
                >
                  Save Settings
                </Button>
              </Box>
            </Grid>
          </Grid>
        </CardContent>
      </Card>
      {/* TODO: Add self account delete */}
      {/*<Card sx={{ mt: 4 }}>*/}
      {/*  <CardContent>*/}
      {/*    <Grid*/}
      {/*      container*/}
      {/*      spacing={3}*/}
      {/*    >*/}
      {/*      <Grid*/}
      {/*        item*/}
      {/*        md={4}*/}
      {/*        xs={12}*/}
      {/*      >*/}
      {/*        <Typography variant="h6">*/}
      {/*          Delete Account*/}
      {/*        </Typography>*/}
      {/*      </Grid>*/}
      {/*      <Grid*/}
      {/*        item*/}
      {/*        md={8}*/}
      {/*        xs={12}*/}
      {/*      >*/}
      {/*        <Typography*/}
      {/*          sx={{ mb: 3 }}*/}
      {/*          variant="subtitle1"*/}
      {/*        >*/}
      {/*          Delete your account and all of your source data. This is irreversible.*/}
      {/*        </Typography>*/}
      {/*        <Button*/}
      {/*          color="error"*/}
      {/*          variant="outlined"*/}
      {/*        >*/}
      {/*          Delete account*/}
      {/*        </Button>*/}
      {/*      </Grid>*/}
      {/*    </Grid>*/}
      {/*  </CardContent>*/}
      {/*</Card>*/}
    </Box>
  );
};
