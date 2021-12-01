import { useSelector } from "react-redux";
import { RootState } from "../features";

export const useCatalog = () =>
  useSelector((state: RootState) => {
    return state.connectors.catalog
  });

export const useConnectors = () =>
  useSelector((state: RootState) => state.connectors);

export const useDataSources = () =>
  useSelector((state: RootState) => state.connectors.connectors);

