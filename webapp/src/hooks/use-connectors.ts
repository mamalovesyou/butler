import { useSelector } from "react-redux";
import { RootState } from "../features";

export const useConnectors = () =>
  useSelector((state: RootState) => state.connectors);
