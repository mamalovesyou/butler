import * as yup from 'yup';
import {JSONSchema4} from "json-schema";
import {labelize} from "../../string";
import {FORMAT_VALIDATOR_MAP, ValidateFormatOrThrowError} from "../types";


export const ToYupString = (name: string, property: JSONSchema4, required: string[]): yup.StringSchema => {
    const schema: yup.StringSchema<string, any> = yup.string()
    const label: string = property.title || labelize(name);

    console.log("string property", name, property)

    if (required.length > 0 && required.includes(name)) {
        schema.required(`${label} is required`)
    }

    if (property.pattern) {
        const re = new RegExp(property.pattern);
        console.log("property format", property.pattern, re, "2021-21-0".match(re), "2021-21-01".match(re))
        schema.matches(re, `${label} does not match the required format`)
    }

    if (property.format) {
        ValidateFormatOrThrowError(property.format)
        const { regex , message } = FORMAT_VALIDATOR_MAP[property.format]
        schema.matches(regex, { message })
    }

    return schema
}