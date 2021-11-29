import type { FC } from 'react';
import { Avatar, Box, Button, Card, CardContent, Grid, Typography } from '@mui/material';
import { useDispatch } from 'react-redux';
import { push } from 'redux-first-history';
import { DASHBOARD_ROOT_PATH } from '../../routes';


export const ConnectDataSourceStep: FC = (props) => {
  const { ...other } = props;
  const dispatch = useDispatch();

  const handleSkip = () => dispatch(push(DASHBOARD_ROOT_PATH));

  return (
    <div {...other}>
      <Typography variant="h6">
        Connect a Data Source
      </Typography>
      <Typography
        color="textSecondary"
        sx={{ mt: 2 }}
        variant="body2"
      >
        You can think a workspace as a group. For instance, if you work for multiples companies you will have a different workspace for each of them.
      </Typography>
      <Grid container>
        {[1, 2, 3, 4, 5, 6].map((val, index) =>
          <Grid item key={index} xs={12} md={6} lg={3} sx={{pl: 2, pr: 2}}>
            <Card
              sx={{ mt: 4 }}
            >
              <CardContent>
                <Box
                  sx={{
                    display: 'flex',
                    flexDirection: {
                      xs: 'column',
                      sm: 'row'
                    }
                  }}
                >
                  <Avatar
                    component="a"
                    // src={company.logo}
                    sx={{
                      background: 'transparent',
                      mr: 2,
                      mb: {
                        xs: 2,
                        md: 0
                      }
                    }}
                    variant="rounded"
                  >
                    IC
                  </Avatar>
                  <div>
                    <Typography
                      color="textPrimary"
                      variant="h6"
                    >
                      Integration {val}
                    </Typography>
                    <Typography variant="body2">
                      Integration subtitle
                    </Typography>
                  </div>
                </Box>
              </CardContent>
            </Card>
          </Grid>
        )}

      </Grid>
      <Box sx={{ pt: 2 }}>
        <Button
          variant="contained"
          onClick={handleSkip}
        >
          Skip
        </Button>
      </Box>
    </div>
  );
};