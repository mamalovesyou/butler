import * as yup from 'yup';
import {JSONSchema4} from "json-schema";
import {ValidateTypeOrThrowError} from "./types";
import {ToYupArray, ToYupString, ToYupObject, ToYupNumber} from "./handlers";

export type AllowedSchema = yup.StringSchema | yup.NumberSchema | yup.ArraySchema<any> | yup.ObjectSchema<any>;

interface SchemaBuilder {
    (name: string, property: JSONSchema4, required: string[]): AllowedSchema;
}

export const GetSchemaBuilder = (type: string | string[]): SchemaBuilder => {
    ValidateTypeOrThrowError(type)
    switch(type) {
        case "object":
            return ToYupObject
        case "string":
            return ToYupString
        case "array":
            return ToYupArray

        case "number":
        case "integer":
            return ToYupNumber
        default:
            throw Error("Invalid type: " + type)
    }
}

export const BuildYup = (schema: JSONSchema4): yup.AnyObjectSchema => {
    if (schema.type !== "object") {
        throw Error("JSON Schema root type must be object")
    }
    return ToYupObject("", schema, [])
}