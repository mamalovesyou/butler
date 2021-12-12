import { V1User } from '../../api';
import * as ActionType from './AuthAction.types';

export type AuthStateType = {
  isAuthenticated: boolean;
  accessToken: string;
  refreshToken: string;
  user: V1User;
};

const initialAuthState: AuthStateType = {
  isAuthenticated: false,
  accessToken: undefined,
  refreshToken: undefined,
  user: undefined
};

const authReducer = (
  state: AuthStateType = initialAuthState,
  action: ActionType.AuthActionType
) => {
  switch (action.type) {
    case ActionType.LOGIN_SUCCESS:
    case ActionType.SIGNUP_SUCCESS:
    case ActionType.REFRESH_TOKEN_SUCCESS:
      return { ...state, ...action.payload, isAuthenticated: true };

    case ActionType.LOGIN_FAILURE:
    case ActionType.SIGNUP_FAILURE:
    case ActionType.REFRESH_TOKEN_FAILURE:
    case ActionType.LOGOUT:
      return initialAuthState;

    default:
      return state;
  }
};

export default authReducer;
