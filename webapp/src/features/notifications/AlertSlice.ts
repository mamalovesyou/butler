import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import {IAlert} from "./types";

export type AlertStateType = {
    alerts: IAlert[];
};

const initialState: AlertStateType = {
    alerts: []
}

export const AlertSlice = createSlice({
    name: "alert",
    initialState,
    reducers: {
        createAlert: (state: AlertStateType, action: PayloadAction<IAlert>) => {
            state.alerts.push({
                message: action.payload.message,
                type: action.payload.type
            });
        }
    },
});

export const notificationsActions = AlertSlice.actions;

export default AlertSlice;
