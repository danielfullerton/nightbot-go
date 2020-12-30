import React from 'react'
import ReactDOM from 'react-dom'
import './styles.scss'

if (module.hot) module.hot.accept();

ReactDOM.render(
  React.createElement('div', null, `Hi there`),
  document.getElementById('app')
)
