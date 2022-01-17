import {Cancel, Tag} from "@mui/icons-material";
import {Box, Chip, InputAdornment, TextField} from "@mui/material";
import {FC, useRef, useState} from "react";
import {Mail as MailIcon} from "../../icons/mail";
import {isValidEmail} from "../../utils/email";
import {V1CatalogConnector, V1WorkspaceConnector} from "../../api";


export const MultiEmailInput: FC = () => {
    const [emails, setEmails] = useState([]);
    const [error, setError] = useState('');

    const handleDelete = (value) => {
        console.log(value);
        const newtags = emails.filter((val) => val !== value);
        setEmails(newtags);
    };

    const onValueChange = (e) => {
        e.preventDefault();
        const value = e.target.value;
        const lastkey = value.charAt(value.length - 1);
        console.log("update: ", value)
        if (lastkey === "," || lastkey === " ") {
            const email = value.slice(0, -1);
            console.log("isValidEmail", email, isValidEmail(email))
            if (isValidEmail(email)) {
                setEmails([...emails, email]);
                e.target.value = "";
            } else {
                const errMsg = `Invalid email: ${email}`;
                setError(errMsg);
                e.target.value = email;
            }
        } else {
            setError('');
        }
    };

    return <TextField
        label="Email address"
        placeholder="Add multiple addresses separated by commas or space"
        onChange={onValueChange}
        error={Boolean(error.length)}
        helperText={error}
        multiline
        sx={{
            m: 1.5,
            flexGrow: 1,
            maxWidth: '100%'
        }}
        margin='none'
        InputProps={{
            startAdornment: (
                <Box sx={{margin: "0 0.2rem 0 0", display: "flex", maxWidth: 120 }}>
                    {emails.map((data, index) => <Chip key={index} label={data} onDelete={() => handleDelete(data)}/>
                    )}
                </Box>
            )
        }}/>
}

export default MultiEmailInput;
