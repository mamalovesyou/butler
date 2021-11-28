import { PersistedState } from 'redux-persist';
import { AuthStateType } from "../auth";

// REDUX
export const REDUX_INIT = '@@INIT';

// REDUX PERSIST
export const PERSIST_INIT = 'perist/PERSIST';
export const PERSIST_REHYDRATE = 'persist/REHYDRATE';
export const PERSIST_PURGE = 'perist/PURGE';

// ROUTER
export const ROUTER_LOCATION_CHANGE = '@@router/LOCATION_CHANGE';


export interface IActionRehydrate {
    type: typeof PERSIST_REHYDRATE,
    payload: PersistedState & { auth: AuthStateType }
}