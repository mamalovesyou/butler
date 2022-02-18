import * as yup from 'yup';
import {JSONSchema4} from "json-schema";
import {labelize} from "../../string";
import isArray from "lodash/isArray";


export const ToYupNumber = (name: string, property: JSONSchema4, required: string[]): yup.NumberSchema => {
    const schema: yup.NumberSchema = yup.number()
    const label: string = property.title || labelize(name);

    if (required.length > 0 && required.includes(name)) {
        schema.required(`${label} is required`)
    }
    if (property.type === "integer") {
        schema.integer(label + " must be an integer.")
    }
    if (property.maximum) {
        schema.max(property.maximum)
    }

    if (property.minimum) {
        schema.min(property.maximum)
    }

    return schema
}