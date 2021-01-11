import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import { QueueComponent } from './components/QueueComponent';
import { Home } from './components/Home';
import { Close } from './components/Close';
import { ChannelComponent } from './components/ChannelComponent';
import { Store } from './store';
import { DataLoader } from './components/DataLoader';

export class App extends Component<any, any> {
  render() {
    return (
      <Store>
        <DataLoader>
          <Router>
            <Switch>
              <Route exact path="/queue">
                <QueueComponent />
              </Route>
              <Route exact path="/channel">
                <ChannelComponent />
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
