import { createStore, applyMiddleware, StoreEnhancer } from 'redux';
import { createBrowserHistory } from 'history';
import createSagaMiddleware from 'redux-saga';
import { createReduxHistoryContext } from 'redux-first-history';
import { persistStore } from 'redux-persist';
import { composeWithDevTools } from 'redux-devtools-extension';
import createRootReducer from './configureReducer';
import configureEffects from './configureEffects';
import { isProductionPlatformEnv } from '../config';

// Create the saga middleware
const sagaMiddleware = createSagaMiddleware();

// Remove redux-dev-tools on production environment
const applyReduxDevTools = (enhancer: StoreEnhancer): StoreEnhancer => (isProductionPlatformEnv()
  ? enhancer
  : composeWithDevTools(enhancer));

export const configureStore = () => {
  
  const { createReduxHistory, routerMiddleware, routerReducer } = createReduxHistoryContext({
    history: createBrowserHistory(),
  });

  const store = createStore(
    createRootReducer(routerReducer),
    applyReduxDevTools(
      applyMiddleware(
        routerMiddleware,
        sagaMiddleware,
      ),
    ),
  );

  const rootSaga = configureEffects();
  sagaMiddleware.run(rootSaga);

  const history = createReduxHistory(store);
  const persistor = persistStore(store);
  return { store, persistor, history };
};
