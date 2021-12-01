import { configureStore } from "./configureStore";

export const { store, persistor, history } = configureStore();

export type RootState = ReturnType<typeof store.getState>;

export type AppDispatch = typeof store.dispatch;
