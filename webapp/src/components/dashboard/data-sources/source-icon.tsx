import {Avatar} from "@mui/material";

interface SourceIconProps {
    name: string;
    xml: string;
}

export const SourceIcon = (props: SourceIconProps) => {
    const src = `data:image/svg+xml;utf8,${encodeURI(props.xml)}`;
    return <Avatar alt={props.name} src={src} />
}

export default SourceIcon;