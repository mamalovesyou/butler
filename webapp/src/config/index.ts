import { AxiosInstance } from "axios";

enum PlatformEnv {
  PRODUCTION = "production",
  DEVELOPMENT = "development",
}

const { PLATFORM_ENV } = process.env;
const { APP_VERSION } = process.env;
const { API_BASE_URL } = process.env;
const { WEBAPP_BASE_URL } = process.env;

export const updateAxiosInstance = (instance: AxiosInstance): void => {
  instance.defaults.baseURL = API_BASE_URL;
  instance.defaults.headers["App-Version"] = APP_VERSION;
};

export const isProductionPlatformEnv = (): boolean =>
    PLATFORM_ENV === PlatformEnv.PRODUCTION;
