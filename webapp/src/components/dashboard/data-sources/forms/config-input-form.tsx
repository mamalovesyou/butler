import * as React from "react";

const {buildYup} = require("schema-to-yup");
import {useFormik} from 'formik';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import {
    FormHelperText,
    TextField,
} from "@mui/material";
import {labelize} from "../../../../utils/string";
import {Api, V1Connector, V1DataSource} from "../../../../api";

interface JSONFieldProperty {
    description?: string;
    format?: string;
    type: string
}

interface PropertyFieldProps {
    label: string;
    name: string;
    value: string | number;
    jsonProps: JSONFieldProperty;
    error: boolean;
    onBlur?: React.FocusEventHandler<HTMLInputElement | HTMLTextAreaElement>;
    onChange?: React.ChangeEventHandler<HTMLTextAreaElement | HTMLInputElement>;
}

export const PropertyField = (props: PropertyFieldProps) => {
    const {name, label, value, error, onChange, onBlur} = props;
    const {description, type, format} = props.jsonProps;

    const getInputType = (type: string): string => {
        switch (type) {
            case "number":
            case "integer":
                return "number"
            case "string":
            default:
                return "text"
        }
    }

    return <TextField
        error={error}
        type={getInputType(type)}
        fullWidth
        helperText={description}
        label={label}
        margin="normal"
        name={name}
        onBlur={onBlur}
        onChange={onChange}
        value={value}
    />
}

interface ConfigInputFormProps {
    connectorId: string;
    source: V1DataSource;
    onComplete: (connector: V1Connector) => void;
}

export const ConfigInputForm = (props: ConfigInputFormProps) => {
    const { source, connectorId, onComplete } = props;
    const schemaObj = JSON.parse(source.configurationInputJSONSchema.trim());
    const {required, properties} = schemaObj;

    const getYupSchema = (obj: object, fields: string[]): object => {
        let config = {}
        fields.forEach((name: string) => config[name] = {
            required: "This field is required.",
            format: "Wrong format."
        });
        return buildYup(obj, config)
    }

    const getInitialValues = (properties: Record<string, JSONFieldProperty>) => {
        let initialValues = {}
        Object.entries(properties).map(([name, property]) => {
            initialValues[name] = ''
        })
        return initialValues;
    }

    const formik = useFormik({
        initialValues: getInitialValues(properties),
        validationSchema: getYupSchema(schemaObj, required),
        onSubmit: async (values): Promise<void> => {
            const response = await Api.v1.connectorsServiceMutateConnector({
                    connectorId,
                    config: {
                        ...values
                    }
                });
            onComplete(response.data)
            // TODO: Handle error
        }
    });


    return (
        <form noValidate onSubmit={formik.handleSubmit}>
            {formik.errors.submit && (
                <Box sx={{ py: 2 }}>
                    <FormHelperText error>{formik.errors.submit}</FormHelperText>
                </Box>
            )}
            <Box sx={{p: 0}}>
                {Object.entries(properties).map(([name, jsonProps]: [string, JSONFieldProperty]) =>
                    <PropertyField
                        key={name}
                        error={Boolean(formik.touched[name] && formik.errors[name])}
                        name={name} onBlur={formik.handleBlur}
                        onChange={formik.handleChange}
                        value={formik.values[name]}
                        jsonProps={jsonProps}
                        label={labelize(name)}
                    />
                )}
            </Box>
            <Box sx={{display: 'flex', pt: 2}}>
                <Button
                    type="submit"
                    variant="contained"
                    disabled={formik.isSubmitting}
                >
                    Create
                </Button>
            </Box>
        </form>
    );
};