import React from 'react'
import ReactDOM from 'react-dom'

if (module.hot) module.hot.accept();

ReactDOM.render(
  React.createElement('div', null, `Hi there`),
  document.getElementById('app')
)
