import * as React from "react";
import convertToYup from "json-schema-yup-transformer";
import {useFormik, getIn} from 'formik';
import Box from '@mui/material/Box';
import {labelize} from "../../../../utils/string";
import {useEffect, useState} from "react";
import {JSONFieldProperty, PropertyField} from "./PropertyField";
import {Alert, AlertTitle, Button, CircularProgress} from "@mui/material";
import isEqual from "lodash/isEqual";
import {Api} from "../../../../api";


interface ConfigInputFormProps {
    connectorId: string;
    initialValues: object;
    configJSONSchema: string;
    secretsJSONSchema: string;
}

export const ConfigInputForm = (props: ConfigInputFormProps) => {

    const getInitialValues = (properties: Record<string, JSONFieldProperty>) => {
        let allInitialValues = {...initialValues}
        Object.entries(properties).map(([name, property]) => {
            if (!(name in allInitialValues)) {
                if (property.default) {
                    allInitialValues[name] = property.default
                } else {
                    allInitialValues[name] = property.type === 'array' ? [] : ''
                }
            }
        })
        return allInitialValues;
    }

    const {connectorId, configJSONSchema, secretsJSONSchema, initialValues} = props;

    const configJSONSchemaObj = JSON.parse(configJSONSchema.trim());
    const configYupSchema = convertToYup(configJSONSchemaObj);
    const configInitialValues =  getInitialValues(configJSONSchemaObj.properties);
    const secretsJSONSchemaObj = JSON.parse(secretsJSONSchema.trim());
    const secretsYupSchema = convertToYup(secretsJSONSchemaObj);
    const secretsInitialValues =  getInitialValues(secretsJSONSchemaObj.properties)




    const [enableUpdate, setEnableUpdate] = useState(false);
    const [updating, setUpdating] = useState(false);
    const [error, setError] = useState('');

    const configForm = useFormik({
        initialValues: configInitialValues,
        validationSchema: configYupSchema,
        onSubmit: () => {}
    });
    const secretsForm = useFormik({
        initialValues: secretsInitialValues,
        validationSchema: secretsYupSchema,
        onSubmit: () => {}
    });

    useEffect(() => {
        const isValid = configYupSchema.isValidSync(configForm.values);
        const formValuesUpdates = !isEqual(configInitialValues, configForm.values) || !isEqual(secretsInitialValues, secretsForm.values);
        console.log("is valid: ", isValid, "config changed: ", !isEqual(configInitialValues, configForm.values), "secret changed: ", !isEqual(secretsInitialValues, secretsForm.values), formValuesUpdates)
        setEnableUpdate(isValid && formValuesUpdates);
    }, [configForm.values, secretsForm.values])

    const updateConfig = async () => {
        setError('');
        setUpdating(true);
        try {
            const response = await Api.v1.connectorsServiceMutateConnector({
                connectorId,
                config: configForm.values,
                secrets: isEqual(secretsInitialValues, secretsForm.values) ? {} : secretsForm.values
            });
            const {status, message} = response.data;
            if (status !== "succeeded") {
                setError(`Failed: ${message}`)
            }
        } catch (err) {
            console.log("Unable to update config", err);
            setError("Unable to update config: ");
        }
        setUpdating(false);
    }


    return <>
        <Box sx={{p: 2}}>
            <form noValidate onSubmit={secretsForm.handleSubmit}>
                {Object.entries(secretsJSONSchemaObj.properties).map(([name, jsonProps]: [string, JSONFieldProperty]) => {
                        // const fieldName = `config.${name}`
                        return <PropertyField
                            isSecret
                            key={name}
                            form={secretsForm}
                            error={Boolean(getIn(secretsForm.touched, name) && getIn(secretsForm.errors, name))}
                            name={name} onBlur={secretsForm.handleBlur}
                            onChange={secretsForm.handleChange}
                            value={getIn(secretsForm.values, name)}
                            jsonProps={jsonProps}
                            label={labelize(name)}
                        />
                    }
                )}
            </form>
            <form noValidate onSubmit={configForm.handleSubmit}>
                {Object.entries(configJSONSchemaObj.properties).map(([name, jsonProps]: [string, JSONFieldProperty]) => {
                        // const fieldName = `config.${name}`
                        return <PropertyField
                            key={name}
                            form={configForm}
                            error={Boolean(getIn(configForm.touched, name) && getIn(configForm.errors, name))}
                            name={name} onBlur={configForm.handleBlur}
                            onChange={configForm.handleChange}
                            value={getIn(configForm.values, name)}
                            jsonProps={jsonProps}
                            label={labelize(name)}
                        />
                    }
                )}
            </form>
        </Box>
        {error !== "" && (
            <Box sx={{py: 1}}>
                <Alert severity="error">
                    <AlertTitle>Error</AlertTitle>{error}
                </Alert>
            </Box>
        )}
        <Box sx={{p: 1, width: '100%', display: 'flex', alignItems: "center", justifyContent: "center"}}>
            {updating ? <CircularProgress/>
                : <Button disabled={!enableUpdate} variant="contained" color="primary"
                          onClick={updateConfig}>Update</Button>}
        </Box>
    </>
};