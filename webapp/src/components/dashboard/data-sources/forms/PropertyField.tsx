import {TextField} from "@mui/material";
import * as React from "react";
import { FormikProps } from 'formik';
import ChipInput from "../../../common/MultiInput";

type AllowedDataTypes = string | number | string[] | number[];
export interface JSONFieldProperty {
    title?: string;
    default?: AllowedDataTypes;
    description?: string;
    format?: string;
    type: string
    items?: JSONFieldProperty;
}

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

interface PropertyFieldProps {
    form: FormikProps<object>;
    label: string;
    name: string;
    value: any;
    jsonProps: JSONFieldProperty;
    error: boolean;
    onBlur?: React.FocusEventHandler<HTMLInputElement | HTMLTextAreaElement>;
    onChange?: React.ChangeEventHandler<HTMLTextAreaElement | HTMLInputElement>;
}

export const PropertyField = (props: PropertyFieldProps) => {
    const { form, jsonProps, name, label, value, error, onBlur} = props;
    const {description, type, format} = props.jsonProps;

    const handleChange = (value) => {
        form.setFieldValue(name, value);
    }

    const getComponent = (property: JSONFieldProperty): React.ReactElement => {
        switch (property.type) {
            case "array":
                return <ChipInput
                    error={error}
                                   values={value}
                                   name={name}
                                   label={label}
                                   helperText={description}
                                   type={getInputType(property.items?.type)}
                                   onChange={handleChange}
                                   onBlur={onBlur}
                />
            case "number":
            case "integer":
            case "string":
            default:
                return <TextField
                    error={error}
                    type={getInputType(type)}
                    fullWidth
                    helperText={jsonProps.description}
                    label={label}
                    margin="normal"
                    name={name}
                    onBlur={onBlur}
                    onChange={(e) => handleChange(e.target.value)}
                    value={value}
                />
        }
    }

    return getComponent(jsonProps)
}