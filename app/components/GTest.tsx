import React, { useContext } from 'react';
import { GlobalContext } from '../store';

export const GTest = () => {
  const context = useContext(GlobalContext);
  console.log(context);

  return (
    <div>
      Hello World
    </div>
  );
};
