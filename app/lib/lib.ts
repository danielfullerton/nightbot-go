import { Dispatch } from 'react';
import { GlobalAction } from '../store';

interface DispatchFunction {
  (dispatch: Dispatch<GlobalAction>): void;
}

export const FetchChannel: DispatchFunction = (dispatch) => {
  dispatch({
    type: 'channel-updated',
    payload: null
  });
}

export const FetchQueue: DispatchFunction = (dispatch) => {
  dispatch({
    type: 'queue-updated',
    payload: null
  });
}
