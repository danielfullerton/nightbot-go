import React, { createContext, Dispatch, useCallback, useReducer } from 'react';
import { Channel, ChannelResponse } from './types/Channel';
import { Queue, QueueResponse, Song } from './types/Queue';

export interface GlobalState {
  channel: Channel;
  currentSong: Song;
  queue: Queue;
}

export interface GlobalAction {
  type: string;
  payload: any;
}

export interface GlobalContext {
  state: GlobalState;
  dispatch: Dispatch<GlobalAction>;
}

const initialState: GlobalState = {
  channel: null,
  currentSong: null,
  queue: []
};

export const GlobalContext = createContext({} as GlobalContext);

export const Store = ({ children }: any) => {
  const memoizedReducer = useCallback((prevState: GlobalState, action: GlobalAction) => {
    console.log(action.payload);
    switch (action.type) {
      case 'channel-updated': {
        const { channel } = action.payload as ChannelResponse;
        return { ...prevState, channel: channel };
      }
      case 'queue-updated': {
        const { queue, _currentSong } = action.payload as QueueResponse;
        return { ...prevState, queue: queue, currentSong: _currentSong };
      }
      default: {
        return { ...prevState };
      }
    }
  }, []);
  const [state, dispatch] = useReducer(memoizedReducer, initialState);
  console.log(state);
  return (
    <GlobalContext.Provider value={{state, dispatch}}>
      {children}
    </GlobalContext.Provider>
  );
};
