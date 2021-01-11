import React, { useContext } from 'react';
import { GlobalContext } from '../store';

export const Overlay = () => {
  const { state } = useContext(GlobalContext);

  return (
    <div>
      {state.channel?.displayName}
    </div>
  );
};
