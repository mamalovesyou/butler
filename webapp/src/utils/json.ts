import {JSONSchema4} from "json-schema";
import {labelize} from "./string";
const { buildYup } = require("schema-to-yup");


export const BuildYupFromJSONSchema = (schema: JSONSchema4) => {
    const errMessages = {}
    Object.entries(schema.properties).forEach(([name, property]: [string, JSONSchema4]) => {
        errMessages[name] = errMessages[labelize(name)] || {};
        if (schema.required) {
            errMessages[name].required = `${labelize(name)} is required.`
        }
        if (property.format) {
            errMessages[name].format = `${labelize(name)} dosen't match expected format.`
        }
    })
    console.log(errMessages)
    return buildYup(schema, { errMessages })
}