import { combineReducers, Reducer } from "redux";
import { persistReducer } from "redux-persist";
import storage from "redux-persist/lib/storage";
import { RouterState } from "redux-first-history";
import authReducer from "./auth/AuthReducer";
import workspaceReducer from "./workspace/WorkspaceReducer";
import connectorReducer from "./connectors/ConnectorsReducer";
import onboardingReducer from "./onboarding/OnboardingReducer";
import AlertSlice from "./notifications/AlertSlice";
import dataSourceReducer from "./data-sources/DataSourceReducer";
import * as ActionType from "./auth/AuthAction.types";

// Persist config
// Using a white list to only
const persistConfig = {
  // configuration object for redux-persist
  key: "root",
  storage, // define which storage to use
  whitelist: ["auth"],
};

const workspacePersistConfig = {
    key: 'workspace',
    storage,
    whitelist: ['selected'],
};

const createRootReducer = (routerReducer: Reducer<RouterState>) => {

    const appReducer = combineReducers({
        router: routerReducer,
        notifications: AlertSlice.reducer,
        auth: authReducer,
        workspace: persistReducer(workspacePersistConfig, workspaceReducer),
        onboarding: onboardingReducer,
        connectors: connectorReducer,
        dataSources: dataSourceReducer
    });

    const rootReducer = (state, action) => {
        if (action.type === ActionType.LOGOUT) {
            return appReducer(undefined, action)
        }
        return appReducer(state, action)
    }

    return persistReducer(
        persistConfig,
        rootReducer
    );
}



export default createRootReducer;
