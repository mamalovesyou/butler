import React, {Fragment, ReactNode} from "react";
import {Box, CircularProgress, Container, Divider, Typography} from "@mui/material";

interface IDashboardPage {
    title: string;
    loading?: boolean;
    children?: ReactNode;
}

export const DashboardPage = (props: IDashboardPage = {title: "", loading: true, children: null}) => {

    const {title, loading, children} = props;

    console.log(children)

    return <Box
        component="main"
        sx={{
            flexGrow: 1,
            py: 2,
            height: '100%'
        }}
    >
        <Container maxWidth="md" sx={{ height: '100%'}}>
            <Typography variant="h4">{title}</Typography>
            <Divider sx={{mb: 3}}/>
            { loading || !children ?
                <Box
                    sx={{
                        display: 'flex',
                        alignItems: "center",
                        justifyContent: "center",
                        width: '100%',
                        py: 8
                    }}
                >
                    <CircularProgress/>
                </Box>
                : <Fragment>{children}</Fragment>
            }
        </Container>
    </Box>

}

export default DashboardPage