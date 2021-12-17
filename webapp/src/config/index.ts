import { AxiosInstance } from "axios";

enum PlatformEnv {
  PROD = "prod",
  DEV = "dev",
}

const { BUILD_TARGET } = process.env;
const { APP_VERSION } = process.env;
const { API_BASE_URL } = process.env;
const { APP_BASE_URL } = process.env;

export const updateAxiosInstance = (instance: AxiosInstance): void => {
  instance.defaults.baseURL = API_BASE_URL;
  instance.defaults.headers["App-Version"] = APP_VERSION;
};

export const isProductionPlatformEnv = (): boolean =>
    BUILD_TARGET === PlatformEnv.PROD;
