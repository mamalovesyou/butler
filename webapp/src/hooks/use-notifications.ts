import { useSelector } from "react-redux";
import { RootState } from "../features";
import { IAlert } from "../features/notifications";

export const useAlerts = () => useSelector((state: RootState): IAlert[] => state.notifications.alerts);
