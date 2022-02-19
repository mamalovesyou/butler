import * as yup from 'yup';
import {JSONSchema4} from "json-schema";
import {labelize} from "../../string";
import isArray from "lodash/isArray";


export const ToYupArray = (name: string, property: JSONSchema4, required: string[]): yup.ArraySchema<any> => {
    const schema: yup.ArraySchema<any> = yup.array()
    const label: string = property.title || labelize(name);

    if (required.length > 0 && required.includes(name)) {
        schema.required(`${label} is required`)
    }

    return schema
}