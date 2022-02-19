import { useSelector } from "react-redux";
import { RootState } from "../features";

export const useDataSources = () =>
  useSelector((state: RootState) => {
    return state.dataSources
  });
