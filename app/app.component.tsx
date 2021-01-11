import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import { Home } from './components/Home';
import { Close } from './components/Close';
import { Store } from './store';
import { DataLoader } from './components/DataLoader';
import { Overlay } from './components/Overlay';

export class App extends Component<any, any> {
  render() {
    return (
      <Store>
        <DataLoader>
          <Router>
            <Switch>
              <Route exact path="/overlay">
                <Overlay />
              </Route>
              <Route exact path="/close">
                <Close />
              </Route>
              <Route path="/">
                <Home />
              </Route>
            </Switch>
          </Router>
        </DataLoader>
      </Store>
    );
  }
}
