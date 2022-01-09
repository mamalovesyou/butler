import { updateAxiosInstance } from "../config";
import { Api, ApiConfig } from "./gen";

export class ButlerApi<T> extends Api<T> {
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
  const api = getButlerApiInstance();
  api.instance.defaults.headers.common[key] = value;
  console.log(api.instance.defaults.headers)
}

export const addAuthorization = (accessToken: string): void => {
  addCommonHeader("Authorization", `Bearer ${accessToken}`);

};

export * from "./gen";
export default ButlerApi;
