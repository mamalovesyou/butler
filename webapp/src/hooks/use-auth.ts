import { useSelector } from "react-redux";
import { RootState } from "../features";

export const useAuth = () => useSelector((state: RootState) => state.auth);
