import { useState, FC } from 'react';
import type { ChangeEvent } from 'react';
import { Box, Container, Divider, Tab, Tabs, Typography } from '@mui/material';
import { AccountGeneralSettings } from '../../components/dashboard/account/account-general-settings';
import { AccountTeamSettings } from '../../components/dashboard/account/account-team-settings';


enum TabsEnum {
    GENERAL = 'GENERAL',
    TEAM = 'TEAM'
}
const tabs = [
  { label: 'General', value: TabsEnum.GENERAL },
  { label: 'Team', value:  TabsEnum.TEAM },
];

const Account: FC = () => {
  const [currentTab, setCurrentTab] = useState<string>(TabsEnum.GENERAL);

  const handleTabsChange = (event: ChangeEvent<{}>, value: string): void => {
    setCurrentTab(value);
  };

  return (
    <>
      <Box
        component="main"
        sx={{
          flexGrow: 1,
          py: 8
        }}
      >
        <Container maxWidth="md">
          <Typography variant="h4">
            Account
          </Typography>
          <Tabs
            indicatorColor="primary"
            onChange={handleTabsChange}
            scrollButtons="auto"
            textColor="primary"
            value={currentTab}
            variant="scrollable"
            sx={{ mt: 3 }}
          >
            {tabs.map((tab) => (
              <Tab
                key={tab.value}
                label={tab.label}
                value={tab.value}
              />
            ))}
          </Tabs>
          <Divider sx={{ mb: 3 }} />
          {currentTab === TabsEnum.GENERAL && <AccountGeneralSettings />}
          {currentTab === TabsEnum.TEAM && <AccountTeamSettings />}
        </Container>
      </Box>
    </>
  );
};

export default Account;
