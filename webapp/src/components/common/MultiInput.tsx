import {Autocomplete, Box, Chip, TextField} from "@mui/material";
import {useEffect, useState} from "react";
import * as React from "react";
import isEqual from "lodash/isEqual";

interface MultiInputProps {
    name: string;
    label: string;
    type: string;
    error: boolean;
    onBlur?: React.FocusEventHandler<HTMLInputElement | HTMLTextAreaElement>;
    onChange?: (data: Array<any>) => void;
    helperText?: string;
    placeholder?: string;
    values: any[];
}

export const ChipInput = (props: MultiInputProps) => {
    const {label, type, onChange, helperText, placeholder, error, name, values} = props;
    const [receivers, setReceivers] = useState<string[]>(values.map(String));

    useEffect(() => {
        if (!isEqual(values, receivers) && onChange) {
            if (type === "number"){ onChange(receivers.map(Number)) }
            else { onChange(receivers) }
        }
    }, [receivers])


    return <Autocomplete
        multiple
        id="tags-filled"
        options={[]}
        defaultValue={[]}
        value={receivers}
        freeSolo
        onChange={(e, value) => setReceivers((state) => value)}
        renderTags={(
            value: any[],
            getTagProps: (arg0: { index: any }) => JSX.IntrinsicAttributes
        ) =>
            receivers.map((option: any, index: any) => {
                return (
                    <Chip
                        key={index}
                        color="primary"
                        label={option}
                        {...getTagProps({ index })}
                    />
                );
            })
        }
        renderInput={(params: any) => (
            <TextField
                {...params}
                name={name}
                label={label}
                error={error}
                helperText={helperText}
                placeholder={placeholder}
            />
        )}
    />
}


export default ChipInput;
