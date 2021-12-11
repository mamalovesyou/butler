import { combineReducers, Reducer } from "redux";
import { persistReducer } from "redux-persist";
import storage from "redux-persist/lib/storage";
import { RouterState } from "redux-first-history";
import authReducer from "./auth/AuthReducer";
import workspaceReducer from "./workspace/WorkspaceReducer";
import connectorReducer from "./connectors/ConnectorsReducer";
import errorReducer from "./error/ErrorReducer";
import onboardingReducer from "./onboarding/OnboardingReducer";

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
      error: errorReducer,
      auth: authReducer,
      workspace: workspaceReducer,
      onboarding: onboardingReducer,
      connectors: connectorReducer,
    })
  );

export default createRootReducer;
