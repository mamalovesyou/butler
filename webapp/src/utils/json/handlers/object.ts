import * as yup from 'yup';
import {JSONSchema4} from "json-schema";
import {ValidateTypeOrThrowError} from "../types";
import {AllowedSchema, GetSchemaBuilder} from "../builder";
import isArray from "lodash/isArray";


export const ToYupObject = (name: string, property: JSONSchema4, _required: string[]): yup.ObjectSchema<any> => {
    const schema: yup.ObjectSchema<any> = yup.object();
    const properties = Object.entries(property.properties);
    const shape: Record<string, AllowedSchema> = {}
    const propertyRequired: string[] = isArray(property.required) ? property.required : [];
    properties.map((entry) => {
        const [name, p]: [string, JSONSchema4] = entry;
        ValidateTypeOrThrowError(p.type)
        const builder = GetSchemaBuilder(p.type);
        shape[name] = builder(name, p, propertyRequired);
    })

    return schema.shape(shape);
}