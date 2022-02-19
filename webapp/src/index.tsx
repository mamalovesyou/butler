import React from 'react';
import ReactDOM from 'react-dom';
import { Provider as ReduxProvider } from 'react-redux';
import { PersistGate } from 'redux-persist/integration/react';
import { ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { HistoryRouter } from 'redux-first-history/rr6';
import { AppRoutes } from './routes';
import { createTheme } from './theme';
import { store, persistor, history } from './features';
import {
  SettingsConsumer,
  SettingsProvider
} from './contexts/settings-context';
import NotificationToaster from "./components/notifications-toaster";

ReactDOM.render(
  <React.StrictMode>
    <ReduxProvider store={store}>
      <PersistGate persistor={persistor}>
        <HistoryRouter history={history}>
          <SettingsProvider>
            <SettingsConsumer>
              {({ settings }) => (
                <ThemeProvider
                  theme={createTheme({
                    direction: settings.direction,
                    responsiveFontSizes: settings.responsiveFontSizes,
                    mode: settings.theme
                  })}
                >
                  <CssBaseline />
                  <NotificationToaster />
                  <AppRoutes />
                </ThemeProvider>
              )}
            </SettingsConsumer>
          </SettingsProvider>
        </HistoryRouter>
      </PersistGate>
    </ReduxProvider>
  </React.StrictMode>,
  document.getElementById('root')
);

// Hot Module Replacement (HMR) - Remove this snippet to remove HMR.
// Learn more: https://www.snowpack.dev/concepts/hot-module-replacement
if (undefined /* [snowpack] import.meta.hot */) {
  // @ts-ignore
  undefined /* [snowpack] import.meta.hot */
    .accept();
}
