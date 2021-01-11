import { Dispatch } from 'react';
import { GlobalAction } from '../store';
import axios from 'axios';

interface DispatchFunction {
  (dispatch: Dispatch<GlobalAction>): Promise<void>;
}

export const FetchChannel: DispatchFunction = async (dispatch) => {
  const channelResponse = await axios.get('/api/channel');
  dispatch({
    type: 'channel-updated',
    payload: channelResponse.data
  });
}

export const FetchQueue: DispatchFunction = async (dispatch) => {
  const queueResponse = await axios.get('/api/queue');
  dispatch({
    type: 'queue-updated',
    payload: queueResponse.data
  });
}
