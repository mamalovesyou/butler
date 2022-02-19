import type { FC } from 'react';
import type { ApexOptions } from 'apexcharts';
import {
  Badge, Box,
  Button,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  Divider,
  Grid,
  Tooltip,
  Typography
} from '@mui/material';
import { useTheme } from '@mui/material/styles';
import { ArrowRight as ArrowRightIcon } from '../../../icons/arrow-right';
import { InformationCircleOutlined as InformationCircleOutlinedIcon } from '../../../icons/information-circle-outlined';
import Chart from "react-apexcharts";

const data = {
  series: [
    {
      color: 'rgba(86, 100, 210, 0.5)',
      data: 3400,
      label: 'Linkedin'
    },
    {
      color: '#7BC67E',
      data: 8000,
      label: 'Google'
    },
    {
      color: '#455a64',
      data: 500,
      label: 'Other'
    }
  ]
};

export const AnalyticsSocialSources: FC = () => {
  const theme = useTheme();

  const chartOptions: ApexOptions = {
    chart: {
      background: 'transparent',
      stacked: false,
      toolbar: {
        show: false
      }
    },
    colors: data.series.map((item) => item.color),
    dataLabels: {
      enabled: false
    },
    fill: {
      opacity: 1
    },
    labels: data.series.map((item) => item.label),
    legend: {
      show: false
    },
    stroke: {
      width: 0
    },
    theme: {
      mode: theme.palette.mode
    }
  };

  const chartSeries = data.series.map((item) => item.data);

  return (
    <Card>
      <CardHeader sx={{ p: 2 }} title="Revenue per channel" />
      <Divider />
      <CardContent>
        <Grid container>
          {data.series.map((item) => (
            <Grid
              item
              key={item.label}
              sx={{
                alignItems: 'center',
                display: 'flex',
                p: 1
              }}
              xs={12}
            >
              <Grid item sx={{
                alignItems: 'center',
                display: 'flex',
                p: 1
              }} xs={8}>
                <Box
                  sx={{
                    border: 3,
                    borderColor: item.color,
                    borderRadius: '50%',
                    height: 16,
                    mr: 1,
                    width: 16
                  }}
                />
                <Typography variant="subtitle2">
                  {item.label}:
                </Typography>
              </Grid>
              <Grid item xs={4}>
                <Typography sx={{ px: 2, color: 'back' }} variant="subtitle2">
                  $ {item.data}
                </Typography>
              </Grid>
            </Grid>
          ))}
        </Grid>
      </CardContent>
    </Card>
  );
};
