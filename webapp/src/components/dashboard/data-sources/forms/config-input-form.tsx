import * as React from "react";
import convertToYup from "json-schema-yup-transformer";
import {useFormik} from 'formik';
import Box from '@mui/material/Box';
import {labelize} from "../../../../utils/string";
import {useEffect, useState} from "react";
import {JSONFieldProperty, PropertyField} from "./PropertyField";
import {Api} from "../../../../api";
import {Alert, AlertTitle, Button, CircularProgress} from "@mui/material";
import isEqual from "lodash/isEqual";


interface ConfigInputFormProps {
    connectorId: string;
    initialValues: object;
    inputJSONSchema: string;
}

export const ConfigInputForm = (props: ConfigInputFormProps) => {

    const getInitialFormValues = (properties: Record<string, JSONFieldProperty>) => {
        let allInitialValues = {...initialValues}
        Object.entries(properties).map(([name, property]) => {
            if (!(name in allInitialValues)) {
                allInitialValues[name] = property.type === 'array' ? [] : ''
            }
        })
        return allInitialValues;
    }

    const {connectorId, inputJSONSchema, initialValues} = props;
    const schemaObj = JSON.parse(inputJSONSchema.trim());
    const yupSchema = convertToYup(schemaObj);
    const initialFormValues = getInitialFormValues(schemaObj.properties);
    const [enableUpdate, setEnableUpdate] = useState(false);
    const [error, setError] = useState('');


    const formik = useFormik({
        initialValues: initialFormValues,
        validationSchema: yupSchema,
        onSubmit: async (values) => {
            setError('');
            const { status, message } = await testConfig(values);
            if (status === "succeeded") {
                setError("");
                await updateConfig(values);
            } else {
                setError(`Failed: ${message}`)
            }
        },
    });

    useEffect(() => {
        const isValid = yupSchema.isValidSync(formik.values);
        setEnableUpdate(isValid && !isEqual(initialFormValues, formik.values));
    }, [formik.values])

    const updateConfig = async (values: Record<string, unknown>) => {
        try {
            await Api.v1.connectorsServiceMutateConnector({
                connectorId,
                config: values,
            })
        } catch (err) {
            console.log("Unable to update config", err);
            setError("Unable to update config");
        }
    }

    const testConfig = async (values: Record<string, unknown>) => {
        try {
            const response = await Api.v1.connectorsServiceTestConnection({
                connectorId,
                config: values
            });
            return response.data;
        } catch (err) {
            console.log("Error", err);
            setError(`Failed: ${err}`)
        }
    }


    return <form noValidate>
        <Box sx={{p: 2}}>
            {Object.entries(schemaObj.properties).map(([name, jsonProps]: [string, JSONFieldProperty]) =>
                <PropertyField
                    key={name}
                    form={formik}
                    error={Boolean(formik.touched[name] && formik.errors[name])}
                    name={name} onBlur={formik.handleBlur}
                    onChange={formik.handleChange}
                    value={formik.values[name]}
                    jsonProps={jsonProps}
                    label={labelize(name)}
                />
            )}
        </Box>
        {error !== "" && (
            <Box sx={{ py: 1 }}>
                <Alert severity="error">
                    <AlertTitle>Error</AlertTitle>{error}
                </Alert>
            </Box>
        )}
        <Box sx={{p: 1, width: '100%', display: 'flex', alignItems: "center", justifyContent: "center"}}>
            {formik.isSubmitting ? <CircularProgress/>
                : <Button disabled={ !enableUpdate || !formik.dirty} variant="contained" color="primary" type="submit">Update</Button>}
        </Box>
    </form>
};