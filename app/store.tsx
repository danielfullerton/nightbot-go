import React from 'react';
import { createContext, useReducer } from 'react';
import { Channel } from './types/Channel';
import { Queue, Song } from './types/Queue';

interface GlobalState {
  channel: Channel;
  currentSong: Song;
  queue: Queue;
}

const initialState: GlobalState = {
  channel: null,
  currentSong: null,
  queue: []
};

export const GlobalContext = createContext({});

export const Store = ({ children }: any) => {
  const [state, dispatch] = useReducer(() => {
    return null;
  }, initialState);
  return (
    <GlobalContext.Provider value={{state, dispatch}}>
      {children}
    </GlobalContext.Provider>
  );
};
