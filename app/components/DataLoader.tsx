import React, { useContext, useEffect, useState } from 'react';
import { GlobalContext } from '../store';
import { FetchChannel, FetchQueue } from '../lib/lib';

export const DataLoader = ({ children }: any) => {
  const { dispatch } = useContext(GlobalContext);
  const [state, setState] = useState({ lock: false });

  useEffect(() => {
    setInterval(async () => {
      if (!state.lock) {
        setState({ lock: true });
        await Promise.all([
          FetchQueue(dispatch),
          FetchChannel(dispatch),
        ]);
        setState({ lock: false });
      }
    }, 3000);
  }, [dispatch]);

  return (
    <>
      {children}
    </>
  );
}
