import React from "react";
import {Box, Typography} from "@mui/material";

export const ComingSoon = () => {


    return <Box
        sx={{
            display: 'flex',
            alignItems: "center",
            justifyContent: "center",
            width: '100%',
            py: 8
        }}
    >
        <Typography variant="h4">Coming Soon !</Typography>
    </Box>
}

export default ComingSoon;