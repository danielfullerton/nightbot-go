import React, { createContext, Dispatch, useCallback, useReducer } from 'react';
import { Channel } from './types/Channel';
import { Queue, Song } from './types/Queue';

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
    console.log(action.type);
    return { ...prevState };
  }, []);
  const [state, dispatch] = useReducer(memoizedReducer, initialState);
  return (
    <GlobalContext.Provider value={{state, dispatch}}>
      {children}
    </GlobalContext.Provider>
  );
};
