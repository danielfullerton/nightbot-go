import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import { Queue } from './components/Queue';
import { Home } from './components/Home';
import { Close } from './components/Close';
import { NowPlaying } from './components/NowPlaying';
import { Channel } from './components/Channel';

export class App extends Component<any, any> {
  render() {
    return (
      <Router>
        <Switch>
          <Route exact path="/queue">
            <Queue />
          </Route>
          <Route exact path="/channel">
            <Channel />
          </Route>
          <Route exact path="/now-playing">
            <NowPlaying />
          </Route>
          <Route exact path="/close">
            <Close />
          </Route>
          <Route path="/">
            <Home />
          </Route>
        </Switch>
      </Router>
    );
  }
}
