import {AxiosInstance} from "axios";

enum PlatformEnv {
    PROD = "prod",
    DEV = "dev",
}

export const {BUILD_TARGET} = process.env;
export const {APP_VERSION} = process.env;
export const {API_BASE_URL} = process.env;
export const {WEBAPP_BASE_URL} = process.env;

console.log("Config: ", BUILD_TARGET, APP_VERSION, API_BASE_URL, WEBAPP_BASE_URL)

export const updateAxiosInstance = (instance: AxiosInstance): void => {
    instance.defaults.baseURL = API_BASE_URL;
};

export const isProductionPlatformEnv = (): boolean =>
    BUILD_TARGET === PlatformEnv.PROD;
