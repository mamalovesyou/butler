import { combineReducers, Reducer } from "redux";
import { persistReducer } from "redux-persist";
import storage from "redux-persist/lib/storage";
import { RouterState } from "redux-first-history";
import authReducer from "./auth/AuthReducer";
import workspaceReducer from "./workspace/WorkspaceReducer";
import connectorReducer from "./connectors/ConnectorsReducer";
import onboardingReducer from "./onboarding/OnboardingReducer";
import AlertSlice from "./notifications/AlertSlice";

// Persist config
// Using a white list to only
const persistConfig = {
  // configuration object for redux-persist
  key: "root",
  storage, // define which storage to use
  whitelist: ["auth"],
};

const createRootReducer = (routerReducer: Reducer<RouterState>) =>
  persistReducer(
    persistConfig,
    combineReducers({
      router: routerReducer,
      notifications: AlertSlice.reducer,
      auth: authReducer,
      workspace: workspaceReducer,
      onboarding: onboardingReducer,
      connectors: connectorReducer,
    })
  );

export default createRootReducer;
