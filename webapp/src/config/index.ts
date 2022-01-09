import {AxiosInstance} from "axios";
import {addCommonHeader} from "../api";

enum PlatformEnv {
    PROD = "prod",
    DEV = "dev",
}

const {BUILD_TARGET} = process.env;
const {APP_VERSION} = process.env;
const {API_BASE_URL} = process.env;
const {APP_BASE_URL} = process.env;

console.log("env", BUILD_TARGET, APP_VERSION, API_BASE_URL, APP_BASE_URL)

export const updateAxiosInstance = (instance: AxiosInstance): void => {
    instance.defaults.baseURL = API_BASE_URL;
    // addCommonHeader("App-Version", APP_VERSION);
};

export const isProductionPlatformEnv = (): boolean =>
    BUILD_TARGET === PlatformEnv.PROD;
