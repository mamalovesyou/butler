import { useState } from 'react';
import type { FC } from 'react';
import numeral from 'numeral';
import {
  Box,
  Button,
  Card,
  CardActions,
  CardHeader,
  Divider,
  Table,
  TableBody,
  TableCell,
  TableFooter,
  TableHead,
  TableRow,
  TableSortLabel,
  Tooltip,
  Typography
} from '@mui/material';
import type { CardProps } from '@mui/material';
import { InformationCircleOutlined as InformationCircleOutlinedIcon } from '../../../icons/information-circle-outlined';
import { ArrowRight as ArrowRightIcon } from '../../../icons/arrow-right';

interface CPCStat {
  channel: string;
  spent: number;
  leads: number;
  mqls: number;
  opportunities: number;
  won: number;
  revenue: number;
  cac: number;
}

const fakeCPCData: CPCStat[] = [
  {
    channel: "Linkedin / CPC",
    spent: 6500,
    leads: 30,
    mqls: 20,
    opportunities: 1,
    revenue: 3500,
    won: 3,
    cac: 2000
  },
  {
    channel: "Google / CPC",
    spent: 6500,
    leads: 30,
    mqls: 20,
    opportunities: 1,
    revenue: 6500,
    won: 2,
    cac: 2900
  },
  {
    channel: "Facebook / CPC",
    spent: 6500,
    leads: 30,
    mqls: 20,
    opportunities: 1,
    revenue: 8500,
    won: 1,
    cac: 3000
  },
  {
    channel: "Unknown",
    spent: 0,
    leads: 2,
    mqls: 12,
    opportunities: 1,
    revenue: 1000,
    won: 1,
    cac: 0
  }
]
export const AnalyticsOverviewTable: FC<CardProps> = (props) => {



  return (
    <Card {...props}>
      <CardHeader sx={{ p: 2 }} title="CPC per channel" />
      <Divider />
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>
              Channel
            </TableCell>
            <TableCell>
              $ Spent
            </TableCell>
            <TableCell>
              Leads
            </TableCell>
            <TableCell>
              MQLS
            </TableCell>
            <TableCell>
              Opportunities
            </TableCell>
            <TableCell>
              Won
            </TableCell>
            <TableCell>
              Revenue
            </TableCell>
            <TableCell>
              CAC
            </TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {fakeCPCData.map((item: CPCStat, index: number) => (
            <TableRow
              key={index}
              sx={{
                '&:last-child td': {
                  border: 0
                }
              }}
            >
              <TableCell>
                {item.channel}
              </TableCell>
              <TableCell>
                $ {item.spent}
              </TableCell>
              <TableCell>
                {item.leads}
              </TableCell>
              <TableCell>
                {item.mqls}
              </TableCell>
              <TableCell>
                {item.opportunities}
              </TableCell>
              <TableCell>
                {item.won}
              </TableCell>
              <TableCell>
                $ {item.revenue}
              </TableCell>
              <TableCell>
                {item.cac}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
        <TableHead>
          <TableRow>
            <TableCell>
              Total
            </TableCell>
            <TableCell>
              $ {fakeCPCData.reduce((accumulator: number, current: CPCStat) => accumulator + current.spent, 0)}
            </TableCell>
            <TableCell>
              {fakeCPCData.reduce((accumulator: number, current: CPCStat) => accumulator + current.leads, 0)}
            </TableCell>
            <TableCell>
              {fakeCPCData.reduce((accumulator: number, current: CPCStat) => accumulator + current.mqls, 0)}
            </TableCell>
            <TableCell>
              {fakeCPCData.reduce((accumulator: number, current: CPCStat) => accumulator + current.opportunities, 0)}
            </TableCell>
            <TableCell>
              {fakeCPCData.reduce((accumulator: number, current: CPCStat) => accumulator + current.won, 0)}
            </TableCell>
            <TableCell>
              $ {fakeCPCData.reduce((accumulator: number, current: CPCStat) => accumulator + current.revenue, 0)}
            </TableCell>
            <TableCell>
              {fakeCPCData.reduce((accumulator: number, current: CPCStat) => accumulator + current.cac, 0)}
            </TableCell>
          </TableRow>
        </TableHead>
      </Table>
    </Card>
  );
};
