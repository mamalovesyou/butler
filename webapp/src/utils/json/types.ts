import isArray from 'lodash/isArray';

export const SUPPORTED_TYPE = ["string", "number", "integer", "array", "object"]

export const ValidateTypeOrThrowError = (type: string | string[]) => {
    if (isArray(type)) {
        throw Error(`Unsupported type array. Supported types are: ${SUPPORTED_TYPE.toString()}`)
    }
    if (!SUPPORTED_TYPE.includes(type)) {
        throw Error(`Type: ${type} is not supported. Supported types are: ${SUPPORTED_TYPE.toString()}`)
    }
}

export const SUPPORTED_FORMAT = ["date"]
export const FORMAT_VALIDATOR_MAP = {
    date: {
        regex: /^\d{4}-\d{2}-\d{2}$/,
        message: "This doesn't match the right format. The date should be formatted as follow: yyyy-mm-dd. Ex: 2020-01-01"
    }
}

export const ValidateFormatOrThrowError = (format: string) => {
    if (!SUPPORTED_FORMAT.includes(format)) {
        throw Error(`Format: ${format} is not supported. Supported format are: ${SUPPORTED_FORMAT.toString()}`)
    }
}