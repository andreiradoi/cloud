// @flow weak

import React, { Component } from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom'

import { FKApiClient } from './api/api';
import '../css/App.css';

import { Landing } from './components/Landing';
import { Signin } from './components/Signin';

import type { ErrorMap } from './common/util';

export class App extends Component {
  constructor(props) {
    super(props);

    let API_HOST = 'https://fieldkit.org';
    if (process.env.NODE_ENV === 'development') {
      API_HOST = 'http://localhost:8080';
    }

    FKApiClient.setup(API_HOST, this.onLogout.bind(this));
  }

  onLogout() {
    console.log('Logged out!');
  }

  requestSignIn(username: string, password: string): Promise<?ErrorMap> {
    return FKApiClient.get().signIn(username, password);
  }

  render() {
    return (
      <Router>
        <div>
          <Route exact={true} path="/" component={Landing} />
          <Route path="/signin" render={() =>
            <Signin requestSignIn={this.requestSignIn.bind(this)} />
          }/>
        </div>
      </Router>
    );
  }
}
