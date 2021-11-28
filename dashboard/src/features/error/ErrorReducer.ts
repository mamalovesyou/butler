
import { GoogleRpcStatus } from '../../api';
import * as ActionTypes from './ErrorAction.types';

export type ErrorStateType = {
    error: GoogleRpcStatus | null;
}

const initialErrorState: ErrorStateType = {
    error: null,
}

const errorReducer = (state: ErrorStateType = initialErrorState, action: any) => {
    const { error } = action;

    // This line catch a all actions that have an error field.
    if (error) {
        return {
            error: error
        }
    } else {
        const errorAction: ActionTypes.ErrorActionType = action;
        switch (errorAction.type) {
            case ActionTypes.CLEAR_ERROR:
                return initialErrorState;
        }
    }

    return state;
}

export default errorReducer;