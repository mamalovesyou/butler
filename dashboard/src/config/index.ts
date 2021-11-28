import { AxiosInstance } from 'axios';

enum PlatformEnv {
  PRODUCTION = 'production',
  DEVELOPMENT = 'development'
}

const NODE_ENV = process.env.NODE_ENV;
const PLATFORM_ENV = process.env.PLATFORM_ENV;
const VERSION = process.env.VERSION;
const API_BASE_URL = process.env.API_BASE_URL;

console.log("VALUES:", NODE_ENV, PLATFORM_ENV, VERSION, API_BASE_URL)


// type ButlerAPPConfig = {
//   platformEnv: string;
//   version: string;
//   apiBaseUrl: string;
// }

// const schema = yup.object()
//   .noUnknown()
//   .shape({
//     platformEnv: yup.string().oneOf(Object.values(PlatformEnv)).required(),
//     appVersion: yup.string().required(),
//     apiBaseUrl: yup.string().required(),
//   });

// export const AppConfig: ButlerAPPConfig = yupEnv({ schema });

export const updateAxiosInstance = (instance: AxiosInstance): void => {
  instance.defaults.baseURL = API_BASE_URL;
  instance.defaults.headers['App-Version'] = VERSION;
};

export const isProductionPlatformEnv = (): boolean => NODE_ENV === PlatformEnv.PRODUCTION;
