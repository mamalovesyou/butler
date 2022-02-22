import { Api as GenAPI, ApiConfig } from "./gen";
import {API_BASE_URL} from "../config";

export class ButlerApi<T> extends GenAPI<T> {
  private static apiInstance;

  constructor(apiConfig?: ApiConfig) {
    super(apiConfig);
    if (ButlerApi.apiInstance) {
      // eslint-disable-next-line no-constructor-return
      return ButlerApi.apiInstance;
    }
    ButlerApi.apiInstance = this;
  }
}

export const Api = new ButlerApi({baseURL: API_BASE_URL });

export const addCommonHeader = (key: string, value: string) => {
  Api.instance.defaults.headers.common[key] = value;
}

export const addAuthorization = (accessToken: string): void => {
  addCommonHeader("Authorization", `Bearer ${accessToken}`);
};

export * from "./gen";
export default ButlerApi;
