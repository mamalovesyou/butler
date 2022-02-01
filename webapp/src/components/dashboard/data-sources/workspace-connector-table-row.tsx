import * as React from 'react';
import Typography from '@mui/material/Typography';
import {Avatar, Box, Button, TableCell, TableRow} from "@mui/material";
import OAuthPopup from "../../oauth-popup";
import {useMemo} from "react";
import {UserCircle as UserCircleIcon} from "../../../icons/user-circle";
import {V1CatalogConnector} from "../../../api";

type ICatalogTableRowProps = {
    connector: V1CatalogConnector;
    onOAuthConnect: (params: { code: string, provider: string }) => void;
};

export const CatalogTableRow: React.FC<ICatalogTableRowProps> = ({ connector, onOAuthConnect }) => {

    const imageURL = useMemo(() => {
        const blob = new Blob([connector.iconSvg], {type: 'image/svg+xml'});
        return URL.createObjectURL(blob);
    }, [connector.iconSvg]);

    return (
        <TableRow>
            <TableCell>
                <Box
                    sx={{
                        alignItems: 'center',
                        display: 'flex'
                    }}
                >
                    <Avatar
                        sx={{
                            bgcolor: 'transparent',
                            height: 50,
                            width: 50,
                            mr: 1
                        }}
                        alt={`${connector.name} logo`}
                        src={imageURL}
                    />
                    <div>
                        <Typography
                            variant="subtitle2">{connector.name}</Typography>
                    </div>
                </Box>
            </TableCell>
            <TableCell align="right">
                <OAuthPopup
                    url={connector.authUrl}
                    title={`Connect ${connector.name}`}
                    onCode={(code, params) => onOAuthConnect({
                        code,
                        provider: connector.name
                    })
                    }
                >
                    <Button variant="contained">
                        Connect
                    </Button>
                </OAuthPopup>
            </TableCell>
        </TableRow>
    );
}



export default CatalogTableRow;