import {useState, FC, useEffect} from 'react';
import {Box, Container, Divider, Tab, Tabs, Typography} from '@mui/material';
import {Link, Outlet, useLocation} from "react-router-dom";
import {
    ORGANIZATION_WORKSPACES_PATH,
    ORGANIZATION_TEAM_PATH, ORGANIZATION_ROOT_PATH,
} from "../../../routes/constants";



const tabs = [
    {label: 'Workspaces', path: ORGANIZATION_WORKSPACES_PATH },
    {label: 'Team',  path: ORGANIZATION_TEAM_PATH },
];

const Organization: FC = () => {

    const [value, setValue] = useState(ORGANIZATION_WORKSPACES_PATH)
    const location = useLocation();

    useEffect(() => {
        let path = location.pathname.replace(ORGANIZATION_ROOT_PATH, "");
        while(path.charAt(0) === "/") path = path.substring(1);
        if (path === "") path = ORGANIZATION_WORKSPACES_PATH;
        setValue(path);
    }, [location.pathname])

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
                    <Typography variant="h4">Account</Typography>
                    <Tabs
                        indicatorColor="primary"
                        scrollButtons="auto"
                        textColor="primary"
                        value={value}
                        variant="scrollable"
                        sx={{mt: 3}}
                    >
                        {tabs.map((tab) => (
                            <Tab key={tab.path} label={tab.label} component={Link} to={tab.path} value={tab.path}/>
                        ))}
                    </Tabs>
                    <Divider sx={{mb: 3}}/>
                    <Outlet/>
                </Container>
            </Box>
        </>
    );
};

export default Organization;
