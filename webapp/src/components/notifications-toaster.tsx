import React, {useEffect, useState} from "react";
import {IAlert} from "../features/notifications";
import {useAlerts} from "../hooks/use-notifications";
import Snackbar from '@mui/material/Snackbar';
import MuiAlert, {AlertProps} from '@mui/material/Alert';


const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref,
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});


const NotificationToaster = () => {

    const alerts = useAlerts();
    const [alert, setAlert] = useState<IAlert>(null);
    const [open, setOpen] = useState(false);

    useEffect(() => {
        console.log("Got alerts", alerts)
        if (alerts.length > 0) {
            setAlert(alerts[alerts.length - 1]);
            setOpen(true);
        }
    }, [alerts]);

    const handleClose = () => {
        setOpen(false);
    };

    return <Snackbar open={open} autoHideDuration={5000} anchorOrigin={{
        vertical: 'top',
        horizontal: 'right'
    }} onClose={handleClose}>
        { alert ? <Alert onClose={handleClose} severity={alert.type} sx={{width: '100%'}}>
            {alert.message}
        </Alert> : null }
    </Snackbar>

};

export default NotificationToaster;
