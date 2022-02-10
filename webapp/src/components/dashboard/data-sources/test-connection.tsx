import {Alert, AlertTitle, Box, Typography} from "@mui/material";
import {Api, V1Connector} from "../../../api";
import {LoadingButton} from "@mui/lab";
import {useState} from "react";
import * as React from "react";
import {ConnectorWithSource} from "./connectors-table-list";

interface TestConnectionProps {
    connector: ConnectorWithSource;
}

export const TestConnection = (props: TestConnectionProps) => {

    const [loading, setLoading] = useState(false);
    const [message, setMessage] = useState("");
    const [status, setStatus] = useState("");

    const testConnection = async () => {
        setLoading(true);
        const {data} = await Api.v1.connectorsServiceTestConnection({
            connectorId: props.connector.id
        })
        setLoading(false);
        setStatus(data.status);
        setMessage(data.message);
    }

    return (
        <Box>
            <Typography variant="h4">Test Connection</Typography>
            {(!loading && status !== "") ?
                <Box sx={{width: '100%', p: 4}}>
                    <Alert severity={status === "failed" ? "error" : "success"}>
                        <AlertTitle>{status.toUpperCase()}</AlertTitle>
                        {message}
                    </Alert>
                </Box> : null}
            <LoadingButton loading={loading} onClick={testConnection}>
                Try it !
            </LoadingButton>
        </Box>
    );
};

export default TestConnection;