import {all} from "redux-saga/effects";
import {initEffects} from "./init/InitEffects";
import {authEffects} from "./auth/AuthEffects";
import {workspaceEffects} from "./workspace/WorkspaceEffects";
import ButlerApi from "../api";
import connectorEffects from "./connectors/ConnectorsEffects";
import dataSourceEffects from "./data-sources/DataSourcesEffects";
import onboardingEffects from "./onboarding/OnboardingEffects";

export const Api = new ButlerApi();

export default function configureEffects() {
    return function* rootSaga() {
        yield all([...initEffects, ...authEffects, ...workspaceEffects, ...connectorEffects, ...onboardingEffects, ...dataSourceEffects]);
    };
}
