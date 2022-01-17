import { updateAxiosInstance } from "../config";
import { Api as GenAPI, ApiConfig } from "./gen";

export class ButlerApi<T> extends GenAPI<T> {
  private static apiInstance;

  constructor(apiConfig?: ApiConfig) {
    super(apiConfig);
    updateAxiosInstance(this.instance);
    if (ButlerApi.apiInstance) {
      // eslint-disable-next-line no-constructor-return
      return ButlerApi.apiInstance;
    }
    ButlerApi.apiInstance = this;
  }
}

export const getButlerApiInstance = <T>(): ButlerApi<T> => new ButlerApi();

export const addCommonHeader = (key: string, value: string) => {
  const GenAPI = getButlerApiInstance();
  GenAPI.instance.defaults.headers.common[key] = value;
  console.log(GenAPI.instance.defaults.headers)
}

export const addAuthorization = (accessToken: string): void => {
  addCommonHeader("Authorization", `Bearer ${accessToken}`);

};

export const Api = new ButlerApi();

export * from "./gen";
export default ButlerApi;
